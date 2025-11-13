package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"th-iot-server/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			ReturnError(c, 401, fmt.Errorf("未提供 token"))
			c.Abort()
			return
		}

		// 支持 "Bearer xxx" 格式
		if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
			tokenStr = tokenStr[7:]
		}

		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			ReturnError(c, 401, fmt.Errorf("token 无效"))
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

