package main

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
)

// ---------------------- 签名函数 ----------------------

// GenerateSign 根据 OneNET 签名算法生成 sign
func GenerateSign(params map[string]string, key string) (string, error) {
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

// GenerateToken 生成产品级或设备级 token
func GenerateToken(res, key, version, method string, expireTime int64) (string, error) {
	params := map[string]string{
		"version": version,
		"res":     res,
		"et":      fmt.Sprintf("%d", expireTime),
		"method":  method,
	}
	sign, err := GenerateSign(params, key)
	if err != nil {
		return "", err
	}
	token := fmt.Sprintf("version=%s&res=%s&et=%s&method=%s&sign=%s",
		url.QueryEscape(version),
		url.QueryEscape(res),
		params["et"],
		url.QueryEscape(method),
		url.QueryEscape(sign),
	)
	return token, nil
}

// ---------------------- HTTP 请求 ----------------------

func DoRequest(method, urlStr string, body []byte, token string, result interface{}) error {
	fmt.Println("HTTP 请求信息:")
	fmt.Println("URL:", urlStr)
	fmt.Println("Method:", method)
	fmt.Println("Body:", string(body))
	fmt.Println("Authorization:", token)
	fmt.Println("-----------------------------------")

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

	fmt.Println("HTTP 响应 JSON:")
	rawJSON, _ := json.MarshalIndent(rawResp, "", "  ")
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

// ---------------------- 1. 获取设备列表 ----------------------

func GetDeviceList(productID, productKey, version, method string) {
	fmt.Println("\n=======================================================")
	fmt.Println("--- 获取设备列表 ---")
	urlStr := fmt.Sprintf("https://iot-api.heclouds.com/device/list?product_id=%s&offset=0&limit=100", productID)

	var res struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			List []struct {
				Name   string `json:"name"`
				Status int    `json:"status"`
				Imei   string `json:"imei"`
			} `json:"list"`
		} `json:"data"`
	}

	expireTime := time.Now().Add(5 * time.Minute).Unix()
	token, err := GenerateToken(fmt.Sprintf("products/%s", productID), productKey, version, method, expireTime)
	if err != nil {
		fmt.Println("生成 token 失败:", err)
		return
	}

	if err := DoRequest("GET", urlStr, nil, token, &res); err != nil {
		fmt.Println("获取设备列表失败:", err)
		return
	}

	if res.Data.List != nil {
		fmt.Println("设备列表:")
		for _, d := range res.Data.List {
			fmt.Printf("  设备名: %s, 状态: %d, IMEI: %s\n", d.Name, d.Status, d.Imei)
		}
	} else {
		fmt.Println("设备列表为空或解析失败。")
	}
}

// ---------------------- 2. 获取设备详情 ----------------------

func GetDeviceDetail(productID, productKey, deviceName, version, method string) {
	fmt.Println("\n=======================================================")
	fmt.Printf("--- 获取设备详情 (%s) ---\n", deviceName)
	urlStr := fmt.Sprintf("https://iot-api.heclouds.com/device/detail?product_id=%s&device_name=%s", productID, deviceName)

	var result map[string]interface{}
	expireTime := time.Now().Add(5 * time.Minute).Unix()
	token, err := GenerateToken(fmt.Sprintf("products/%s", productID), productKey, version, method, expireTime)
	if err != nil {
		fmt.Println("生成 token 失败:", err)
		return
	}

	if err := DoRequest("GET", urlStr, nil, token, &result); err != nil {
		fmt.Println("获取设备详情失败:", err)
		return
	}

	fmt.Println("设备详情获取完成。")
}

// ---------------------- 3. 设置设备属性 ----------------------

type SetPropertyBody struct {
	ProductID  string                 `json:"product_id"`
	DeviceName string                 `json:"device_name"`
	Params     map[string]interface{} `json:"params"`
}

func SetDeviceProperty(productID, productKey, deviceName string, properties map[string]interface{}, version, method string) error {
	const apiURL = "https://iot-api.heclouds.com/thingmodel/set-device-property"

	bodyStruct := SetPropertyBody{
		ProductID:  productID,
		DeviceName: deviceName,
		Params:     properties,
	}

	jsonData, err := json.Marshal(bodyStruct)
	if err != nil {
		return fmt.Errorf("请求体序列化失败: %v", err)
	}

	expireTime := time.Now().Add(5 * time.Minute).Unix()
	token, err := GenerateToken(fmt.Sprintf("products/%s", productID), productKey, version, method, expireTime)
	if err != nil {
		return fmt.Errorf("生成 token 失败: %v", err)
	}

	var result map[string]interface{}
	if err := DoRequest("POST", apiURL, jsonData, token, &result); err != nil {
		return fmt.Errorf("物模型设置属性请求失败: %v", err)
	}

	if data, ok := result["data"].(map[string]interface{}); ok {
		if deviceCode, ok := data["code"].(float64); ok && deviceCode != 200 {
			deviceMsg := "Unknown Device Error"
			if m, ok := data["msg"].(string); ok {
				deviceMsg = m
			}
			return fmt.Errorf("设备端执行失败 (Code: %.0f, Msg: %s)", deviceCode, deviceMsg)
		}
	}

	return nil
}

// 控制继电器
func ControlRelay(productID, productKey, deviceName string, on bool, version, method string) {
	action := "关闭"
	relayValue := 0
	if on {
		action = "打开"
		relayValue = 1
	}
	fmt.Printf("\n尝试通过物模型 API %s继电器 (relay=%d)...\n", action, relayValue)

	properties := map[string]interface{}{"relay": relayValue}
	if err := SetDeviceProperty(productID, productKey, deviceName, properties, version, method); err != nil {
		fmt.Printf("【失败】%s继电器失败: %v\n", action, err)
	} else {
		fmt.Printf("【成功】通过物模型 API %s继电器。\n", action)
	}
}

// ---------------------- 主函数 ----------------------

func main() {
	productID := "Ay3w00GD25"
	productKey := "w7G5OVd5u9/BD+l/42FtbYcJe9d362EvJaFbWY0nHcU="
	deviceName := "866560088599200"
	version := "2022-05-01"
	method := "sha1"

	GetDeviceList(productID, productKey, version, method)
	GetDeviceDetail(productID, productKey, deviceName, version, method)
	//ControlRelay(productID, productKey, deviceName, true, version, method)
	time.Sleep(2 * time.Second)
	ControlRelay(productID, productKey, deviceName, false, version, method)
}
