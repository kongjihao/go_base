package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.liwenzhou.com");
	if err != nil {
		fmt.Println("Error:", err)
		return 
	} else {
		fmt.Println("Response status:", resp.Status)
	}
	
}