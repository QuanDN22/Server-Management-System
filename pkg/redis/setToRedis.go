package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func SetToRedis(ctx context.Context, redisClient *redis.Client, key, val string) error {
	err := redisClient.Set(ctx, key, val, 30).Err()
	if err != nil {
		return err
	}

	return nil
}
