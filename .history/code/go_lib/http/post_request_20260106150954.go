package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// POST请求
func main() {
	baseURL := "http://localhost:8080/post"

	// 构建POST请求的参数
	params := url.Values{}
	params.Add("name", "Alice")
	params.Add("age", "25")

	// 发送POST请求
	// resp, err := http.PostForm(baseURL, params)
	http.Post(baseURL, "application/json", pa)
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
