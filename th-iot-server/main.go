package main

import (
    "fmt"
    "th-iot-server/middleware"
    "th-iot-server/router"
    "th-iot-server/utils"
)

func main() {
    fmt.Println("Hello, 世界")
    utils.InitDB()
    //utils.InitRedis()

    r := router.InitRouter()
    r.Use(middleware.ResponseMiddleware())
    fmt.Println("Server running at http://localhost:9090")
    r.Run(":9090")
}
