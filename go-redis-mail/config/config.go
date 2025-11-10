package config

var (
	RedisAddr     = "localhost:6379"
	RedisPassword = ""
	RedisDB       = 0

	SMTPHost = "smtp.qq.com"
	SMTPPort = 587
	SMTPUser = "1598253545@qq.com" // 发送邮箱
	SMTPPass = ""  // 邮箱授权码（不是密码）
)
