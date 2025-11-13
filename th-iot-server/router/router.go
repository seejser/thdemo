package router

import (
	"github.com/gin-gonic/gin"
	"th-iot-server/controllers"
	"th-iot-server/middleware"
)

func InitRouter() *gin.Engine {
	// 默认带 Logger + Recovery 中间件
	r := gin.Default()

	// ✅ 全局中间件
	r.Use(
		middleware.CorsMiddleware(),     // 跨域
		middleware.ResponseMiddleware(), // 统一响应
	)

	// ✅ 路由分组
	api := r.Group("/api/v1")
	{
		// ---- 认证相关 ----
		auth := api.Group("/auth")
		{
			auth.GET("/captcha", controllers.GetCaptcha)
			auth.GET("/email_code", controllers.GetEmailCode)
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
			auth.POST("/refresh", controllers.RefreshToken)
		}

		// ---- 用户相关 ----
		user := api.Group("/user")
		user.Use(middleware.AuthMiddleware())
		{
			user.GET("/info", controllers.UserInfo)
		}

		// ---- 设备相关 ----
		device := api.Group("/device")
		//device.Use(middleware.AuthMiddleware()) // 需要登录
		{
			device.POST("", controllers.CreateDevice)       // 新增设备
			device.PUT("/:id", controllers.UpdateDevice)    // 更新设备
			device.DELETE("/:id", controllers.DeleteDevice) // 删除设备
			device.GET("/:id", controllers.GetDeviceByID)   // 获取设备详情
			device.GET("", controllers.ListDevices)         // 分页 + 关键字查询
			// 同步接口
			device.POST("/sync", controllers.SyncDeviceList)              // 同步所有设备
			device.POST("/sync/:device_id", controllers.SyncDeviceDetail) // 同步单个设备
		}
	}

	return r
}
