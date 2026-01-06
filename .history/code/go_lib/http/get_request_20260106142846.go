package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.liwenzhou.com");
	if err != nil {
		fmt.Printf("Request %s error: %v\n", resp.Request.URL, err)
		return 
	} 
	
	
	fmt.Println("Response status:", resp.Status)

}