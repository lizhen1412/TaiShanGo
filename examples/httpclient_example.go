package main

import (
	"fmt"

	"github.com/lizhen1412/TaiShanGo/httpclient"
)

type RequestBody struct {
	DeviceID    string `json:"deviceId"`
	AccessToken string `json:"accessToken"`
	ServerID    string `json:"serverID"`
	AgentID     string `json:"agentID"`
	Account     string `json:"account"`
}

type ResponseBody struct {
	Platform string `json:"platform"`
	Channel  string `json:"channel"`
	DeviceID string `json:"deviceId"`
	UserID   string `json:"userId"`
	ErrMsg   string `json:"errMsg"`
	Code     int    `json:"code"`
}

func main() {
	url := "http://127.0.0.1:8081/auth"

	requestBody := RequestBody{
		DeviceID:    "be70a492ea6e446a72a04de79a8bae84decec40e",
		AccessToken: "V0ZVL0lzd21VTllQVHcrdjR2M3BYMmJIcTRGQ2QvQVVXUnNvcDZQRWlLeEw4Um5pVVgxVmtmaUl4VHhCdnQzWUdaaTJoU0dxZ0ZLaDBYcDE4SnR6K3BIdXhCSDZlNld3TllRLzM0c1FLd2kwR2xUZkJjQUoyL3l5YTUydlVFWVNIVmtRZ0l1ZmdRb2pOWlZIQmJ3cDJheWYwWUhiakhUVVZLSi9zNkE1cGxrPQ==",
		ServerID:    "1001",
		AgentID:     "",
		Account:     "1708243878823",
	}

	var responseBody ResponseBody

	err := httpclient.PostJson(url, requestBody, &responseBody, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Response: %+v\n", responseBody)
}
