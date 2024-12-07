package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitCache(redisAddr string) {
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
}

func SetCache(key string, value interface{}) error {
	ctx := context.Background()
	if key == "" {
		return fmt.Errorf("cache key cannot be empty")
	}

	err := redisClient.Set(ctx, key, value, 10*time.Second).Err()
	return err
}

func GetCache(key string) (string, error) {
	ctx := context.Background()
	if key == "" {
		return "", fmt.Errorf("cache key cannot be empty")
	}

	value, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("cache miss for key: %s", key)
	}
	return value, err
}
