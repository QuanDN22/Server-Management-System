package monitor

import (
	"context"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type MonitorService struct {
	MonitorProducer *kafka.Writer
	MonitorConsumer *kafka.Reader
	logger          *zap.Logger
}

func NewMonitorService(MonitorProducer *kafka.Writer, MonitorConsumer *kafka.Reader, logger *zap.Logger) *MonitorService {
	return &MonitorService{
		MonitorProducer: MonitorProducer,
		MonitorConsumer: MonitorConsumer,
		logger:          logger,
	}
}

func (m *MonitorService) Start(ctx context.Context) {
	go m.StartMonitorConsumer(ctx)
	go m.StartMonitorProducer(ctx)
	<-ctx.Done()
}
