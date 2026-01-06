package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.liwenzhou.com");
	if err != nil {
		fmt.Printf("Request %S error: %v\n", err)
		return 
	} 
	
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

}