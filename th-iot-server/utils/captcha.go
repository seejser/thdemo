package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/mojocn/base64Captcha"
	// "th-iot-server/config"
)

func GenerateCaptcha(ctx context.Context) (string, string, error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	id, _, answer := captcha.Driver.GenerateIdQuestionAnswer()
	item, err := captcha.Driver.DrawCaptcha(answer)
	if err != nil {
		return "", "", err
	}
	// 2️⃣ 保存到 Redis (1分钟有效)
	// store.Set(ctx, fmt.Sprintf("captcha:%s", id), answer, time.Minute)
	Rdb.Set(ctx, fmt.Sprintf("captcha:%s", id), answer, 10*time.Minute)
	return id, item.EncodeB64string(), nil
}
