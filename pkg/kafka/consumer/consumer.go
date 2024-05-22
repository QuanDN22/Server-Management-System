package consumer

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
)

func NewConsumer(ctx context.Context, BrokerAddress string, topic string, groupID string) *kafka.Reader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{BrokerAddress},
		Topic:       topic,
		GroupID:     groupID,
		StartOffset: kafka.FirstOffset,
	})

	return r
}
