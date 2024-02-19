package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DoRequest 函数用于执行 HTTP 请求并处理响应。
// 参数 method: 请求方法，例如 "GET"、"POST" 等。
// 参数 url: 请求的 URL。
// 参数 body: 请求体，通常是请求内容的数据流。
// 参数 ret: 用于存储响应数据的变量，通常是一个结构体指针。
// 参数 headers: 请求头信息，以键值对形式表示。
// 参数 bAuth: 基本认证信息，包括用户名和密码，如果不需要基本认证则为 nil。
// 函数返回一个 error，如果请求成功并且响应状态码为 200，则返回 nil；否则返回包含错误信息的 error。
func DoRequest(method, url string, body io.Reader, ret interface{}, headers map[string]string, bAuth []string) error {

	// 创建 HTTP 请求对象。
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	// 如果提供了基本认证信息，则设置请求的基本认证头部。
	if bAuth != nil && len(bAuth) == 2 {
		// 设置基本认证头部。
		req.SetBasicAuth(bAuth[0], bAuth[1])
	}

	// 如果提供了自定义请求头信息，则设置请求头。
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	// 发送 HTTP 请求并获取响应。
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取响应体的内容。
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// 检查响应状态码是否为 200，如果不是则返回错误信息。
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code: %d, body: %s", resp.StatusCode, bodyBytes)
	}

	// 解析响应体的 JSON 数据并存储到 ret 变量中。
	err = json.Unmarshal(bodyBytes, ret)
	if err != nil {
		return err
	}

	// 返回 nil 表示请求成功。
	return nil
}

// PostJson 函数用于执行 HTTP POST 请求，并以 JSON 格式发送数据。
// 参数 url: 请求的 URL。
// 参数 info: 要发送的数据，通常是一个结构体或其他数据类型。
// 参数 result: 用于存储响应数据的变量，通常是一个结构体指针。
// 参数 bAuth: 基本认证信息，包括用户名和密码，如果不需要基本认证则为 nil。
// 函数返回一个 error，如果请求成功并且响应状态码为 200，则返回 nil；否则返回包含错误信息的 error。
func PostJson(url string, info interface{}, result interface{}, bAuth []string) error {
	// 将 info 变量中的数据编码为 JSON 格式。
	contentJson, err := json.Marshal(info)
	if err != nil {
		return err
	}

	// 设置请求头，指定内容类型为 JSON 格式。
	headers := map[string]string{
		"Content-Type": "application/json;charset=utf-8",
	}

	// 调用 DoRequest 函数执行 HTTP POST 请求，发送 JSON 数据。
	return DoRequest(http.MethodPost, url, bytes.NewReader(contentJson), result, headers, bAuth)
}
