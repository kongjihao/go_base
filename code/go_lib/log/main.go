package main

import (
	"log"
	"net/http"
	"os"
)

// Go语言提供的基本日志功能。Go语言提供的默认日志包是https://golang.org/pkg/log/。
// 优势：
// 它最大的优点是使用非常简单。我们可以设置任何io.Writer作为日志记录输出并向其发送要写入的日志。
//
// 劣势：
// 仅限基本的日志级别
// 只有一个Print选项。不支持INFO/DEBUG等多个级别。
// 对于错误日志，它有Fatal和Panic
// Fatal日志通过调用os.Exit(1)来结束程序
// Panic日志在写入日志消息之后抛出一个panic
// 但是它缺少一个ERROR日志级别，这个级别可以在不抛出panic或退出程序的情况下记录错误
// 缺乏日志格式化的能力——例如记录调用者的函数名和行号，格式化日期和时间格式。等等。
// 不提供日志切割的能力。

func main() {
	setLogger()
	simpleHttpGet01("http://www.google.com")
	simpleHttpGet01("http://www.baidu.com")
	simpleHttpGet01("www.google.com") // 故意写错的URL，触发错误日志
}

// 实现一个Go语言中的日志记录器非常简单——创建一个新的日志文件，然后设置它为日志的输出位置。
func setLogger() {
	// 日志打印到console
	// log.SetOutput(os.Stdout) // 或者不设置，默认打印到控制台

	// 日志打印到文件中
	file, err := os.OpenFile("./code/go_lib/log/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(file)
}

func simpleHttpGet01(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetch url... url: %s error: %v", url, err)
	} else {
		log.Printf("Http get succeeded url: %s statusCode: %d status: %s", url, resp.StatusCode, resp.Status)
	}
}
