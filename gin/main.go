// package main
// import (
//     "net/http"
//     "github.com/gin-gonic/gin"
// )
// func main() {
//     r := gin.Default()
//     r.GET("/ping", func(c *gin.Context) {
//         c.JSON(http.StatusOK, gin.H{"message": "pong"})
//     })
//     r.Run(":9090")
// }

package main

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
    // 1️⃣ 创建 Redis 客户端
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis 地址
        Password: "",               // 无密码则留空
        DB:       0,                // 使用默认DB
    })

    // 2️⃣ 测试连接
    pong, err := rdb.Ping(ctx).Result()
    if err != nil {
        fmt.Println("❌ Redis 连接失败:", err)
        return
    }
    fmt.Println("✅ Redis 连接成功:", pong)

    // 3️⃣ 写入数据
    err = rdb.Set(ctx, "test_key", "hello redis", 0).Err()
    if err != nil {
        fmt.Println("❌ 写入失败:", err)
        return
    }

    // 4️⃣ 读取数据
    val, err := rdb.Get(ctx, "test_key").Result()
    if err != nil {
        fmt.Println("❌ 读取失败:", err)
        return
    }
    fmt.Println("📦 读取到的值:", val)

    // 5️⃣ 测试过期键
    err = rdb.Set(ctx, "temp_key", "expire test", 5 * time.Second).Err()
    if err != nil {
        fmt.Println("❌ 设置过期键失败:", err)
        return
    }
    fmt.Println("⏱ 已设置 temp_key，5秒后自动过期")
}

