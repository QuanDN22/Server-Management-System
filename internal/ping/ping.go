package ping

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type PingService struct {
	logger   *zap.Logger
	PingProducer *kafka.Writer
}

// PingBrokerAddress string, PingTopic string,

func NewPingService(ctx context.Context, PingProducer *kafka.Writer, logger *zap.Logger) *PingService {
	return &PingService{
		PingProducer: PingProducer,
		logger:   logger,
	}
}

func (p *PingService) Start(ctx context.Context, numberOfServer uint) {
	time.Sleep(time.Second * 3)
	for {
		var src = rand.NewSource(time.Now().UnixNano())
		var r = rand.New(src)

		server_id := r.Intn(int(numberOfServer)) + 1
		status := "off"
		if r.Intn(2) == 1 {
			status = "on"
		}

		err := p.PingProducer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(server_id)),
			Value: []byte(fmt.Sprintf("%d,%s", server_id, status)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		msg := fmt.Sprintf(`{"server_id": %d, "status": %s}`, server_id, status)
		p.logger.Info(msg)
		time.Sleep(time.Second * 1)
	}
}
