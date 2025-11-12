package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"th-iot-server/middleware"
	"th-iot-server/services"
	"th-iot-server/utils"
)

func GetCaptcha(c *gin.Context) {
	//ctx := c.Request.Context() // 使用请求上下文
	id, img, err := utils.GenerateCaptcha()
	if err != nil {
		middleware.ReturnError(c, 500, fmt.Errorf("生成验证码失败: %w", err))
		return
	}
	middleware.ReturnSuccess(c, gin.H{
		"captcha_id": id,
		"image":      img,
	})
}
func GetEmailCode(c *gin.Context) {
	//ctx := c.Request.Context() // 使用请求上下文
	var req struct {
		Username    string `json:"username" form:"username"`
		Email       string `json:"email" form:"email"`
		CaptchaCode string `json:"captcha_code" form:"captcha_code"`
		CaptchaId   string `json:"captcha_id" form:"captcha_id"`
	}
	if err := c.ShouldBind(&req); err != nil {
		middleware.ReturnError(c, 400, fmt.Errorf("参数错误: %w", err))
		return
	}
	// 验证图形验证码
	valid, err := utils.VerifyCaptcha(req.CaptchaId, req.CaptchaCode)
	if err != nil || !valid {
		middleware.ReturnError(c, 400, fmt.Errorf("图形验证码错误"))
		return
	}

	// 发送邮件验证码
	if err := utils.SendVerificationEmail(req.Email); err != nil {
		middleware.ReturnError(c, 500, fmt.Errorf("发送邮件验证码失败: %w", err))
		return
	}

	middleware.ReturnSuccess(c, gin.H{
		"message": "邮件验证码发送成功",
	})
}

func Register(c *gin.Context) {
	var req struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		EmailCode string `json:"email_code"` // 邮件验证码
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.ReturnError(c, 400, fmt.Errorf("参数错误: %w", err))
		return
	}

	if req.Username == "" || req.Password == "" || req.Email == "" {
		middleware.ReturnError(c, 400, fmt.Errorf("用户名、邮箱或密码不能为空"))
		return
	}

	// 验证邮件验证码
	valid, err := utils.VerifyEmailCode(req.Email, req.EmailCode)
	if err != nil || !valid {
		middleware.ReturnError(c, 400, fmt.Errorf("邮箱验证码错误"))
		return
	}

	if err := services.RegisterUser(req.Username, req.Password, req.Email); err != nil {
		middleware.ReturnError(c, 400, err)
		return
	}

	middleware.ReturnSuccess(c, gin.H{"message": "注册成功"})
}

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.ReturnError(c, 400, fmt.Errorf("参数错误"))
		return
	}

	user, err := services.LoginUser(req.Username, req.Password)
	if err != nil {
		middleware.ReturnError(c, 400, err)
		return
	}
	// token 逻辑
	accessToken, refreshToken, err := utils.GenerateTokenPair(user.ID, user.Username)
	if err != nil {
		middleware.ReturnError(c, 500, fmt.Errorf("生成令牌失败: %w", err))
		return
	}

	middleware.ReturnSuccess(c, gin.H{"access_token": accessToken, "refresh_token": refreshToken, "message": "登录成功"})
}

// RefreshToken 刷新 Access Token
func RefreshToken(c *gin.Context) {
	// 从请求中获取 refresh_token
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.ReturnError(c, 400, err)
		return
	}

	// 调用 utils 刷新 access token
	newAccessToken, err := utils.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		middleware.ReturnError(c, 401, err)
		return
	}

	middleware.ReturnSuccess(c, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": req.RefreshToken,
		"message":       "token刷新成功",
	})
}

// 获取当前登录用户信息
func UserInfo(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		middleware.ReturnError(c, 401, fmt.Errorf("用户未登录"))
		return
	}

	userClaims, ok := claims.(*utils.Claims)
	if !ok {
		middleware.ReturnError(c, 500, fmt.Errorf("Claims 类型错误"))
		return
	}

	user, err := services.GetUserByID(userClaims.UserID) // 使用 services
	if err != nil {
		middleware.ReturnError(c, 404, fmt.Errorf("用户不存在"))
		return
	}

	middleware.ReturnSuccess(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"created":  user.CreatedAt,
	})
}
