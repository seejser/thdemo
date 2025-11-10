package utils

import (
    "fmt"
    "gopkg.in/gomail.v2"
    "go-redis-mail/config"
)

func SendEmail(to string, code string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", config.SMTPUser)
    m.SetHeader("To", to)
    m.SetHeader("Subject", "【Go 验证码测试】")
    m.SetBody("text/plain", fmt.Sprintf("您的验证码是：%s（5分钟内有效）", code))

    d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SMTPUser, config.SMTPPass)
    return d.DialAndSend(m)
}

