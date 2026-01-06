package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// 关于GET请求的参数需要使用Go语言内置的net/url这个标准库来处理。
// 我们可以使用url.Values类型来构建查询参数，然后将其编码成URL格式的字符串，并将其附加到基础URL上。
func main() {
	baseURL := "http://localhost:8080/info"
	params := url.Values{}
	params.Add("name", "John")
	params.Add("age", "30")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Printf("Request baseURL error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response URL:", resp.Request.URL)
	fmt.Println("Response status code:", resp.StatusCode)
	
}

