package monitor

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func (m *MonitorService) StartMonitorProducer(ctx context.Context) {
	for {
		err := m.MonitorProducer.WriteMessages(ctx, kafka.Message{
			Key:   []byte("monitor"),
			Value: []byte(time.Now().String()),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}
		time.Sleep(time.Second * 1)
	}
}
