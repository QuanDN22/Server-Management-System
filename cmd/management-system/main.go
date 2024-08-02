package main

import (
	"context"
	"fmt"
	"log"

	"github.com/QuanDN22/Server-Management-System/internal/management-system/domain"
	"github.com/QuanDN22/Server-Management-System/internal/management-system/gRPCServer"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/pkg/kafka/consumer"
	"github.com/QuanDN22/Server-Management-System/pkg/logger"
	"github.com/QuanDN22/Server-Management-System/pkg/middleware"
	"github.com/QuanDN22/Server-Management-System/pkg/postgres"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/QuanDN22/Server-Management-System/proto/auth"
	"github.com/QuanDN22/Server-Management-System/proto/mail"
	mt "github.com/QuanDN22/Server-Management-System/proto/monitor"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("creating management system server")

	// config
	cfg, err := config.NewConfig("./cmd/management-system", ".env.management-system")
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

	// database
	db := postgres.NewPostgresDB(cfg.PGDatabaseHost, cfg.PGDatabaseUser, cfg.PGDatabasePassword, cfg.PGDatabaseDBName, cfg.PGDatabasePort)
	l.Info("database connected...")
	// // delete table if it doesn't exist
	// err = db.Migrator().DropTable(&domain.Server{})
	// if err != nil {
	// 	log.Fatalf("Failed to drop table servers: %v", err)
	// } else {
	// 	log.Println("Dropped table servers")
	// }
	// Auto migrate the Server model
	err = db.AutoMigrate(&domain.Server{})
	if err != nil {
		log.Fatalf("Failed to migrate servers datable: %v", err)
	} else {
		log.Println("migrate servers datable successfully")
	}

	l.Info("database migrated...")

	// users := []domain.Server{
	// 	{Server_Name: "server#1", Server_IPv4: "192.168.1.1", Server_Status: "on"},
	// 	{Server_Name: "server#2", Server_IPv4: "192.168.1.2", Server_Status: "off"},
	// 	{Server_Name: "server#3", Server_IPv4: "192.168.1.3", Server_Status: "off"},
	// 	{Server_Name: "server#4", Server_IPv4: "192.168.1.4", Server_Status: "on"},
	// 	{Server_Name: "server#5", Server_IPv4: "192.168.1.5", Server_Status: "off"},
	// 	{Server_Name: "server#6", Server_IPv4: "192.168.1.6", Server_Status: "off"},
	// 	{Server_Name: "server#7", Server_IPv4: "192.168.1.7", Server_Status: "on"},
	// 	{Server_Name: "server#8", Server_IPv4: "192.168.1.8", Server_Status: "on"},
	// 	{Server_Name: "server#9", Server_IPv4: "192.168.1.9", Server_Status: "on"},
	// 	{Server_Name: "server#10", Server_IPv4: "192.168.1.10", Server_Status: "off"},
	// }

	// for _, user := range users {
	// 	db.Create(&user)
	// }

	// cache
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	// Ping Redis to check if the connection is working
	// pong, err := client.Ping(ctx).Result()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(pong)

	// middleware
	mw, err := middleware.NewMiddleware(cfg.PathPublicKey)
	if err != nil {
		l.Error("failed to create middleware", zap.Error(err))
	}
	l.Info("middleware created...")

	// monitor Consumer
	monitorConsumer := consumer.NewConsumer(ctx, cfg.MonitorBrokerAddress, cfg.MonitorTopic, cfg.MonitorConsumerGroupID)
	l.Info("monitor consumer created...")

	// monitor Client
	monitorConnect, err := grpc.Dial(
		cfg.MonitorServerPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect to monitor server: %v", err)
	}
	defer monitorConnect.Close()
	l.Info("monitor client created...")

	monitorClient := mt.NewMonitorClient(monitorConnect)

	// mail client
	mailConnect, err := grpc.Dial(
		cfg.MailServerPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptor),
	)

	if err != nil {
		log.Fatalf("did not connect to mail server: %v", err)
	}
	defer mailConnect.Close()
	l.Info("mail client created...")

	mailClient := mail.NewMailClient(mailConnect)

	// auth client
	authConnect, err := grpc.Dial(
		cfg.AuthServerGrpcPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptor),
	)

	if err != nil {
		log.Fatalf("did not connect to auth server: %v", err)
	}
	defer authConnect.Close()
	l.Info("auth client created...")

	authClient := auth.NewAuthServiceClient(authConnect)

	// grpc server
	grpcserver := grpc.NewServer(
		grpc.UnaryInterceptor(mw.UnaryServerInterceptor),
		grpc.StreamInterceptor(mw.StreamServerInterceptor),
	)

	management_system_grpcserver := gRPCServer.NewManagementSystemGrpcServer(
		cfg,
		l,
		grpcserver,
		db,
		nil,
		monitorConsumer,
		monitorClient,
		mailClient,
		authClient,
	)

	management_system_grpcserver.Start(ctx, cancel)
}
