package main

import (
	"context"
	"fmt"
	"log"

	"github.com/QuanDN22/Server-Management-System/internal/mail"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/pkg/logger"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("creating mail service")

	// config
	cfg, err := config.NewConfig("./cmd/mail", ".env.mail")
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

	// middleware
	mw, err := middleware.NewMiddleware(cfg.PathPublicKey)
	// mw, err := middleware.NewMiddleware(os.Args[1])
	if err != nil {
		l.Error("failed to create middleware", zap.Error(err))
	}
	l.Info("middleware created...")

	// grpc server
	grpcserver := grpc.NewServer(
		grpc.UnaryInterceptor(mw.UnaryServerInterceptor),
	)

	mailService := mail.NewMailService(l, cfg, grpcserver)

	mailService.Start(ctx, cancel)
}
