package producer

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
)

func NewProducer(ctx context.Context, BrokerAddress string, topic string) *kafka.Writer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{BrokerAddress},
		Topic:   topic,
	})

	return w
}
