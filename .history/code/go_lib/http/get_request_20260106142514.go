package main

import (
	"fmt"
	"net/http"
)

func main() {
	if _, err := http.Get("https://www.liwenzhou.com"); err != nil {
		return 
	}
	
}