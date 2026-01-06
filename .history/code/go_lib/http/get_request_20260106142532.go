package main

import (
	"fmt"
	"net/http"
)

func main() {
	if resp, err := http.Get("https://www.liwenzhou.com"); err != nil {
		return 
	} else {
		fmt.Println("Response status:", resp.Status)
	}
	
}