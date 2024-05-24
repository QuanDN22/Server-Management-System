package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// "localhost:6379"

func NewRedisCache(ctx context.Context, host string, port string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Ping Redis to check if the connection is working
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	return client, nil
}
