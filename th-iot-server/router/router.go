package router

import (
	"github.com/gin-gonic/gin"
	"th-iot-server/controllers"
	"th-iot-server/middleware"
)

func InitRouter() *gin.Engine {
	//  默认带 Logger + Recovery 中间件
	r := gin.Default()

	// ✅ 全局中间件
	r.Use(
		middleware.CorsMiddleware(),     // 跨域
		middleware.ResponseMiddleware(), // 统一响应
	)

	// ✅ 路由分组
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/captcha", controllers.GetCaptcha)
			auth.GET("/email_code", controllers.GetEmailCode)
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
			auth.POST("/refresh", controllers.RefreshToken)
		}
		// 需要登录的接口
		user := api.Group("/user")
		user.Use(middleware.AuthMiddleware())
		{
			user.GET("/info", controllers.UserInfo)
		}
	}

	return r
}
