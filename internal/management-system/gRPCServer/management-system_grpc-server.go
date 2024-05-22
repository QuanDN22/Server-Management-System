package gRPCServer

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/QuanDN22/Server-Management-System/internal/management-system/domain"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	managementsystem "github.com/QuanDN22/Server-Management-System/proto/management-system"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ManagementSystemGrpcServer struct {
	managementsystem.UnimplementedManagementSystemServer
	config       *config.Config
	logger       *zap.Logger
	gRPCServer   *grpc.Server
	db           *gorm.DB
	pingConsumer *kafka.Reader

	monitorConsumer *kafka.Reader
	monitorProducer *kafka.Writer
}

func NewManagementSystemGrpcServer(
	config *config.Config,
	logger *zap.Logger,
	grpcserver *grpc.Server,
	db *gorm.DB,
	pingConsumer *kafka.Reader,
	monitorConsumer *kafka.Reader,
	monitorProducer *kafka.Writer,
) (ms *ManagementSystemGrpcServer) {
	ms = &ManagementSystemGrpcServer{
		config:          config,
		logger:          logger,
		gRPCServer:      grpcserver,
		db:              db,
		pingConsumer:    pingConsumer,
		monitorConsumer: monitorConsumer,
		monitorProducer: monitorProducer,
	}

	// Attach the Greeter service to the server
	managementsystem.RegisterManagementSystemServer(ms.gRPCServer, ms)
	return ms
}

func (ms *ManagementSystemGrpcServer) Start(ctx context.Context, cancel context.CancelFunc) {
	// start consumer of topic ping
	// go func() {
	// 	for {
	// 		// the `ReadMessage` method blocks until we receive the next event
	// 		msg, err := ms.pingConsumer.ReadMessage(ctx)
	// 		if err != nil {
	// 			panic("could not read message " + err.Error())
	// 		}

	// 		go ms.Woker(ctx, msg)
	// 	}
	// }()

	// start consumer of topic monitor
	// go func() {
	// 	for {
	// 		// the `ReadMessage` method blocks until we receive the next event
	// 		msg, err := ms.monitorConsumer.ReadMessage(ctx)
	// 		if err != nil {
	// 			panic("could not read message " + err.Error())
	// 		}

	// 		go ms.Woker(ctx, msg)
	// 	}
	// }()

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

	<-ctx.Done()
}

func (ms *ManagementSystemGrpcServer) Woker(ctx context.Context, msg kafka.Message) {
	if msg.Topic == ms.config.PingTopic {
		split := strings.Split(string(msg.Value), ",")
		server_id := split[0]
		status := split[1]
		fmt.Println(server_id, status)
		var server domain.Server

		res := ms.db.First(&server, server_id)
		if res.RowsAffected == 0 {
			ms.logger.Info("server not found")
			return
		}

		server.Server_Status = status
		res = ms.db.Save(&server)
		if res.Error == nil {
			ms.logger.Info("server updated")
		} else {
			ms.logger.Info("Failed to update user")
		}
		return
	}

	if msg.Topic == ms.config.MonitorTopic {
		// consumer topic
		fmt.Println(string(msg.Value))
		time_monitor := string(msg.Value)
		var server_ids []uint

		res := ms.db.Model(&domain.Server{}).Select("server_id").Where("server_status = ?", "on").Find(&server_ids)

		if res.Error != nil {
			ms.logger.Info("Failed to get server")
			return
		}

		// producer result topic
		resultTopic := struct {
			TimeMonitor string `json:"time_monitor"`
			ServerIDs   []uint `json:"server_ids"`
		}{
			TimeMonitor: time_monitor,
			ServerIDs:   server_ids,
		}

		// Marshal the struct/map to JSON for efficient message encoding
		messageBytes, err := json.Marshal(resultTopic)
		if err != nil {
			ms.logger.Error("Failed to marshal resultTopic: " + err.Error())
			return // Handle error appropriately
		}

		err = ms.monitorProducer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(msg.Value),
			Value: []byte(messageBytes),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}
	}
}
