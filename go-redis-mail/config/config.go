package config

var (
    RedisAddr     = "localhost:6379"
    RedisPassword = ""
    RedisDB       = 0

    SMTPHost = "smtp.qq.com"
    SMTPPort = 587
    SMTPUser = "你的邮箱@qq.com" // 发送邮箱
    SMTPPass = "授权码"         // 邮箱授权码（不是密码）
)

