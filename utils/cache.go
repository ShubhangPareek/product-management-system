package utils

import (
    "context"
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
    err := redisClient.Set(ctx, key, value, 10*time.Minute).Err()
    return err
}

func GetCache(key string) (string, error) {
    ctx := context.Background()
    value, err := redisClient.Get(ctx, key).Result()
    return value, err
}