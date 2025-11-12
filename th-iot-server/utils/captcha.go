package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
)

// GenerateCaptcha 生成图形验证码并保存到 Redis
func GenerateCaptcha() (string, string, error) {
	// 配置验证码：数字 + 字母
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80) // 高80，宽240，6位
	c := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	// 生成验证码
	id, b64s,code, err := c.Generate()
	if err != nil {
		return "", "", fmt.Errorf("生成验证码失败: %w", err)
	}

	// 保存到 Redis，5分钟过期
	err = WithRedisCtx(5*time.Second, func(ctx context.Context, rdb *redis.Client) error {
		return rdb.Set(ctx, fmt.Sprintf("captcha:%s", id), code, 5*time.Minute).Err()
	})
	if err != nil {
		return "", "", err
	}

	return id, b64s, nil
}


// VerifyCaptcha 验证图形验证码
func VerifyCaptcha(captchaId, code string) (bool, error) {
	if captchaId == "" || code == "" {
		return false, fmt.Errorf("验证码 ID 或 Code 为空")
	}

	var valid bool
	err := WithRedisCtx(3*time.Second, func(ctx context.Context, rdb *redis.Client) error {
		key := fmt.Sprintf("captcha:%s", captchaId)
		stored, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			return fmt.Errorf("验证码不存在或已过期")
		} else if err != nil {
			return fmt.Errorf("获取验证码失败: %w", err)
		}

		// ✅ 打印 Redis 中保存的验证码和传入的验证码
		fmt.Printf("[CAPTCHA] 校验中 => Redis值: %s, 输入值: %s\n", stored, code)

		if stored == code {
			valid = true
			// 验证成功后删除，避免复用
			_ = rdb.Del(ctx, key).Err()
			fmt.Printf("[CAPTCHA] ✅ 验证通过 (ID: %s)\n", captchaId)
		} else {
			fmt.Printf("[CAPTCHA] ❌ 验证失败 (ID: %s)\n", captchaId)
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return valid, nil
}
