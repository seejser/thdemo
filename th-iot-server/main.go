package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"th-iot-server/router"
	"th-iot-server/utils"
)

func init() {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ 没有找到 .env 文件，使用系统环境变量")
	}

	// 设置 Gin 模式
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
}
func main() {
	log.Println("🚀 Server starting...")

	// 初始化 MySQL
	utils.InitDB()
	defer utils.CloseDB()

	// 初始化 Redis
	if err := utils.InitRedis(); err != nil {
		log.Fatalf("Redis 初始化失败: %v", err)
	}
	defer utils.CloseRedis()

	// 初始化路由
	r := router.InitRouter()

	log.Println("✅ Server running at http://localhost:9090")
	if err := r.Run(":9090"); err != nil {
		log.Fatalf("Server 启动失败: %v", err)
	}
}
