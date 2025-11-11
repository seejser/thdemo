package router

import (
	"github.com/gin-gonic/gin"
	"th-iot-server/controllers"
	"th-iot-server/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1/auth")
	{
		api.GET("/captcha", controllers.GetCaptcha)
	//api.GET("/email_code", controllers.GetEmailCode)
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.GET("/profile", middleware.AuthMiddleware(), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "已登录"})
		})
	}

	return r
}
