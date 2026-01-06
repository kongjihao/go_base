package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// POST请求
func main() {
	baseURL := "http://localhost:8080/post"

	// 构建POST请求的参数
	// 方式一：
	// params := url.Values{}
	// params.Add("name", "Alice")
	// params.Add("age", "25")

	// 方式二：
	data := `{"name": "Alice", "age": 25}`

	// 发送POST请求
	// 方式一：此种方式比第二种方式更简单，直接使用http.PostForm方法发送表单数据
	// resp, err := http.PostForm(baseURL, params)

	// 方式二：此种方式指定了Content-Type为application/json，并将数据作为字符串发送
	resp, err := http.Post(baseURL, "application/json", strings.NewReader(data))

	if err != nil {
		fmt.Printf("Request error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response URL:", resp.Request.URL)
	fmt.Println("Response status code:", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read response body error: %v\n", err)
		return
	}
	fmt.Println("Response body:", string(body))
}
