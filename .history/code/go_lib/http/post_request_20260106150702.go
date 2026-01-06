package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// POST请求的参数需要使用Go语言内置的net/url这个标准库来处理。
func main() {
	baseURL := "http://example.com/api"

	