package utils

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"

	"th-iot-server/config"
)

// SendVerificationEmail 发送邮箱验证码，并控制发送频率
func SendVerificationEmail(email string) error {
	return WithRedisCtx(5*time.Second, func(ctx context.Context, rdb *redis.Client) error {
		emailKey := fmt.Sprintf("email:verify:%s", email)
		limitKey := fmt.Sprintf("email:limit:%s", email)

		// 防止频繁请求
		if rdb.Exists(ctx, limitKey).Val() > 0 {
			return fmt.Errorf("请求过于频繁，请稍后再试")
		}

		code, err := generateSecureCode(6)
		if err != nil {
			return fmt.Errorf("生成验证码失败: %w", err)
		}

		// 保存验证码 5 分钟有效
		if err := rdb.Set(ctx, emailKey, code, 5*time.Minute).Err(); err != nil {
			return fmt.Errorf("保存验证码失败: %w", err)
		}

		// 限制 1 分钟只能发一次
		if err := rdb.Set(ctx, limitKey, "1", time.Minute).Err(); err != nil {
			return fmt.Errorf("设置发送限制失败: %w", err)
		}

		// 发送邮件
		if err := sendEmail(email, code); err != nil {
			return fmt.Errorf("发送邮件失败: %w", err)
		}

		return nil
	})
}

// VerifyEmailCode 校验邮箱验证码是否正确
func VerifyEmailCode(email, code string) (bool, error) {
	var success bool

	err := WithRedisCtx(5*time.Second, func(ctx context.Context, rdb *redis.Client) error {
		key := fmt.Sprintf("email:verify:%s", email)
		storedCode, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			return fmt.Errorf("验证码不存在或已过期")
		} else if err != nil {
			return fmt.Errorf("查询验证码失败: %w", err)
		}

		if storedCode != code {
			return fmt.Errorf("验证码错误")
		}

		// 验证成功后删除验证码，避免重复使用
		if err := rdb.Del(ctx, key).Err(); err != nil {
			return fmt.Errorf("删除验证码失败: %w", err)
		}

		success = true
		return nil
	})

	return success, err
}

// 发送邮件
func sendEmail(to, code string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.SMTPUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "【Go 验证码测试】")
	m.SetBody("text/plain", fmt.Sprintf("您的验证码是：%s（5分钟内有效）", code))

	d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SMTPUser, config.SMTPPass)
	return d.DialAndSend(m)
}

// 生成安全随机验证码
func generateSecureCode(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	code := ""
	for _, v := range b {
		code += fmt.Sprintf("%d", v%10)
	}
	return code, nil
}
