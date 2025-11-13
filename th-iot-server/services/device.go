package services

import (
	"strings"
	"fmt"

	"th-iot-server/dao"
	"th-iot-server/models"
	"th-iot-server/utils"
)

type deviceService struct{}

var DeviceService = &deviceService{}

// Create 新增设备
func (s *deviceService) Create(device *models.Device) error {
	return dao.Device.Create(device)
}

// Update 更新设备
func (s *deviceService) Update(device *models.Device) error {
	return dao.Device.Update(device)
}

// Delete 删除设备
func (s *deviceService) Delete(id uint64) error {  // 改为 uint64
	dev, err := dao.Device.FindByID(id)
	if err != nil {
		return err
	}
	return dao.Device.Delete(dev)
}

// GetByID 获取单个设备
func (s *deviceService) GetByID(id uint64) (*models.Device, error) {  // 改为 uint64
	return dao.Device.FindByID(id)
}

// List 分页 + 关键字查询
func (s *deviceService) List(page, limit int, keyword string) ([]models.Device, int64, error) {
	offset := (page - 1) * limit

	// 支持模糊搜索
	var list []models.Device
	var total int64
	var err error

	if keyword != "" {
		keyword = strings.TrimSpace(keyword)
		list, total, err = dao.Device.FindByKeyword(keyword, offset, limit)
	} else {
		list, total, err = dao.Device.FindAllPaginate(offset, limit)
	}

	return list, total, err
}
// SyncList 从 OneNET 拉取设备列表并保存到本地数据库
func (s *deviceService) SyncList() error {
	// 调用 utils 获取设备列表
	devicesData, err := utils.GetDeviceList()
	if err != nil {
		return fmt.Errorf("获取 OneNET 设备列表失败: %w", err)
	}

	for _, d := range devicesData {
		// 构建 Device 对象
		device := &models.Device{
			DeviceID: fmt.Sprintf("%v", d["device_id"]),
			Name:     fmt.Sprintf("%v", d["name"]),
			Product:  fmt.Sprintf("%v", d["product_id"]),
			Region:   fmt.Sprintf("%v", d["region"]),
			Status:   2, // 默认离线
		}

		// 判断是否存在
		exist, _ := dao.Device.FindByDeviceID(device.DeviceID)
		if exist != nil {
			device.ID = exist.ID
			if err := dao.Device.Update(device); err != nil {
				fmt.Println("更新设备失败:", device.DeviceID, err)
			}
		} else {
			if err := dao.Device.Create(device); err != nil {
				fmt.Println("创建设备失败:", device.DeviceID, err)
			}
		}
	}

	return nil
}

// SyncDetail 拉取单个设备详情并更新本地
func (s *deviceService) SyncDetail(deviceID string) (*models.Device, error) {
	data, err := utils.GetDeviceDetail(deviceID)
	if err != nil {
		return nil, fmt.Errorf("获取设备详情失败: %w", err)
	}

	device := &models.Device{
		DeviceID: deviceID,
		Name:     fmt.Sprintf("%v", data["name"]),
		Product:  fmt.Sprintf("%v", data["product_id"]),
		Region:   fmt.Sprintf("%v", data["region"]),
		Status:   2, // 默认离线
	}

	exist, _ := dao.Device.FindByDeviceID(deviceID)
	if exist != nil {
		device.ID = exist.ID
		if err := dao.Device.Update(device); err != nil {
			return nil, err
		}
	} else {
		if err := dao.Device.Create(device); err != nil {
			return nil, err
		}
	}

	return device, nil
}