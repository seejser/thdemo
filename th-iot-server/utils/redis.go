package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"th-iot-server/config"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

func InitRedis() error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	ctx, cancel := context.WithTimeout(Ctx, 5*time.Second)
	defer cancel()

	if err := Rdb.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("Redis 连接失败: %w", err)
	}

	fmt.Println("[Redis] ✅ 连接成功:", config.RedisAddr)
	return nil
}

func CloseRedis() {
	if Rdb != nil {
		if err := Rdb.Close(); err != nil {
			fmt.Println("[Redis] ❌ 关闭失败:", err)
		} else {
			fmt.Println("[Redis] ✅ 已关闭连接")
		}
	}
}

// WithRedisCtx 安全执行 Redis 操作
func WithRedisCtx(timeout time.Duration, fn func(ctx context.Context, rdb *redis.Client) error) error {
	ctx, cancel := context.WithTimeout(Ctx, timeout)
	defer cancel()

	if Rdb == nil {
		return fmt.Errorf("Redis 客户端未初始化")
	}

	return fn(ctx, Rdb)
}
