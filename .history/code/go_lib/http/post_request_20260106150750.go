package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// POST请求
func main() {
	baseURL := "http://localhost:8080/post"

	// 