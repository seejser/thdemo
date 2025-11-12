package main

import (
	"fmt"
	"th-iot-server/router"
	"th-iot-server/utils"
)

func main() {
	fmt.Println("🚀 Server starting...")

	utils.InitDB()
	// 初始化 Redis
	if err := utils.InitRedis(); err != nil {
		panic(err)
	}
	defer utils.CloseRedis()
	r := router.InitRouter()

	fmt.Println("✅ Server running at http://localhost:9090")
	r.Run(":9090")
}
