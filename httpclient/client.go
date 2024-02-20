package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// defaultClient 是默认的 HTTP 客户端，具有 3 秒的超时时间。
var defaultClient = &http.Client{
	Timeout: 3 * time.Second, // 设置默认的超时时间为 3 秒
}

// ClientConfig 结构体用于配置HTTP客户端
type ClientConfig struct {
	Timeout           time.Duration // 请求超时时间
	DisableKeepAlives bool          // 是否禁用连接池
}

// NewCustomClient 根据提供的 ClientConfig 创建并返回一个定制的 HTTP 客户端。
// 参数:
//
//	config: ClientConfig 结构体，用于配置 HTTP 客户端的参数。
//
// 返回值:
//
//	*http.Client: 创建的定制化的 HTTP 客户端。
func NewCustomClient(config ClientConfig) *http.Client {
	// 创建一个新的 HTTP 客户端，并根据提供的配置进行定制化设置
	return &http.Client{
		Timeout: config.Timeout, // 设置超时时间为配置中指定的超时时间
		Transport: &http.Transport{
			DisableKeepAlives: config.DisableKeepAlives, // 根据配置禁用或启用连接池
		},
	}
}

// PostJson 执行 HTTP POST 请求，以 JSON 格式发送数据，并解析响应到提供的 result 参数中。
// 参数:
//
//	url: 请求的 URL 地址。
//	info: 要发送的数据，通常是一个结构体或者一个指针。
//	result: 用于存储响应数据的接收器，通常是一个指向结构体的指针。
//	bAuth: 包含用户名和密码的切片，用于 HTTP 基本认证。
//	client: 执行请求的 HTTP 客户端，如果未提供，则使用默认客户端。
//
// 返回值:
//
//	error: 如果执行过程中发生错误，将返回相应的错误信息；否则返回 nil。
func PostJson(url string, info interface{}, result interface{}, bAuth []string, client *http.Client) error {
	// 如果未提供客户端，则使用默认客户端
	if client == nil {
		client = defaultClient
	}

	// 将 info 转换为 JSON 格式的字节切片
	contentJson, err := json.Marshal(info)
	if err != nil {
		return err
	}

	// 设置请求头
	headers := map[string]string{
		"Content-Type": "application/json", // 设置内容类型为 JSON
	}

	// 调用 DoRequest 函数执行 HTTP 请求，并将结果存储在提供的 result 参数中
	return DoRequest(http.MethodPost, url, bytes.NewReader(contentJson), result, headers, bAuth, client)
}

// DoRequest 执行 HTTP 请求并处理响应。
// 参数:
//
//	method: 请求方法，例如 GET、POST 等。
//	url: 请求的 URL 地址。
//	body: 请求的内容体，通常是一个 io.Reader。
//	ret: 用于存储响应数据的接收器，通常是一个指向结构体的指针。
//	headers: 请求的头部信息，一个映射类型，存储了键值对形式的头部信息。
//	bAuth: 包含用户名和密码的切片，用于 HTTP 基本认证。
//	client: 执行请求的 HTTP 客户端。
//
// 返回值:
//
//	error: 如果执行过程中发生错误，将返回相应的错误信息；否则返回 nil。
func DoRequest(method, url string, body io.Reader, ret interface{}, headers map[string]string, bAuth []string, client *http.Client) error {
	// 创建一个新的 HTTP 请求
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	// 如果提供了 Basic Auth 信息，则设置请求的基本认证信息
	if bAuth != nil && len(bAuth) == 2 {
		req.SetBasicAuth(bAuth[0], bAuth[1])
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 使用提供的客户端执行 HTTP 请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查响应状态码，如果不是 200 OK，则返回错误信息
	if resp.StatusCode != http.StatusOK {
		// 读取响应体内容
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		// 返回状态码和响应体内容组成的错误信息
		return fmt.Errorf("status code: %d, body: %s", resp.StatusCode, bodyBytes)
	}

	return json.NewDecoder(resp.Body).Decode(ret)
}
