package gRPCServer

import (
	"context"
	"net"

	"github.com/QuanDN22/Server-Management-System/pkg/config"
	managementsystem "github.com/QuanDN22/Server-Management-System/proto/management-system"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ManagementSystemGrpcServer struct {
	managementsystem.UnimplementedManagementSystemServer
	config     *config.Config
	logger     *zap.Logger
	gRPCServer *grpc.Server
	db         *gorm.DB
}

func NewManagementSystemGrpcServer(
	config *config.Config,
	logger *zap.Logger,
	grpcserver *grpc.Server,
	db *gorm.DB,
) (ms *ManagementSystemGrpcServer) {
	ms = &ManagementSystemGrpcServer{
		config:     config,
		logger:     logger,
		gRPCServer: grpcserver,
		db:         db,
	}

	// Attach the Greeter service to the server
	managementsystem.RegisterManagementSystemServer(ms.gRPCServer, ms)
	return ms
}

func (ms *ManagementSystemGrpcServer) Start(ctx context.Context, cancel context.CancelFunc) {
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
}
