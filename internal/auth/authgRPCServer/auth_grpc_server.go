package authgRPCServer

import (
	"context"
	"net"

	"github.com/QuanDN22/Server-Management-System/internal/auth/issuer"
	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/proto/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AuthGrpcServer struct {
	auth.UnimplementedAuthServiceServer
	iss        *issuer.Issuer // create token
	config     *config.Config
	logger     *zap.Logger
	gRPCServer *grpc.Server
}

// // login
// func (a *AuthGrpcServer) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {

// }

// // signup
// func (a *AuthGrpcServer) Signup(ctx context.Context, in *auth.SignupRequest) (*emptypb.Empty, error) {

// }

// // logout
// func (a *AuthGrpcServer) Logout(ctx context.Context, in *auth.LogoutRequest) (*emptypb.Empty, error) {

// }

// // change password
// func (a *AuthGrpcServer) ChangePassword(ctx context.Context, in *auth.ChangePasswordRequest) (*emptypb.Empty, error) {

// }

// admin delete a user by ID
// func (a *AuthGrpcServer) DeleteUserByID(ctx context.Context, in *auth.DeleteUserRequest) (*emptypb.Empty, error) {
// }

func NewAuthGRPCServer(
	issuer *issuer.Issuer,
	config *config.Config,
	logger *zap.Logger,
	grpcserver *grpc.Server,
) (s *AuthGrpcServer) {
	s = &AuthGrpcServer{
		iss:        issuer,
		config:     config,
		logger:     logger,
		gRPCServer: grpcserver,
	}

	// Attach the Greeter service to the server
	auth.RegisterAuthServiceServer(s.gRPCServer, s)
	return s
}

func (a *AuthGrpcServer) Start(ctx context.Context, cancel context.CancelFunc) {
	// Create listening on TCP port
	lis, err := net.Listen("tcp", a.config.GrpcPort)
	if err != nil {
		cancel()
		a.logger.Info("Failed to listen: ", zap.Error(err), zap.String("port", a.config.GrpcPort))
		return
	}

	// Serve gRPC Server
	a.logger.Info("Auth gRPC server started", zap.String("port", a.config.GrpcPort))
	if err := a.gRPCServer.Serve(lis); err != nil {
		cancel()
		a.logger.Info("error starting grpc server", zap.Error(err), zap.String("port", a.config.GrpcPort))
		return
	}

	<-ctx.Done()
	if err := lis.Close(); err != nil {
		cancel()
		a.logger.Info("error closing listener", zap.Error(err))
		return
	}
}
