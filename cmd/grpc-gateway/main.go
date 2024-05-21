package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/pkg/logger"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authpb "github.com/QuanDN22/Server-Management-System/proto/auth"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("creating gRPC-gateway")

	// config
	cfg, err := config.NewConfig("./cmd/grpc-gateway", ".env.grpc-gateway")
	if err != nil {
		cancel()
		log.Fatalf("failed get config %v", err)
	}
	log.Println("config parsed...")

	// new logger
	// Create a logger with lumberjack integration
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

	mw, err := middleware.NewMiddleware(cfg.PathPublicKey)
	// mw, err := middleware.NewMiddleware(os.Args[1])
	if err != nil {
		l.Error("failed to create middleware", zap.Error(err))
	}
	l.Info("middleware created...")

	// Create a client connection to the gRPC server we just created
	// This is where the gRPC-gateway proxies the requests
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptor),
	}

	err = authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, gwmux, cfg.AuthServerPort, opts)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    cfg.GrpcGatewayPort,
		Handler: mw.HandleHTTP(gwmux),
	}

	l.Info(fmt.Sprintf("Serving gRPC-Gateway is running on %s", cfg.GrpcGatewayPort))
	log.Fatalln(gwServer.ListenAndServe())
}
