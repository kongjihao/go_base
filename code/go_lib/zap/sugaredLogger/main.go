package main

import (
	"go.uber.org/zap"
	"net/http"
)

// 性能较低，但使用更简洁，那也比其他log库快上4倍
var sugaredLogger *zap.SugaredLogger

func main() {
	InitLogger01()
	defer sugaredLogger.Sync()
	simpleHttpGet01("http://www.google.com")
	simpleHttpGet01("http://www.baidu.com")
	simpleHttpGet01("www.google.com") // 故意写错的URL，触发错误日志
}

func InitLogger01() {
	logger, _ := zap.NewProduction()
	sugaredLogger = logger.Sugar()
}

func simpleHttpGet01(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugaredLogger.Errorf("Error fetching URL = %s, Error = %s", url, err)
	} else {
		sugaredLogger.Info(
			"Http get succeeded, ",
			"url: ", url,
			", statusCode: ", resp.StatusCode,
			", status: ", resp.Status)
	}
}
