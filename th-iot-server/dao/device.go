package dao

import (
	"fmt"
	"strings"
	"th-iot-server/models"
	"th-iot-server/utils"
)

type DeviceDAO struct{}

var Device = &DeviceDAO{}

// Create 新增设备
func (d *DeviceDAO) Create(device *models.Device) error {
	return utils.DB.Create(device).Error
}

// Update 更新设备
func (d *DeviceDAO) Update(device *models.Device) error {
	return utils.DB.Save(device).Error
}

// Delete 删除设备
func (d *DeviceDAO) Delete(device *models.Device) error {
	return utils.DB.Delete(device).Error
}

// FindByID 查询单个设备
func (d *DeviceDAO) FindByID(id uint64) (*models.Device, error) {
	var device models.Device
	if err := utils.DB.First(&device, id).Error; err != nil {
		return nil, err
	}
	return &device, nil
}


// FindAll 查询所有设备
func (d *DeviceDAO) FindAll() ([]models.Device, error) {
	var list []models.Device
	if err := utils.DB.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// FindAllPaginate 分页查询，不带关键字
func (d *DeviceDAO) FindAllPaginate(offset, limit int) ([]models.Device, int64, error) {
	var list []models.Device
	var total int64

	if err := utils.DB.Model(&models.Device{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := utils.DB.Offset(offset).Limit(limit).Order("id DESC").Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// FindByKeyword 支持 name / device_id / region / product 模糊搜索
func (d *DeviceDAO) FindByKeyword(keyword string, offset, limit int) ([]models.Device, int64, error) {
	var list []models.Device
	var total int64
	kw := "%" + strings.TrimSpace(keyword) + "%"

	query := utils.DB.Model(&models.Device{}).Where(
		utils.DB.Where("name LIKE ?", kw).
			Or("device_id LIKE ?", kw).
			Or("region LIKE ?", kw).
			Or("product LIKE ?", kw),
	)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("统计总数失败: %w", err)
	}

	if err := query.Offset(offset).Limit(limit).Order("id DESC").Find(&list).Error; err != nil {
		return nil, 0, fmt.Errorf("查询列表失败: %w", err)
	}

	return list, total, nil
}
func (d *DeviceDAO) FindByDeviceID(deviceID string) (*models.Device, error) {
	var device models.Device
	if err := utils.DB.Where("device_id = ?", deviceID).First(&device).Error; err != nil {
		return nil, err
	}
	return &device, nil
}