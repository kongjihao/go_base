package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.liwenzhou.com");
	if err != nil {
		fmt.Println("Request error:", err)
		return 
	} else {
		fmt.Println("Response status:", resp.Status)
	}
	
}