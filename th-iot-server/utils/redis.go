package utils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"th-iot-server/config"
)

var (
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
	Ctx = context.Background()
)

func RedisCtx(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(Ctx, timeout)
}
