package controllers

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
	"th-iot-server/middleware"
	"th-iot-server/models"
	"th-iot-server/services"
)

// CreateDevice 新增设备
func CreateDevice(c *gin.Context) {
	var req models.Device
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.ReturnError(c, http.StatusBadRequest, err)
		return
	}

	if err := services.DeviceService.Create(&req); err != nil {
		middleware.ReturnError(c, http.StatusInternalServerError, err)
		return
	}

	middleware.ReturnSuccess(c, gin.H{"message": "设备创建成功", "device": req})
}

// UpdateDevice 更新设备
func UpdateDevice(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		middleware.ReturnError(c, http.StatusBadRequest, err)
		return
	}

	var req models.Device
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.ReturnError(c, http.StatusBadRequest, err)
		return
	}

	req.ID = id
	if err := services.DeviceService.Update(&req); err != nil {
		middleware.ReturnError(c, http.StatusInternalServerError, err)
		return
	}

	middleware.ReturnSuccess(c, gin.H{"message": "设备更新成功", "device": req})
}

// DeleteDevice 删除设备
func DeleteDevice(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		middleware.ReturnError(c, http.StatusBadRequest, err)
		return
	}

	if err := services.DeviceService.Delete(id); err != nil {
		middleware.ReturnError(c, http.StatusInternalServerError, err)
		return
	}

	middleware.ReturnSuccess(c, gin.H{"message": "设备删除成功"})
}

// GetDeviceByID 获取单个设备详情
func GetDeviceByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		middleware.ReturnError(c, http.StatusBadRequest, err)
		return
	}

	device, err := services.DeviceService.GetByID(id)
	if err != nil {
		middleware.ReturnError(c, http.StatusNotFound, err)
		return
	}

	middleware.ReturnSuccess(c, gin.H{"device": device})
}

// ListDevices 分页 + 关键字查询设备列表
func ListDevices(c *gin.Context) {
	// 页码与分页大小
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	// 关键字搜索 (name / device_id / region / product)
	keyword := c.Query("keyword")

	list, total, err := services.DeviceService.List(page, limit, keyword)
	if err != nil {
		middleware.ReturnError(c, http.StatusInternalServerError, err)
		return
	}

	middleware.ReturnSuccess(c, gin.H{
		"list":  list,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
// SyncDeviceList 同步设备列表
func SyncDeviceList(c *gin.Context) {
	if err := services.DeviceService.SyncList(); err != nil {
		middleware.ReturnError(c, http.StatusInternalServerError, err)
		return
	}
	middleware.ReturnSuccess(c, gin.H{"message": "设备列表同步成功"})
}

// SyncDeviceDetail 同步单个设备详情
func SyncDeviceDetail(c *gin.Context) {
	deviceID := c.Param("device_id")
	if deviceID == "" {
		middleware.ReturnError(c, http.StatusBadRequest, fmt.Errorf("device_id 不能为空"))
		return
	}

	device, err := services.DeviceService.SyncDetail(deviceID)
	if err != nil {
		middleware.ReturnError(c, http.StatusInternalServerError, err)
		return
	}

	middleware.ReturnSuccess(c, gin.H{"device": device, "message": "设备详情同步成功"})
}
