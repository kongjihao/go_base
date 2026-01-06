package main

import (
	"fmt"
	"net/http"
	"net/url"
)

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

