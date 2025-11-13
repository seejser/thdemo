package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	// 优先加载 .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env 文件未找到，使用默认值或系统环境变量")
	}

	// Redis
	RedisAddr = getEnv("REDIS_ADDR", RedisAddr)
	RedisPassword = getEnv("REDIS_PASSWORD", RedisPassword)
	RedisDB = getEnvAsInt("REDIS_DB", RedisDB)

	// MySQL
	MySQLAddr = getEnv("MYSQL_ADDR", MySQLAddr)
	MySQLUser = getEnv("MYSQL_USER", MySQLUser)
	MySQLPassword = getEnv("MYSQL_PASSWORD", MySQLPassword)
	MySQLDB = getEnv("MYSQL_DB", MySQLDB)

	// SMTP
	SMTPHost = getEnv("SMTP_HOST", SMTPHost)
	SMTPPort = getEnvAsInt("SMTP_PORT", SMTPPort)
	SMTPUser = getEnv("SMTP_USER", SMTPUser)
	SMTPPass = getEnv("SMTP_PASS", SMTPPass)

	// OneNET
	OneNETProductID = getEnv("ONENET_PRODUCT_ID", OneNETProductID)
	OneNETProductAccessKey = getEnv("ONENET_ACCESS_KEY", OneNETProductAccessKey)
	OneNETVersion = getEnv("ONENET_VERSION", OneNETVersion)
	OneNETMethod = getEnv("ONENET_METHOD", OneNETMethod)
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return defaultVal
}

// 默认配置值
var (
	RedisAddr     = "localhost:6379"
	RedisPassword = ""
	RedisDB       = 0

	MySQLAddr     = "127.0.0.1:3306"
	MySQLUser     = "root"
	MySQLPassword = "aA123456"
	MySQLDB       = "thiotdb"

	SMTPHost = "smtp.qq.com"
	SMTPPort = 587
	SMTPUser = "1598253545@qq.com"
	SMTPPass = "nfuzpssryuheiebi"

	OneNETProductID        = "Ay3w00GD25"
	OneNETProductAccessKey = "w7G5OVd5u9/BD+l/42FtbYcJe9d362EvJaFbWY0nHcU="
	OneNETVersion          = "2022-05-01"
	OneNETMethod           = "sha1"
)
