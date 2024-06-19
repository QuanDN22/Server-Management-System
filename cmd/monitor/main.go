package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/QuanDN22/Server-Management-System/internal/monitor"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/pkg/kafka/producer"
	"github.com/QuanDN22/Server-Management-System/pkg/logger"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	managementsystem "github.com/QuanDN22/Server-Management-System/proto/management-system"
)

func main() {
	// This is the main entry point for the application
	// It should be used to start the application
	// and handle any errors that occur during the application's lifecycle

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("creating monitor service")

	// config
	cfg, err := config.NewConfig("./cmd/monitor", ".env.monitor")
	if err != nil {
		cancel()
		log.Fatalf("failed get config %v", err)
	}
	log.Println("config parsed...")

	fmt.Println("monitor duration: ", cfg.MonitorDurationMinute)
	fmt.Println("token_internal: ", cfg.TokenInternal)
	fmt.Println("max concurrent ping servers: ", cfg.MaxConurrentPingServers)

	// new logger
	l, err := logger.NewLogger(
		fmt.Sprintf("%s%s.log", cfg.LogFilename, cfg.ServiceName),
		int(cfg.LogMaxSize),
		int(cfg.LogMaxBackups),
		int(cfg.LogMaxAge),
		true,
		zapcore.InfoLevel,
	)
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	l.Info("logger created...")

	monitor_producer := producer.NewProducer(ctx, cfg.MonitorBrokerAddress, cfg.MonitorTopic)
	l.Info("monitor producer created...")

	// middleware
	mw, err := middleware.NewMiddleware(cfg.PathPublicKey)
	if err != nil {
		l.Error("failed to create middleware", zap.Error(err))
	}
	l.Info("middleware created...")

	// connect to elasticsearch
	es, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		// Addresses: []string{"http://127.0.0.1:9200"},
		Addresses: []string{"http://elasticsearch:9200"},

		Logger: &elastictransport.ColorLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true},
	})

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	l.Info("elasticsearch connected...")

	res, err := es.Info().Do(context.Background())
	if err != nil {
		log.Fatalf("error reading Info request: %s", err)
	}

	if res.Tagline != "You Know, for Search" {
		log.Fatalf("invalid tagline, got: %s", res.Tagline)
	}

	l.Info("elasticsearch info: ", zap.String("tagline", res.Tagline))

	// create an index named uptime-server-monitor
	// and provide a mapping for
	// the field timestamp which will be date
	// and the field server_id which will be integer
	// and the field duration which will be integer
	indexName := "uptime-server-monitor"
	// If the index doesn't exist we create it with a mapping.
	if exists, err := es.Indices.Exists(indexName).IsSuccess(context.Background()); !exists && err == nil {
		res, err := es.Indices.Create(indexName).
			Mappings(&types.TypeMapping{
				Properties: map[string]types.Property{
					"timestamp": types.DateProperty{},
					"server_id": types.IntegerNumberProperty{},
					"duration":  types.IntegerNumberProperty{},
				},
			}).
			Do(context.Background())

		if err != nil {
			log.Fatalf("error creating index uptime-server-monitor: %s", err)
		}

		if !res.Acknowledged && res.Index != indexName {
			log.Fatalf("unexpected error during index creation, got : %#v", res)
		}

		l.Info("index uptime-server-monitor created...")
	} else if err != nil {
		log.Fatal(err)
	}

	// managementsystem Client
	managementsystemConnect, err := grpc.Dial(
		cfg.ManagementSystemServerPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect to management server: %v", err)
	}
	defer managementsystemConnect.Close()

	managementsystemClient := managementsystem.NewManagementSystemClient(managementsystemConnect)

	l.Info("management system client connected...")

	// grpc server
	grpcserver := grpc.NewServer(
		grpc.UnaryInterceptor(mw.UnaryServerInterceptor),
	)

	l.Info("grpc server created...")

	monitorService := monitor.NewMonitorService(monitor_producer, managementsystemClient, l, cfg, grpcserver, es)

	l.Info("monitor service created...")

	l.Info("monitor service starting...")
	monitorService.Start(ctx)
}
