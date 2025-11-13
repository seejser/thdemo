package utils

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"th-iot-server/config"
)

var DB *gorm.DB

// InitDB 初始化 MySQL 数据库连接
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQLUser,
		config.MySQLPassword,
		config.MySQLAddr,
		config.MySQLDB,
	)

	// 自定义日志模式（可选 info / silent）
	newLogger := logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败: %v", err))
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("获取数据库底层连接失败: %v", err))
	}
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大连接数

	DB = db
	log.Println("✅ MySQL 数据库连接成功")
}

// CloseDB 关闭数据库连接
func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Printf("关闭数据库连接失败: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("关闭数据库连接失败: %v", err)
		} else {
			log.Println("✅ MySQL 数据库连接已关闭")
		}
	}
}
