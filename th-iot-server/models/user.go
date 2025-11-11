package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;comment:用户ID，自增主键"`
	Username  string         `gorm:"size:50;not null;unique;comment:用户名，唯一标识用户"`
	Password  string         `gorm:"size:255;not null;comment:用户密码（加密存储）"`
	Email     string         `gorm:"size:100;comment:用户邮箱，用于联系或找回密码"`
	CreatedAt time.Time      `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:逻辑删除标志，软删除"`
}
