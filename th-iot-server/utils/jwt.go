package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT 签名密钥
var jwtKey = []byte("th-iot-server-secret")

// 自定义声明结构体
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成 Access Token 和 Refresh Token
func GenerateTokenPair(userID uint, username string) (accessToken, refreshToken string, err error) {
	now := time.Now()

	// Access Token（2 小时有效）
	accessClaims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "th-iot-server",
			Subject:   "access_token",
		},
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	// Refresh Token（7 天有效）
	refreshClaims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "th-iot-server",
			Subject:   "refresh_token",
		},
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// 验证 Token 并返回 Claims
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效的 token")
}

// 刷新 Access Token（需提供 refresh token）
func RefreshAccessToken(refreshToken string) (newAccessToken string, err error) {
	claims, err := ParseToken(refreshToken)
	if err != nil {
		return "", errors.New("refresh token 无效")
	}

	// 确保是 refresh_token 类型
	if claims.Subject != "refresh_token" {
		return "", errors.New("无效的 refresh token 类型")
	}

	// 如果 refresh token 已过期
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return "", errors.New("refresh token 已过期，请重新登录")
	}

	// 生成新的 access token（有效期2小时）
	newClaims := &Claims{
		UserID:   claims.UserID,
		Username: claims.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "th-iot-server",
			Subject:   "access_token",
		},
	}
	newAccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims).SignedString(jwtKey)
	if err != nil {
		return "", errors.New("生成新 access token 失败")
	}
	return newAccessToken, nil
}
