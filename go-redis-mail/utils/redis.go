package utils

import (
    "context"
    "github.com/redis/go-redis/v9"
    "go-redis-mail/config"
)

var (
    Rdb = redis.NewClient(&redis.Options{
        Addr:     config.RedisAddr,
        Password: config.RedisPassword,
        DB:       config.RedisDB,
    })
    Ctx = context.Background()
)

