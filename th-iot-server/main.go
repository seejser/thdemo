package main

import (
	"fmt"
	"th-iot-server/router"
	"th-iot-server/utils"
)

func main() {
	fmt.Println("🚀 Server starting...")

	utils.InitDB()

	r := router.InitRouter()

	fmt.Println("✅ Server running at http://localhost:9090")
	r.Run(":9090")
}
