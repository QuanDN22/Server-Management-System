package main

import (
	"context"
	"fmt"
	"log"

	"github.com/QuanDN22/Server-Management-System/internal/monitor"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/pkg/kafka/consumer"
	"github.com/QuanDN22/Server-Management-System/pkg/kafka/producer"
	"github.com/QuanDN22/Server-Management-System/pkg/logger"
	"go.uber.org/zap/zapcore"
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

	monitor_consumer := consumer.NewConsumer(ctx, cfg.MonitorBrokerAddress, cfg.MonitorResultsTopic, cfg.MonitorConsumerGroupID)

	l.Info("monitor consumer created...")

	monitor_producer := producer.NewProducer(ctx, cfg.MonitorBrokerAddress, cfg.MonitorTopic)
	l.Info("monitor producer created...")

	monitorService := monitor.NewMonitorService(monitor_producer, monitor_consumer, l)

	monitorService.Start(ctx)
}
