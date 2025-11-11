package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// 响应结构
type ResponseFormat struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResponseMiddleware 包装返回数据为统一格式
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 捕获响应
		bodyWriter := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		c.Next()

		// Controller 自行跳过中间件
		if _, ok := c.Get("skipResponseMiddleware"); ok {
			return
		}

		statusCode := c.Writer.Status()
		respBody := bodyWriter.body.Bytes()

		var data interface{}
		if len(respBody) > 0 {
			if err := json.Unmarshal(respBody, &data); err != nil {
				data = string(respBody) // 非 JSON 内容直接返回字符串
			}
		}

		c.JSON(statusCode, ResponseFormat{
			Code: func() int { if statusCode >= 400 { return 1 } else { return 0 } }(),
			Msg:  func() string { if statusCode >= 400 { return "error" } else { return "success" } }(),
			Data: data,
		})
	}
}

// bodyWriter 用于捕获响应
type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
