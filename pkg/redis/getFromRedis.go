package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func GetFromRedis(ctx context.Context, redisClient *redis.Client , key string) string {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val
}
