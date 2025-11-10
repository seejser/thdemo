# go-redis-mail

用 go+gin+redis 实现邮箱验证码

## 创建项目

```sh
#创建项目gon-demo
go mod init go-redis-mail

#安装检查
go mod tidy

# 运行
go run main.go
```

## 测试

```sh
# 发送验证码接口
curl -X POST http://localhost:9090/send -H "Content-Type: application/json" -d '{"email": "seejser@gmail.com"}'
```
```go
# 图片base64Captcha
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/mojocn/base64Captcha"
    "github.com/redis/go-redis/v9"
    "context"
    "time"
    "fmt"
)

var ctx = context.Background()
var store *redis.Client

func init() {
    store = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
}

func main() {
    r := gin.Default()

    r.GET("/captcha", func(c *gin.Context) {
        // 1️⃣ 生成图形验证码
        driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
        captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

        id, b64s, answer := captcha.Driver.GenerateIdQuestionAnswer()
        item, _ := captcha.Driver.DrawCaptcha(answer)

        // 2️⃣ 保存到 Redis (1分钟有效)
        store.Set(ctx, fmt.Sprintf("captcha:%s", id), answer, time.Minute)

        // 3️⃣ 返回给前端
        c.JSON(200, gin.H{
            "captcha_id": id,
            "image":      item.EncodeB64string(),
        })
    })

    r.Run(":8080")
}


```
