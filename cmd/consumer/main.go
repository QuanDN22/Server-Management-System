package main

import (
	"fmt"
	"log"

	"github.com/QuanDN22/Server-Management-System/pkg/config"
)

func main() {
	// config
	_, err := config.NewConfig("./cmd/consumer", ".env.consumer")
	if err != nil {
		log.Fatalf("failed get config %v", err)
	}
	fmt.Println("consumer")
	// ctx := context.Background()
	// cs := consumer.NewConsumer(ctx, cfg.KafkaBrokerAddress, cfg.KafkaTopic, cfg.KafkaConsumerGroupId)
	// // cs := consumer.NewConsumer(ctx)
	// cs.Start(ctx)
}
