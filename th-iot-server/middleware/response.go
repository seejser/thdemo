package middleware

import (
    "fmt"
    "net/http"
    "runtime/debug"

    "github.com/gin-gonic/gin"
)

const ResponseDataKey = "response_data"

// APIResponse 统一响应结构
type APIResponse struct {
    Code    int         `json:"code"`
    Message string      `json:"msg"`
    Data    interface{} `json:"data"`
    TraceId string      `json:"trace_id,omitempty"`
    Stack   string      `json:"stack,omitempty"`
}

// 中间件，自动处理响应
func ResponseMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 捕获 panic
        defer func() {
            if r := recover(); r != nil {
                err := fmt.Errorf("%v", r)
                ReturnError(c, 500, err, http.StatusInternalServerError)
            }
        }()

        // 继续执行请求
        c.Next()

        // 检查是否已经被 ReturnError 或 ReturnSuccess 处理过
        if _, exists := c.Get(ResponseDataKey); exists {
            return
        }

        // 获取处理结果
        respData, exists := c.Get("resp_data")
        if exists {
            // 如果业务中间件或 handler 已经设置 resp_data
            ReturnSuccess(c, respData)
            return
        }

        // 如果 handler 返回 nil，默认返回空 data
        if c.Writer.Written() {
            // 已经写入 response，什么也不做
            return
        }

        ReturnSuccess(c, gin.H{})
    }
}

// ReturnSuccess 统一成功返回
func ReturnSuccess(c *gin.Context, data interface{}) {
    resp := &APIResponse{
        Code:    0,
        Message: "success",
        Data:    data,
        TraceId: getTraceId(c),
    }
    c.Set(ResponseDataKey, resp)
    c.JSON(http.StatusOK, resp)
}

// ReturnError 统一错误返回
func ReturnError(c *gin.Context, code int, err error, httpStatus ...int) {
    status := http.StatusOK
    if len(httpStatus) > 0 {
        status = httpStatus[0]
    }
    resp := &APIResponse{
        Code:    code,
        Message: err.Error(),
        Data:    nil,
        TraceId: getTraceId(c),
        Stack:   string(debug.Stack()),
    }
    c.Set(ResponseDataKey, resp)
    c.Error(err)
    c.JSON(status, resp)
    c.Abort()
}

// 可自定义 TraceId 获取逻辑
func getTraceId(c *gin.Context) string {
    if v, ok := c.Get("trace_id"); ok {
        return fmt.Sprintf("%v", v)
    }
    return ""
}
