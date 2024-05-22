package main

import (
	"context"
	"fmt"
	"log"

	"github.com/QuanDN22/Server-Management-System/internal/ping"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
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

	pingService := ping.NewPingService(ctx, pingProducer, l)

	pingService.Start(ctx, uint(cfg.NumberOfServer))
}
