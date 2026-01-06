package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.liwenzhou.com")
	if err != nil {
		fmt.Printf("Request %s error: %v\n", resp.Request.URL, err)
		return
	}

	defer resp.Body.Close() // 记得一定要
	fmt.Printf("Request URL: %s\n", resp.Request.URL)
	fmt.Println("Response status code:", resp.StatusCode)
	respBody, err := ioutil.ReadAll(resp.Body) // 返回的是liwenzhou网站的HTML内容
	if err != nil {
		fmt.Printf("Read response body error: %v\n", err)
		return
	}
	fmt.Println("Response body:", string(respBody))
}
