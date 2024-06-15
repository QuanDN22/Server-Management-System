package monitor

import (
	"context"
	"net"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/QuanDN22/Server-Management-System/pkg/config"

	managementsystem "github.com/QuanDN22/Server-Management-System/proto/management-system"
	mt "github.com/QuanDN22/Server-Management-System/proto/monitor"
)

type MonitorService struct {
	mt.UnimplementedMonitorServer
	MonitorProducer  *kafka.Writer
	managementClient managementsystem.ManagementSystemClient
	config           *config.Config
	logger           *zap.Logger
	gRPCServer       *grpc.Server
	elasticClient    *elasticsearch.TypedClient
}

func NewMonitorService(
	MonitorProducer *kafka.Writer,

	managementClient managementsystem.ManagementSystemClient,

	logger *zap.Logger,
	config *config.Config,
	gRPCServer *grpc.Server,
	elasticClient *elasticsearch.TypedClient,
) (ms *MonitorService) {
	ms = &MonitorService{
		MonitorProducer: MonitorProducer,

		managementClient: managementClient,
		config:           config,
		logger:           logger,
		gRPCServer:       gRPCServer,
		elasticClient:    elasticClient,
	}

	// Attach the Monitor service to the grpc server
	mt.RegisterMonitorServer(ms.gRPCServer, ms)
	return ms
}

func (m *MonitorService) Start(ctx context.Context) {
	// grpc server
	go func() {
		// Create listening on TCP port
		lis, err := net.Listen("tcp", m.config.MonitorServerPort)
		if err != nil {
			m.logger.Info("Failed to listen: ", zap.Error(err), zap.String("port", m.config.MonitorServerPort))
			return
		}

		// Serve gRPC Server
		m.logger.Info("Management System gRPC server started", zap.String("port", m.config.MonitorServerPort))
		if err := m.gRPCServer.Serve(lis); err != nil {
			m.logger.Info("error starting grpc server", zap.Error(err), zap.String("port", m.config.MonitorServerPort))
			return
		}

		<-ctx.Done()
		if err := lis.Close(); err != nil {
			m.logger.Info("error closing listener", zap.Error(err))
			return
		}
	}()

	// monitor server
	go m.WorkDailyMonitorServer(ctx)

	<-ctx.Done()
}
