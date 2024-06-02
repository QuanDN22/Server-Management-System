package mail

import (
	"context"
	"net"

	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"github.com/QuanDN22/Server-Management-System/proto/mail"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type MailService struct {
	mail.UnimplementedMailServer
	config *config.Config
	logger *zap.Logger

	gRPCServer *grpc.Server
}

func NewMailService(
	logger *zap.Logger,
	config *config.Config,
	gRPCServer *grpc.Server,
) (ms *MailService) {
	ms = &MailService{
		config:     config,
		logger:     logger,
		gRPCServer: gRPCServer,
	}

	// Attach the Mail service to the grpc server
	mail.RegisterMailServer(ms.gRPCServer, ms)
	return ms

}

func (ms *MailService) Start(ctx context.Context, cancel context.CancelFunc) {
	// Create listening on TCP port
	lis, err := net.Listen("tcp", ms.config.MailServerPort)
	if err != nil {
		cancel()
		ms.logger.Info("Failed to listen: ", zap.Error(err), zap.String("port", ms.config.MailServerPort))
		return
	}

	// Serve gRPC Server
	ms.logger.Info("Mail gRPC server started", zap.String("port", ms.config.MailServerPort))
	if err := ms.gRPCServer.Serve(lis); err != nil {
		cancel()
		ms.logger.Info("error starting grpc server", zap.Error(err), zap.String("port", ms.config.MailServerPort))
		return
	}

	<-ctx.Done()
	if err := lis.Close(); err != nil {
		cancel()
		ms.logger.Info("error closing listener", zap.Error(err))
		return
	}
}
