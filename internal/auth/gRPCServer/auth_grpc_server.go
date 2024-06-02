package gRPCServer

import (
	"context"
	"net"

	"github.com/QuanDN22/Server-Management-System/internal/auth/issuer"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/proto/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type AuthGrpcServer struct {
	auth.UnimplementedAuthServiceServer
	issuer     *issuer.Issuer // create token
	config     *config.Config
	logger     *zap.Logger
	gRPCServer *grpc.Server
	db         *gorm.DB
}

func NewAuthGRPCServer(
	issuer *issuer.Issuer,
	config *config.Config,
	logger *zap.Logger,
	grpcserver *grpc.Server,
	db *gorm.DB,
) (s *AuthGrpcServer) {
	s = &AuthGrpcServer{
		issuer:     issuer,
		config:     config,
		logger:     logger,
		gRPCServer: grpcserver,
		db:         db,
	}

	// Attach the Auth service to the server
	auth.RegisterAuthServiceServer(s.gRPCServer, s)
	return s
}

func (a *AuthGrpcServer) Start(ctx context.Context, cancel context.CancelFunc) {
	// Create listening on TCP port
	lis, err := net.Listen("tcp", a.config.AuthServerPort)
	if err != nil {
		cancel()
		a.logger.Info("Failed to listen: ", zap.Error(err), zap.String("port", a.config.AuthServerPort))
		return
	}

	// Serve gRPC Server
	a.logger.Info("Auth gRPC server started", zap.String("port", a.config.AuthServerPort))
	if err := a.gRPCServer.Serve(lis); err != nil {
		cancel()
		a.logger.Info("error starting grpc server", zap.Error(err), zap.String("port", a.config.AuthServerPort))
		return
	}

	<-ctx.Done()
	if err := lis.Close(); err != nil {
		cancel()
		a.logger.Info("error closing listener", zap.Error(err))
		return
	}
}
