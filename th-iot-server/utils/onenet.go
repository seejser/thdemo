package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"th-iot-server/config"
)


var (
	ProductID  = config.OneNETProductID
	ProductKey = config.OneNETProductAccessKey
	Version    = config.OneNETVersion
	Method     = config.OneNETMethod
	OneNETBase = "https://iot-api.heclouds.com"
)
	

// ---------------------- 签名函数 ----------------------

func GenerateSign(params map[string]string, key string) (string, error) {
	fmt.Println("ProductID:", ProductID)
	fmt.Println("ProductKey:", ProductKey)
	fmt.Println("Version:", Version)
	fmt.Println("Method:", Method)
	fmt.Println("OneNETBase:", OneNETBase)
	strToSign := fmt.Sprintf("%s\n%s\n%s\n%s",
		params["et"], params["method"], params["res"], params["version"])

	decodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", fmt.Errorf("key base64 decode failed: %v", err)
	}

	var hmacBytes []byte
	switch strings.ToLower(params["method"]) {
	case "sha1":
		h := hmac.New(sha1.New, decodedKey)
		h.Write([]byte(strToSign))
		hmacBytes = h.Sum(nil)
	case "sha256":
		h := hmac.New(sha256.New, decodedKey)
		h.Write([]byte(strToSign))
		hmacBytes = h.Sum(nil)
	default:
		return "", fmt.Errorf("unsupported method: %s", params["method"])
	}

	return base64.StdEncoding.EncodeToString(hmacBytes), nil
}

// ---------------------- Token 生成 ----------------------

func GenerateToken(res string, expireTime int64) (string, error) {
	params := map[string]string{
		"version": Version,
		"res":     res,
		"et":      fmt.Sprintf("%d", expireTime),
		"method":  Method,
	}
	sign, err := GenerateSign(params, ProductKey)
	if err != nil {
		return "", err
	}
	token := fmt.Sprintf("version=%s&res=%s&et=%s&method=%s&sign=%s",
		url.QueryEscape(Version),
		url.QueryEscape(res),
		params["et"],
		url.QueryEscape(Method),
		url.QueryEscape(sign),
	)
	return token, nil
}

// ---------------------- HTTP 请求 ----------------------

func DoRequest(method, urlStr string, body []byte, token string, result interface{}) error {
	reqBody := bytes.NewReader([]byte{})
	if body != nil {
		reqBody = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, urlStr, reqBody)
	if err != nil {
		return fmt.Errorf("请求创建失败: %v", err)
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json, text/plain, */*")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求发送失败: %v", err)
	}
	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应体失败: %v", err)
	}

	var rawResp map[string]interface{}
	if err := json.Unmarshal(responseBytes, &rawResp); err != nil {
		fmt.Printf("HTTP 响应体 (非JSON): %s\n", string(responseBytes))
		return fmt.Errorf("解析响应 JSON 失败: %v", err)
	}

	if result != nil {
		data, _ := json.Marshal(rawResp)
		json.Unmarshal(data, result)
	}

	if code, ok := rawResp["code"].(float64); ok && code != 0 {
		msg := "Unknown Error"
		if m, ok := rawResp["msg"].(string); ok {
			msg = m
		}
		return fmt.Errorf("API 调用失败 (Code: %.0f, Msg: %s)", code, msg)
	}

	return nil
}

// ---------------------- 设备结构体 ----------------------

type Device struct {
	Did             string  `json:"did"`
	ProductID       string  `json:"pid"`
	Name            string  `json:"name"`
	Status          int     `json:"status"`
	Imei            string  `json:"imei"`
	AccessProtocol  int     `json:"access_pt"`
	DataProtocol    int     `json:"data_pt"`
	CreateTime      string  `json:"create_time"`
	ActivateTime    string  `json:"activate_time"`
	LastConnectTime string  `json:"last_connect_time"`
	Lat             string  `json:"lat"`
	Lon             string  `json:"lon"`
	EnableStatus    bool    `json:"enable_status"`
	Private         bool    `json:"private"`
	Obsv            bool    `json:"obsv"`
	ObsvSt          bool    `json:"obsv_st"`
	IntelligentWay  int     `json:"intelligent_way"`
}

type DeviceDetail struct {
	Name        string                 `json:"name"`
	ProductID   string                 `json:"product_id"`
	Status      int                    `json:"status"`
	Imei        string                 `json:"imei"`
	Properties  map[string]interface{} `json:"properties"`
	Desc        string                 `json:"desc"`
	AuthInfo    map[string]string      `json:"auth_info"`
	CreateTime  string                 `json:"create_time"`
	ActivateTime string                `json:"activate_time"`
	LastConnectTime string             `json:"last_connect_time"`
	Lat         string                 `json:"lat"`
	Lon         string                 `json:"lon"`
	EnableStatus bool                  `json:"enable_status"`
	Private      bool                  `json:"private"`
	Obsv         bool                  `json:"obsv"`
	ObsvSt       bool                  `json:"obsv_st"`
	IntelligentWay int                 `json:"intelligent_way"`
}

// ---------------------- 1. 获取设备列表 ----------------------

func GetDeviceList(offset, limit int) ([]Device, error) {
	urlStr := fmt.Sprintf("%s/device/list?product_id=%s&offset=%d&limit=%d", OneNETBase, ProductID, offset, limit)

	var res struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			List   []Device `json:"list"`
			Offset int      `json:"offset"`
			Limit  int      `json:"limit"`
		} `json:"data"`
	}

	expireTime := time.Now().Add(5 * time.Minute).Unix()
	token, err := GenerateToken(fmt.Sprintf("products/%s", ProductID), expireTime)
	if err != nil {
		return nil, fmt.Errorf("生成 token 失败: %v", err)
	}

	if err := DoRequest("GET", urlStr, nil, token, &res); err != nil {
		return nil, fmt.Errorf("获取设备列表失败: %v", err)
	}

	return res.Data.List, nil
}

// ---------------------- 2. 获取设备详情 ----------------------

func GetDeviceDetail(deviceName string) (*DeviceDetail, error) {
	urlStr := fmt.Sprintf("%s/device/detail?product_id=%s&device_name=%s", OneNETBase, ProductID, url.QueryEscape(deviceName))

	var res struct {
		Code int          `json:"code"`
		Msg  string       `json:"msg"`
		Data DeviceDetail `json:"data"`
	}

	expireTime := time.Now().Add(5 * time.Minute).Unix()
	token, err := GenerateToken(fmt.Sprintf("products/%s/devices/%s", ProductID, deviceName), expireTime)
	if err != nil {
		return nil, fmt.Errorf("生成 token 失败: %v", err)
	}

	if err := DoRequest("GET", urlStr, nil, token, &res); err != nil {
		return nil, fmt.Errorf("获取设备详情失败: %v", err)
	}

	return &res.Data, nil
}

// ---------------------- 3. 设置设备属性 ----------------------

type SetPropertyBody struct {
	ProductID  string                 `json:"product_id"`
	DeviceName string                 `json:"device_name"`
	Params     map[string]interface{} `json:"params"`
}

func SetDeviceProperty(deviceName string, properties map[string]interface{}) error {
	apiURL := fmt.Sprintf("%s/thingmodel/set-device-property", OneNETBase)

	bodyStruct := SetPropertyBody{
		ProductID:  ProductID,
		DeviceName: deviceName,
		Params:     properties,
	}

	jsonData, err := json.Marshal(bodyStruct)
	if err != nil {
		return fmt.Errorf("请求体序列化失败: %v", err)
	}

	expireTime := time.Now().Add(5 * time.Minute).Unix()
	token, err := GenerateToken(fmt.Sprintf("products/%s", ProductID), expireTime)
	if err != nil {
		return fmt.Errorf("生成 token 失败: %v", err)
	}

	var result map[string]interface{}
	if err := DoRequest("POST", apiURL, jsonData, token, &result); err != nil {
		return fmt.Errorf("物模型设置属性请求失败: %v", err)
	}

	if data, ok := result["data"].(map[string]interface{}); ok {
		if code, ok := data["code"].(float64); ok && code != 200 {
			msg := "Unknown Device Error"
			if m, ok := data["msg"].(string); ok {
				msg = m
			}
			return fmt.Errorf("设备端执行失败 (Code: %.0f, Msg: %s)", code, msg)
		}
	}

	return nil
}

// ---------------------- 4. 控制继电器 ----------------------

func ControlRelay(deviceName string, on bool) error {
	relayValue := 0
	if on {
		relayValue = 1
	}

	properties := map[string]interface{}{"relay": relayValue}
	return SetDeviceProperty(deviceName, properties)
}