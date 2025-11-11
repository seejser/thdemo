package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"th-iot-server/services"
	"th-iot-server/utils"
)

func GetCaptcha(c *gin.Context) {
	ctx := c.Request.Context() // 使用请求上下文
	id, img, err := utils.GenerateCaptcha(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": "生成验证码失败"})
		return
	}

	c.JSON(200, gin.H{
		"captcha_id": id,
		"image":      img,
	})
}
func GetEmailCode(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	if err := services.RegisterUser(req.Username, req.Password, req.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	if err := services.RegisterUser(req.Username, req.Password, req.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	user, err := services.LoginUser(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: 生成 JWT 返回
	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "user": user})
}
