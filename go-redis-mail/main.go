package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"
    "go-redis-mail/utils"
)

func main() {
    r := gin.Default()

    r.POST("/send", func(c *gin.Context) {
        type Request struct {
            Email string `json:"email" binding:"required,email"`
        }
        var req Request
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱格式不正确"})
            return
        }

        emailKey := fmt.Sprintf("email:verify:%s", req.Email)
        limitKey := fmt.Sprintf("email:limit:%s", req.Email)

        // 限制发送频率（1分钟内只能发一次）
        if utils.Rdb.Exists(utils.Ctx, limitKey).Val() == 1 {
            c.JSON(http.StatusTooManyRequests, gin.H{"error": "请求过于频繁，请稍后再试"})
            return
        }

        // 生成6位验证码
        rand.Seed(time.Now().UnixNano())
        code := fmt.Sprintf("%06d", rand.Intn(1000000))

        // 保存到Redis，5分钟过期
        err := utils.Rdb.Set(utils.Ctx, emailKey, code, 5*time.Minute).Err()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis写入失败"})
            return
        }

        // 限制键，1分钟过期
        utils.Rdb.Set(utils.Ctx, limitKey, "1", time.Minute)

        // 发送邮件
        if err := utils.SendEmail(req.Email, code); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "邮件发送失败"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "验证码发送成功"})
    })

    r.Run(":8080")
}

