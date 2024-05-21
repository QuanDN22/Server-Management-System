package kafka

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type Producer struct {
	Producer *kafka.Writer
}

func NewProducer(ctx context.Context, BrokerAddress string, topic string) *Producer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{BrokerAddress},
		Topic:   topic,
	})

	return &Producer{
		Producer: w,
	}
}

func (p *Producer) Start(ctx context.Context) {
	for {
		var src = rand.NewSource(time.Now().UnixNano())
		var r = rand.New(src)

		user_id := r.Intn(100) + 1
		status := r.Intn(2) == 1

		err := p.Producer.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(user_id)),
			Value: []byte(fmt.Sprintf("%d,%t", user_id, status)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		msg := fmt.Sprintf(`{"user_id": %d, "status": %t}`, user_id, status)
		fmt.Println(msg)
		time.Sleep(time.Second * 3)
	}
}
