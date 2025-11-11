package config

var (
	// Redis 配置
	RedisAddr     = "localhost:6379"
	RedisPassword = ""
	RedisDB       = 0

	// MySQL 配置
	MySQLAddr     = "127.0.0.1:3306"
	MySQLUser     = "root"
	MySQLPassword = "aA123456"
	MySQLDB       = "thiotdb"

	// SMTP 配置
	SMTPHost = "smtp.qq.com"
	SMTPPort = 587
	SMTPUser = "1598253545@qq.com" // 发送邮箱
	SMTPPass = "nfuzpssryuheiebi"  // 邮箱授权码（不是密码）
)
