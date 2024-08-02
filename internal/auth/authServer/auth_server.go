package gRPCServer

import (
	"context"
	"net"
	"net/http"

	"github.com/QuanDN22/Server-Management-System/internal/auth/issuer"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/proto/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	gRPCServer *grpc.Server
	httpServer *http.Server

	issuer     *issuer.Issuer // create token
	config     *config.Config
	logger     *zap.Logger
	db         *gorm.DB
}

func NewAuthServer(
	grpcserver *grpc.Server,
	httpserver *http.Server,
	issuer *issuer.Issuer,
	config *config.Config,
	logger *zap.Logger,
	db *gorm.DB,
) (s *AuthServer) {
	s = &AuthServer{
		gRPCServer: grpcserver,
		httpServer: httpserver,
		issuer:     issuer,
		config:     config,
		logger:     logger,
		db:         db,
	}

	// Attach the Auth service to the server
	auth.RegisterAuthServiceServer(s.gRPCServer, s)
	return s
}

func (a *AuthServer) Start(ctx context.Context, cancel context.CancelFunc) {
	// Create listening on TCP port
	lis, err := net.Listen("tcp", a.config.AuthServerGrpcPort)
	if err != nil {
		cancel()
		a.logger.Info("Failed to listen: ", zap.Error(err), zap.String("port", a.config.AuthServerGrpcPort))
		return
	}

	// http server
	go func(ctx context.Context, cancel context.CancelFunc) {
		a.logger.Info("Serving http is running on %s", zap.String("port", a.config.AuthServerHttpPort))
		if err := a.httpServer.ListenAndServe(); err != nil {
			cancel()
			a.logger.Info("error starting http server", zap.Error(err), zap.String("port", a.config.AuthServerHttpPort))
			return
		}
	}(ctx, cancel)

	// Serve gRPC Server
	go func(ctx context.Context, cancel context.CancelFunc) {
		a.logger.Info("Auth gRPC server started", zap.String("port", a.config.AuthServerGrpcPort))
		if err := a.gRPCServer.Serve(lis); err != nil {
			cancel()
			a.logger.Info("error starting grpc server", zap.Error(err), zap.String("port", a.config.AuthServerGrpcPort))
			return
		}
	}(ctx, cancel)

	<-ctx.Done()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		a.logger.Info("Error shutting down HTTP server", zap.Error(err))
	}

	if err := lis.Close(); err != nil {
		cancel()
		a.logger.Info("error closing listener", zap.Error(err))
		return
	}
}
