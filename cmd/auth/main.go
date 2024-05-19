package main

import (
	"context"
	"log"

	"github.com/QuanDN22/Server-Management-System/internal/auth/authgRPCServer"
	"github.com/QuanDN22/Server-Management-System/internal/auth/issuer"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/pkg/logger"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	"github.com/QuanDN22/Server-Management-System/pkg/postgres"
	"go.uber.org/zap"

	"google.golang.org/grpc"
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

	// new logger
	l, err := logger.NewLogger(cfg.ServiceName)
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

	// l.Info("issuer created...")
	// token, err := i.IssueToken("admin", nil)
	// if err != nil {
	// 	cancel()
	// 	log.Fatal(err)
	// }
	// l.Info("token created... " + token)

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
	// // Auto migrate the Server model
	// err = db.AutoMigrate(&domain.User{})
	// if err != nil {
	// 	log.Fatalf("Failed to migrate servers datable: %v", err)
	// } else {
	// 	log.Println("migrate servers datable successfully")
	// }

	// users := []domain.User{
	// 	{UserName: "quan1", Password: "1", Email: "quan1@gmail.com", Role: "user"},
	// 	{UserName: "quan2", Password: "2", Email: "quan2@gmail.com", Role: "user"},
	// 	{UserName: "quan3", Password: "3", Email: "quan3@gmail.com", Role: "user"},

	// 	{UserName: "admin", Password: "pass", Email: "admin@gmail.com", Role: "admin"},
	// }

	// for _, user := range users {
	// 	user.Password, _ = utils.GenerateHashPassword(user.Password)
	// 	db.Create(&user)
	// }

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
	authgrpcserver := authgRPCServer.NewAuthGRPCServer(i, cfg, l, grpcserver, db)
	authgrpcserver.Start(ctx, cancel)
}
