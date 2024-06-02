package monitor

import (
	"context"
	"net"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	"github.com/QuanDN22/Server-Management-System/pkg/config"
	mt "github.com/QuanDN22/Server-Management-System/proto/monitor"
	"github.com/elastic/go-elasticsearch/v8"
	"google.golang.org/grpc"
)

type MonitorService struct {
	mt.UnimplementedMonitorServer
	MonitorProducer *kafka.Writer
	MonitorConsumer *kafka.Reader
	config          *config.Config
	logger          *zap.Logger
	gRPCServer      *grpc.Server
	elasticClient   *elasticsearch.TypedClient
}

func NewMonitorService(
	MonitorProducer *kafka.Writer,
	MonitorConsumer *kafka.Reader,
	logger *zap.Logger,
	config *config.Config,
	gRPCServer *grpc.Server,
	elasticClient *elasticsearch.TypedClient,
) (ms *MonitorService) {
	ms = &MonitorService{
		MonitorProducer: MonitorProducer,
		MonitorConsumer: MonitorConsumer,
		config:          config,
		logger:          logger,
		gRPCServer:      gRPCServer,
		elasticClient:   elasticClient,
	}

	// Attach the Monitor service to the grpc server
	mt.RegisterMonitorServer(ms.gRPCServer, ms)
	return ms
}

func (m *MonitorService) Start(ctx context.Context) {
	go m.StartMonitorConsumer(ctx)
	go m.StartMonitorProducer(ctx)
	// grpc server
	go func() {
		// Create listening on TCP port
		lis, err := net.Listen("tcp", m.config.MonitorServerPort)
		if err != nil {
			m.logger.Info("Failed to listen: ", zap.Error(err), zap.String("port", ":5003"))
			return
		}

		// Serve gRPC Server
		m.logger.Info("Management System gRPC server started", zap.String("port", ":5003"))
		if err := m.gRPCServer.Serve(lis); err != nil {
			m.logger.Info("error starting grpc server", zap.Error(err), zap.String("port", ":5003"))
			return
		}

		<-ctx.Done()
		if err := lis.Close(); err != nil {
			m.logger.Info("error closing listener", zap.Error(err))
			return
		}
	}()
	<-ctx.Done()
}
