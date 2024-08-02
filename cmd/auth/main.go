package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	authpb "github.com/QuanDN22/Server-Management-System/proto/auth"
	"github.com/QuanDN22/Server-Management-System/internal/auth/domain"
	authServer "github.com/QuanDN22/Server-Management-System/internal/auth/authServer"
	"github.com/QuanDN22/Server-Management-System/internal/auth/issuer"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/pkg/logger"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	"github.com/QuanDN22/Server-Management-System/pkg/postgres"
	"github.com/QuanDN22/Server-Management-System/pkg/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("creating auth service")

	// config
	cfg, err := config.NewConfig("./cmd/auth", ".env.auth")
	if err != nil {
		cancel()
		log.Fatalf("failed get config %v", err)
	}
	log.Println("config parsed...")

	// logger
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

	// issuer
	i, err := issuer.NewIssuer(cfg.PathPrivateKey)
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	l.Info("issuer created...")

	// database
	db := postgres.NewPostgresDB(cfg.PGDatabaseHost, cfg.PGDatabaseUser, cfg.PGDatabasePassword, cfg.PGDatabaseDBName, cfg.PGDatabasePort)
	l.Info("database connected...")
	// // delete table if it doesn't exist
	// err = db.Migrator().DropTable(&domain.User{})
	// if err != nil {
	// 	log.Fatalf("Failed to drop table servers: %v", err)
	// } else {
	// 	log.Println("Dropped table servers")
	// }
	// Auto migrate the Server model
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		// log.Fatalf("Failed to migrate servers datable: %v", err)
		l.Info("failed to migrate servers datable")
	} else {
		// log.Println("migrate servers datable successfully")
		l.Info("migrate servers datable successfully")
	}

	l.Info("migrating database...")

	users := []domain.User{
		{UserName: "quan1", Password: "1", Email: "quan1@gmail.com", Role: "user"},
		{UserName: "quan2", Password: "2", Email: "quan2@gmail.com", Role: "user"},
		{UserName: "quan3", Password: "3", Email: "quan3@gmail.com", Role: "user"},
		{UserName: "quan4", Password: "4", Email: "quan4@gmail.com", Role: "user"},
		{UserName: "quan5", Password: "5", Email: "quan5@gmail.com", Role: "user"},
		{UserName: "quan6", Password: "6", Email: "quan6@gmail.com", Role: "user"},
		{UserName: "quan7", Password: "7", Email: "quan7@gmail.com", Role: "user"},
		{UserName: "quan8", Password: "8", Email: "quan8@gmail.com", Role: "user"},
		{UserName: "quan9", Password: "9", Email: "quan9@gmail.com", Role: "user"},
		{UserName: "quan10", Password: "10", Email: "quan10@gmail.com", Role: "user"},

		{UserName: "admin1", Password: "1", Email: "admin1@gmail.com", Role: "admin"},
		{UserName: "admin2", Password: "2", Email: "admin2@gmail.com", Role: "admin"},
		{UserName: "admin3", Password: "3", Email: "admin3@gmail.com", Role: "admin"},
	}

	for _, user := range users {
		user.Password, _ = utils.GenerateHashPassword(user.Password)
		db.Create(&user)
	}

	mw, err := middleware.NewMiddleware(cfg.PathPublicKey)
	if err != nil {
		l.Error("failed to create middleware", zap.Error(err))
	}
	l.Info("middleware created...")

	// gRPC server
	grpcserver := grpc.NewServer(
		grpc.UnaryInterceptor(mw.UnaryServerInterceptor),
	)

	l.Info("grpc server created...")

	// Http server
	// Create a client connection to the gRPC server we just created
	// This is where the gRPC-gateway proxies the requests
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(mw.StreamClientInterceptor),
	}

	gwmux.HandlePath("GET", "/v1/api/auth/hello", handleHello)

	err = authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, gwmux, cfg.AuthServerGrpcPort, opts)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	httpserver := &http.Server{
		Addr:    cfg.AuthServerHttpPort,
		Handler: mw.HandleHTTP(gwmux),
	}

	// Auth server
	authserver := authServer.NewAuthServer(grpcserver, httpserver, i, cfg, l, db)
	l.Info("auth server created...")

	// Start the server
	l.Info("auth grpc server started...")
	authserver.Start(ctx, cancel)
}


func handleHello(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("creating auth service")

	// config
	cfg, err := config.NewConfig("./cmd/auth", ".env.auth")
	if err != nil {
		cancel()
		log.Fatalf("failed get config %v", err)
	}
	log.Println("config parsed...")
	// logger
	l, _ := logger.NewLogger(
		fmt.Sprintf("%s%s.log", cfg.LogFilename, cfg.ServiceName),
		int(cfg.LogMaxSize),
		int(cfg.LogMaxBackups),
		int(cfg.LogMaxAge),
		true,
		zapcore.InfoLevel,
	)
	l.Info("hello world")
	w.Write([]byte("hello world")) //nolint
}