package gRPCServer

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/QuanDN22/Server-Management-System/internal/management-system/domain"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/proto/auth"
	"github.com/QuanDN22/Server-Management-System/proto/mail"

	managementsystem "github.com/QuanDN22/Server-Management-System/proto/management-system"
	mt "github.com/QuanDN22/Server-Management-System/proto/monitor"
)

type ManagementSystemGrpcServer struct {
	managementsystem.UnimplementedManagementSystemServer
	config     *config.Config
	logger     *zap.Logger
	gRPCServer *grpc.Server
	db         *gorm.DB
	cache      *redis.Client

	monitorConsumer *kafka.Reader
	monitorClient   mt.MonitorClient

	mailClient mail.MailClient
	authClient auth.AuthServiceClient
}

func NewManagementSystemGrpcServer(
	config *config.Config,
	logger *zap.Logger,
	grpcserver *grpc.Server,
	db *gorm.DB,
	cache *redis.Client,
	monitorConsumer *kafka.Reader,
	monitorClient mt.MonitorClient,
	mailClient mail.MailClient,
	authClient auth.AuthServiceClient,
) (ms *ManagementSystemGrpcServer) {
	ms = &ManagementSystemGrpcServer{
		config:          config,
		logger:          logger,
		gRPCServer:      grpcserver,
		db:              db,
		cache:           cache,
		monitorConsumer: monitorConsumer,
		monitorClient:   monitorClient,
		mailClient:      mailClient,
		authClient:      authClient,
	}

	// Attach the Greeter service to the server
	managementsystem.RegisterManagementSystemServer(ms.gRPCServer, ms)
	return ms
}

func (ms *ManagementSystemGrpcServer) Start(ctx context.Context, cancel context.CancelFunc) {
	// grpc server
	go func() {
		// Create listening on TCP port
		lis, err := net.Listen("tcp", ms.config.ManagementSystemServerPort)
		if err != nil {
			cancel()
			ms.logger.Info("Failed to listen: ", zap.Error(err), zap.String("port", ms.config.ManagementSystemServerPort))
			return
		}

		// Serve gRPC Server
		ms.logger.Info("Management System gRPC server started", zap.String("port", ms.config.ManagementSystemServerPort))
		if err := ms.gRPCServer.Serve(lis); err != nil {
			cancel()
			ms.logger.Info("error starting grpc server", zap.Error(err), zap.String("port", ms.config.ManagementSystemServerPort))
			return
		}

		<-ctx.Done()
		if err := lis.Close(); err != nil {
			cancel()
			ms.logger.Info("error closing listener", zap.Error(err))
			return
		}
	}()

	// start consumer of topic monitor
	go func() {
		for {
			// the `ReadMessage` method blocks until we receive the next event
			msg, err := ms.monitorConsumer.ReadMessage(ctx)
			if err != nil {
				panic("could not read message " + err.Error())
			}

			go ms.Woker(ctx, msg)
		}
	}()

	// send email work daily
	go ms.WorkDailyReport()

	<-ctx.Done()
}

func (ms *ManagementSystemGrpcServer) Woker(ctx context.Context, msg kafka.Message) {
	if msg.Topic == ms.config.MonitorTopic {
		// split value
		s := string(msg.Value)
		fmt.Println(s)

		value := strings.Split(s, ",")

		server_id, err := strconv.Atoi(value[0])
		server_status := value[1]

		if err != nil {
			fmt.Println("error convert string to int")
			return
		}

		fmt.Println("receive: ", server_id, server_status)

		// save to database
		var server domain.Server
		res := ms.db.Model(&domain.Server{}).Where("server_id = ?", server_id).First(&server)
		if res.RowsAffected == 0 {
			fmt.Println("Server not found")
			return
		}

		server.Server_Status = server_status
		ms.db.Save(&server)

		fmt.Println("server_id: ", server_id, " - server status updated")
	}
}
