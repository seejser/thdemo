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

// ---------------------- 签名服务 ----------------------

// generateSign 根据 OneNET 签名算法生成 sign
func generateSign(params map[string]string, key string) (string, error) {
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

// GenerateToken 生成 OneNET Token
// key 可选，传空则使用产品级 AccessKey
func GenerateToken(res string, expireTime int64, key string) (string, error) {
	if key == "" {
		key = config.OneNETProductAccessKey
	}

	params := map[string]string{
		"version": config.OneNETVersion,
		"res":     res,
		"et":      fmt.Sprintf("%d", expireTime),
		"method":  config.OneNETMethod,
	}

	sign, err := generateSign(params, key)
	if err != nil {
		return "", err
	}

	token := fmt.Sprintf("version=%s&res=%s&et=%s&method=%s&sign=%s",
		url.QueryEscape(params["version"]),
		url.QueryEscape(params["res"]),
		params["et"],
		url.QueryEscape(params["method"]),
		url.QueryEscape(sign),
	)
	return token, nil
}

// ---------------------- HTTP 请求封装 ----------------------

func doRequest(method, urlStr string, body []byte, token string, result interface{}) error {
	fmt.Println("HTTP 请求信息:")
	fmt.Println("URL:", urlStr)
	fmt.Println("Method:", method)
	fmt.Println("Body:", string(body))
	fmt.Println("Authorization:", token)
	fmt.Println("-----------------------------------")

	var reqBody *bytes.Reader
	if body != nil {
		reqBody = bytes.NewReader(body)
	} else {
		reqBody = bytes.NewReader([]byte{})
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

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应体失败: %v", err)
	}

	var rawResp map[string]interface{}
	if err := json.Unmarshal(respBytes, &rawResp); err != nil {
		fmt.Printf("HTTP 响应体 (非JSON): %s\n", string(respBytes))
		return fmt.Errorf("解析响应 JSON 失败: %v", err)
	}

	rawJSON, _ := json.MarshalIndent(rawResp, "", "  ")
	fmt.Println("HTTP 响应 JSON:")
	fmt.Println(string(rawJSON))
	fmt.Println("-----------------------------------")

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

// ---------------------- OneNET 功能函数 ----------------------

// GetDeviceList 获取设备列表 (产品级 Token)
func GetDeviceList() ([]map[string]interface{}, error) {
	urlStr := fmt.Sprintf("https://iot-api.heclouds.com/device/list?product_id=%s&offset=0&limit=100", config.OneNETProductID)
	expireTime := time.Now().Add(5 * time.Minute).Unix()
	token, err := GenerateToken(fmt.Sprintf("products/%s", config.OneNETProductID), expireTime, "")
	if err != nil {
		return nil, fmt.Errorf("生成 token 失败: %v", err)
	}

	var res struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			List []map[string]interface{} `json:"list"`
		} `json:"data"`
	}

	if err := doRequest("GET", urlStr, nil, token, &res); err != nil {
		return nil, err
	}

	return res.Data.List, nil
}

// GetDeviceDetail 获取单个设备详情 (产品级 Token)
func GetDeviceDetail(deviceName string) (map[string]interface{}, error) {
	urlStr := fmt.Sprintf("https://iot-api.heclouds.com/device/detail?product_id=%s&device_name=%s", config.OneNETProductID, deviceName)
	expireTime := time.Now().Add(5 * time.Minute).Unix()
	token, err := GenerateToken(fmt.Sprintf("products/%s", config.OneNETProductID), expireTime, "")
	if err != nil {
		return nil, fmt.Errorf("生成 token 失败: %v", err)
	}

	var result map[string]interface{}
	if err := doRequest("GET", urlStr, nil, token, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SetDeviceProperty 下发物模型属性 (产品级 Token)
func SetDeviceProperty(deviceName string, properties map[string]interface{}) error {
	const apiURL = "https://iot-api.heclouds.com/thingmodel/set-device-property"

	bodyStruct := map[string]interface{}{
		"product_id":  config.OneNETProductID,
		"device_name": deviceName,
		"params":      properties,
	}

	jsonData, err := json.Marshal(bodyStruct)
	if err != nil {
		return fmt.Errorf("请求体序列化失败: %v", err)
	}

	expireTime := time.Now().Add(5 * time.Minute).Unix()
	res := fmt.Sprintf("products/%s", config.OneNETProductID)
	token, err := GenerateToken(res, expireTime, "")
	if err != nil {
		return fmt.Errorf("生成产品 Token 失败: %v", err)
	}

	var result map[string]interface{}
	if err := doRequest("POST", apiURL, jsonData, token, &result); err != nil {
		return err
	}

	// 检查设备端返回的 code
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

// ControlRelay 设置继电器开关
func ControlRelay(deviceName string, on bool) error {
	val := 0
	action := "关闭"
	if on {
		val = 1
		action = "打开"
	}

	fmt.Printf("尝试 %s设备 %s继电器...\n", deviceName, action)
	return SetDeviceProperty(deviceName, map[string]interface{}{"relay": val})
}
