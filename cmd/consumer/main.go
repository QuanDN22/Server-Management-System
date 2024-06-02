package main

import (
	"fmt"
	"log"

	"github.com/QuanDN22/Server-Management-System/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	// create a connection
	conn, err := grpc.Dial(
		"localhost:5003",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// //create a client
	// client := mt.NewMonitorClient(conn)

	// res, err := client.GetUpTime(context.Background(), &emptypb.Empty{})

	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Response: %v", res.Uptime)
}
