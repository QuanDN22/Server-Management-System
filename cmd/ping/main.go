package main

import (
	"context"
	"fmt"
	"log"

	"github.com/QuanDN22/Server-Management-System/internal/ping"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/pkg/kafka/producer"
	"github.com/QuanDN22/Server-Management-System/pkg/logger"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
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

	log.Println("creating ping service")

	// config
	cfg, err := config.NewConfig("./cmd/ping", ".env.ping")
	if err != nil {
		cancel()
		log.Fatalf("failed get config %v", err)
	}
	log.Println("config parsed...")

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

	fmt.Println("producer")

	// pingProducer := ping.NewProducer(ctx, cfg.PingBrokerAddress, cfg.PingTopic, l)
	pingProducer := producer.NewProducer(ctx, cfg.PingBrokerAddress, cfg.PingTopic)

	l.Info("ping producer created...")

	// middleware
	mw, err := middleware.NewMiddleware(cfg.PathPublicKey)
	// mw, err := middleware.NewMiddleware(os.Args[1])
	if err != nil {
		l.Error("failed to create middleware", zap.Error(err))
	}
	l.Info("middleware created...")

	// managementsystem Client
	managementsystemConnect, err := grpc.Dial(
		cfg.ManagementSystemServerPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect to monitor server: %v", err)
	}
	defer managementsystemConnect.Close()

	managementClient := managementsystem.NewManagementSystemClient(managementsystemConnect)

	// pingService := ping.NewPingService(ctx, pingProducer, l, managementClient)

	_ = ping.NewPingService(ctx, pingProducer, l, managementClient)

	// pingService.Start(ctx, uint(1000))
}
