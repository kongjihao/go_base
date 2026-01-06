package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.liwenzhou.com");
	fmt.Printf("Request URL: %s\n", resp.Request.URL)
	if err != nil {
		fmt.Printf("Request %s error: %v\n", resp.Request.URL, err)
		return 
	} 
	
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

}