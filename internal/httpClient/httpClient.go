package httpClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Requester 接口定义
type Requester interface {
	Get(url string, headers map[string]interface{}) ([]byte, error)
	Post(url string, headers map[string]interface{}, body interface{}) ([]byte, error)
	GetJSON(url string, headers map[string]interface{}) (map[string]interface{}, error)
	PostJSON(url string, headers map[string]interface{}, body interface{}) (map[string]interface{}, error)
}

// HttpClient 结构体
type HttpClient struct {
	client *http.Client
}

// NewHttpClient 创建一个新的 HttpClient 实例
func NewHttpClient(timeout time.Duration) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get 方法实现
func (h *HttpClient) Get(url string, headers map[string]interface{}) ([]byte, error) {
	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value.(string))
	}

	// 发送请求
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

// GetJSON 方法实现
func (h *HttpClient) GetJSON(url string, headers map[string]interface{}) (map[string]interface{}, error) {
	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value.(string))
	}

	// 发送请求
	resp, err := h.client.Do(req)
	if err != nil {
		log.Fatal("failed to send GET request: " + err.Error())
		return nil, fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	// 将 JSON 数据解析为 map
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return result, nil
}

// Post 方法实现
func (h *HttpClient) Post(url string, headers map[string]interface{}, body interface{}) ([]byte, error) {
	// 将 body 转换为 JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %w", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value.(string))
	}

	// 发送请求
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return responseBody, nil
}

// Post 方法实现
func (h *HttpClient) PostJSON(url string, headers map[string]interface{}, body interface{}) (map[string]interface{}, error) {
	// 将 body 转换为 JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %w", err)
	}

	// 设置请求头
	//req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value.(string))
	}

	// 发送请求
	resp, err := h.client.Do(req)
	if err != nil {
		log.Fatal("failed to send POST request: " + err.Error())
		return nil, fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	// 将 JSON 数据解析为 map
	var result map[string]interface{}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return result, nil
}

// 示例使用
func main() {
	// 创建 HttpClient 实例
	client := NewHttpClient(10 * time.Second)

	// 使用 GET 请求
	url := "https://jsonplaceholder.typicode.com/posts/1"
	headers := map[string]interface{}{
		"Accept": "application/json",
	}

	response, err := client.Get(url, headers)
	if err != nil {
		fmt.Println("GET request failed:", err)
		return
	}
	fmt.Println("GET response:", string(response))

	// 使用 POST 请求
	postUrl := "https://jsonplaceholder.typicode.com/posts"
	postHeaders := map[string]interface{}{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}
	postBody := map[string]interface{}{
		"title":  "foo",
		"body":   "bar",
		"userId": 1,
	}

	postResponse, err := client.Post(postUrl, postHeaders, postBody)
	if err != nil {
		fmt.Println("POST request failed:", err)
		return
	}
	fmt.Println("POST response:", string(postResponse))
}
