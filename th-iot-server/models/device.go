package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// JSONMap 用于 GORM 的 JSON 字段映射
type JSONMap map[string]interface{}

// Value 实现 driver.Valuer 接口
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	b, err := json.Marshal(j)
	return string(b), err
}

// Scan 实现 sql.Scanner 接口
func (j *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*j = make(JSONMap)
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid JSON data")
	}
	return json.Unmarshal(bytes, j)
}

// Device 设备模型
type Device struct {
	ID             uint64         `gorm:"primaryKey;autoIncrement;comment:主键ID，从1开始" json:"id"`
	DeviceID       string         `gorm:"size:64;unique;not null;comment:平台设备唯一标识（唯一）" json:"device_id"`
	Name           string         `gorm:"size:128;unique;not null;comment:设备名称（唯一）" json:"name"`
	Product        string         `gorm:"size:128;comment:所属产品型号" json:"product,omitempty"`
	Region         string         `gorm:"size:128;comment:地区/地理位置" json:"region,omitempty"`
	Description    string         `gorm:"type:text;comment:设备描述" json:"description,omitempty"`
	Extra          JSONMap        `gorm:"type:json;comment:扩展属性（如地理坐标、标签、自定义参数等）" json:"extra,omitempty"`
	Status         uint8          `gorm:"not null;default:2;comment:设备状态(1=在线,2=离线,3=故障,4=停用)" json:"status"`
	Switch         bool           `gorm:"not null;default:false;comment:主电源开关(1=开,0=关)" json:"switch"`
	Relay          bool           `gorm:"not null;default:false;comment:继电器开关(1=开,0=关)" json:"relay"`
	OutJ9          bool           `gorm:"not null;default:false;comment:J9继电器开关(1=开,0=关)" json:"out_j9"`
	Signal         *int           `gorm:"comment:信号质量(0-31) (映射 csq)" json:"signal,omitempty"`
	Temp           *float64       `gorm:"comment:实时温度(-55~125°C) (映射 temperature)" json:"temp,omitempty"`
	Warning        string         `gorm:"size:255;comment:预警事件(断电/过流等) (映射 alarm event)" json:"warning,omitempty"`
	CellInfo       string         `gorm:"type:text;comment:基站信息 (映射 cell_info)" json:"cell_info,omitempty"`
	SimNumber      string         `gorm:"size:32;comment:SIM卡号 (映射 imsi)" json:"sim_number,omitempty"`
	ReportCycle    *int           `gorm:"comment:上报周期(秒) (映射 interval)" json:"report_cycle,omitempty"`
	Mac            string         `gorm:"type:text;comment:MAC地址 (映射 macs)" json:"mac,omitempty"`
	IPAddress      string         `gorm:"size:64;comment:最近在线时的设备IP" json:"ip_address,omitempty"`
	FirmwareVersion string        `gorm:"size:64;comment:固件版本号" json:"firmware_version,omitempty"`
	LastOnlineAt   *time.Time     `gorm:"comment:最后一次上线时间" json:"last_online_at,omitempty"`

	CreatedAt      time.Time      `gorm:"autoCreateTime;comment:创建时间" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime;comment:更新时间" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index;comment:软删除时间（NULL表示未删除）" json:"deleted_at,omitempty"`
}
