package main

import (
	"go.uber.org/zap"
	"net/http"
)

// 学习原文：
// https://www.liwenzhou.com/posts/Go/zap/

// 为什么zap？
// zap是非常快的、结构化的，分日志级别的Go日志库。
// 为什么选择Uber-go zap
// 它同时提供了结构化日志记录和printf风格的日志记录
// 它非常的快
// 根据Uber-go Zap的文档，它的性能比类似的结构化日志包更好——也比标准库更快。 以下是Zap发布的基准测试信息

// Zap提供了两种类型的日志记录器—Sugared Logger和Logger。
// Logger:
// 通过调用zap.NewProduction()/zap.NewDevelopment()或者zap.Example()创建一个Logger。
// 上面的每一个函数都将创建一个logger。唯一的区别在于它将记录的信息不同。例如production logger默认记录调用函数信息、日期和时间等。
// 通过Logger调用Info/Error等。
// 默认情况下日志都会打印到应用程序的console界面。

var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync() // flushes buffer, if any

	simpleHttpGet("http://www.google.com")
	simpleHttpGet("http://www.baidu.com")
	simpleHttpGet("www.google.com") // 故意写错的URL，触发错误日志
}

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error( // Error 级别日志
			"Error fetch url...",
			zap.String("url", url),
			zap.Error(err),
		)
	} else { // 注意这里的逻辑，要有这个else
		logger.Info( // Info 级别日志
			"Http get succeeded",
			zap.String("url", url),
			zap.Int("statusCode", resp.StatusCode),
			zap.String("status", resp.Status),
		)
	}
}
