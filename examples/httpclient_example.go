package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lizhen1412/TaiShanGo/httpclient"
)

// RequestBody 结构体用于定义请求体的结构。
type RequestBody struct {
	DeviceID    string `json:"deviceId"`    // 设备ID
	AccessToken string `json:"accessToken"` // 访问令牌
	ServerID    string `json:"serverID"`    // 服务器ID
	AgentID     string `json:"agentID"`     // 代理ID
	Account     string `json:"account"`     // 账号
}

// ResponseBody 结构体用于定义响应体的结构。
type ResponseBody struct {
	Platform string `json:"platform"` // 平台
	Channel  string `json:"channel"`  // 渠道
	DeviceID string `json:"deviceId"` // 设备ID
	UserID   string `json:"userId"`   // 用户ID
	ErrMsg   string `json:"errMsg"`   // 错误消息
	Code     int    `json:"code"`     // 响应码
}

func main() {
	// 定义请求的URL
	url := "http://192.168.0.233:8081/auth"

	client := &http.Client{
		Timeout: 3 * time.Second, // 设置3秒超时
	}

	// 构造请求体数据
	requestBody := RequestBody{
		DeviceID:    "be70a492ea6e446a72a04de79a8bae84decec40e",                                                                                                                                                                                                 // 设备ID
		AccessToken: "V0ZVL0lzd21VTllQVHcrdjR2M3BYMmJIcTRGQ2QvQVVXUnNvcDZQRWlLeEw4Um5pVVgxVmtmaUl4VHhCdnQzWUdaaTJoU0dxZ0ZLaDBYcDE4SnR6K3BIdXhCSDZlNld3TllRLzM0c1FLd2kwR2xUZkJjQUoyL3l5YTUydlVFWVNIVmtRZ0l1ZmdRb2pOWlZIQmJ3cDJheWYwWUhiakhUVVZLSi9zNkE1cGxrPQ==", // 访问令牌
		ServerID:    "1001",                                                                                                                                                                                                                                     // 服务器ID
		AgentID:     "",                                                                                                                                                                                                                                         // 代理ID
		Account:     "1708243878823",                                                                                                                                                                                                                            // 账号
	}

	// 创建用于存储响应数据的变量
	var responseBody ResponseBody

	// 发送POST请求并处理响应
	err := httpclient.PostJson(url, requestBody, &responseBody, nil, client)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印响应数据
	fmt.Printf("Response: %+v\n", responseBody)
}
