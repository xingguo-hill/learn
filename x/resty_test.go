package x

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestGet(t *testing.T) {
	response, err := resty.New().R().Get("https://www.baidu.com")
	// 检查请求是否出错
	if err != nil {
		fmt.Println("请求出错:", err)
		return
	}
	// 输出响应内容
	fmt.Println("状态码:", response.StatusCode())
	fmt.Println("响应体:", string(response.Body()))
}
func TestPost(t *testing.T) {
	// 创建一个 Resty 客户端
	client := resty.New()

	// 构建要发送的数据（这里使用了一个 JSON 示例）
	data := map[string]any{
		"title":  "v1",
		"body":   "v2",
		"userId": 1,
	}

	// 发送 POST 请求
	response, err := client.R().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetBody(data).
		Post("https://jsonplaceholder.typicode.com/posts")

	// 检查请求是否出错
	if err != nil {
		fmt.Println("请求出错:", err)
		return
	}

	// 输出响应内容
	fmt.Println("状态码:", response.StatusCode())
	fmt.Println("响应体:", string(response.Body()))
}
