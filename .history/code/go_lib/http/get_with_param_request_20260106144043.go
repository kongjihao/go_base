package main

import (
	"fmt"
	"net/http"
	"net/url"
)

//关于GET请求的参数需要使用Go语言内置的net/url这个标准库来处理。
func main() {
	baseURL := "http://example.com/api"
	params := url.Values{}
	params.Add("name", "John")
	params.Add("age", "30")

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}

