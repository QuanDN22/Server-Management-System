package main

import (
	"context"
	"log"

	"github.com/QuanDN22/Server-Management-System/internal/auth/authgRPCServer"
	"github.com/QuanDN22/Server-Management-System/internal/auth/issuer"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/pkg/logger"
	"github.com/QuanDN22/Server-Management-System/pkg/postgres"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("creating auth service")

	// config
	cfg, err := config.NewConfig("./pkg/config/", ".env.auth")
	if err != nil {
		cancel()
		log.Fatalf("failed get config %v", err)
	}
	log.Println("config parsed...")

	// new logger
	l, err := logger.NewLogger(cfg.ServiceName)
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	l.Info("logger created...")

	// connect to database
	_ = postgres.NewPostgresDB(cfg.PGDatabaseHost, cfg.PGDatabaseUser, cfg.PGDatabasePassword, cfg.PGDatabaseDBName, cfg.PGDatabasePort)
	l.Info("database connected...")

	// issuer
	i, err := issuer.NewIssuer(cfg.PathPrivateKey)
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	l.Info("issuer created...")
	// token, err := i.IssueToken("admin", nil)
	// if err != nil {
	// 	cancel()
	// 	log.Fatal(err)
	// }
	// l.Info("token created... " + token)

	// grpc server
	grpcserver := grpc.NewServer()
	authgrpcserver := authgRPCServer.NewAuthGRPCServer(i, cfg, l, grpcserver)

	l.Info("grpc server started...")
	authgrpcserver.Start(ctx, cancel)
}
