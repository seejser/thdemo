package utils

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
	"th-iot-server/config"
)

// sendEmail 发送验证码邮件
func sendEmail(to string, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.SMTPUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "【Go 验证码测试】")
	m.SetBody("text/plain", fmt.Sprintf("您的验证码是：%s（5分钟内有效）", code))

	d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SMTPUser, config.SMTPPass)
	return d.DialAndSend(m)
}

// SendVerificationEmail 发送邮箱验证码，并控制发送频率
func SendVerificationEmail(ctx context.Context, email string) error {
	emailKey := fmt.Sprintf("email:verify:%s", email)
	limitKey := fmt.Sprintf("email:limit:%s", email)

	// 限制发送频率（1分钟内只能发一次）
	if Rdb.Exists(ctx, limitKey).Val() > 0 {
		return fmt.Errorf("请求过于频繁，请稍后再试")
	}

	// 生成6位验证码
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	// 保存到Redis，5分钟过期
	if err :=Rdb.Set(ctx, emailKey, code, 5*time.Minute).Err(); err != nil {
		return fmt.Errorf("保存验证码失败: %w", err)
	}

	// 设置发送限制键，1分钟过期
	if err := Rdb.Set(ctx, limitKey, "1", time.Minute).Err(); err != nil {
		return fmt.Errorf("设置发送限制失败: %w", err)
	}

	// 发送邮件
	if err := sendEmail(email, code); err != nil {
		return fmt.Errorf("发送邮件失败: %w", err)
	}

	return nil
}
