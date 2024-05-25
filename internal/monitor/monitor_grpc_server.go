package monitor

import (
	"context"
	"net"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	mt "github.com/QuanDN22/Server-Management-System/proto/monitor"
	"github.com/elastic/go-elasticsearch/v8"
	"google.golang.org/grpc"
)

type MonitorService struct {
	MonitorProducer *kafka.Writer
	MonitorConsumer *kafka.Reader
	logger          *zap.Logger

	mt.UnimplementedMonitorServer
	grpc          *grpc.Server
	elasticClient *elasticsearch.TypedClient
}

func NewMonitorService(MonitorProducer *kafka.Writer, MonitorConsumer *kafka.Reader, logger *zap.Logger, grpc *grpc.Server, elasticClient *elasticsearch.TypedClient) *MonitorService {
	return &MonitorService{
		MonitorProducer: MonitorProducer,
		MonitorConsumer: MonitorConsumer,
		logger:          logger,
		grpc:            grpc,
		elasticClient:   elasticClient,
	}
}

func (m *MonitorService) Start(ctx context.Context) {
	go m.StartMonitorConsumer(ctx)
	go m.StartMonitorProducer(ctx)
	// grpc server
	go func() {
		// Create listening on TCP port
		lis, err := net.Listen("tcp", "localhost:5003")
		if err != nil {
			m.logger.Info("Failed to listen: ", zap.Error(err), zap.String("port", ":5003"))
			return
		}

		// Serve gRPC Server
		m.logger.Info("Management System gRPC server started", zap.String("port", ":5003"))
		if err := m.grpc.Serve(lis); err != nil {
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
