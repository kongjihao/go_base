package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// POST请求
func main() {
	baseURL := "http://example.com/api"

	