package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var logger *zap.Logger

// 自定义zap logger，使用zap.New(core)
// zapcore.Core需要三个配置——Encoder(可修改全部日志打印格式)，WriteSyncer（写到哪里），LogLevel（写入的日志级别）
// 使用第三方库Lumberjack来实现按照日志大小切割的能力，并且支持日志文件的归档和删除

func main() {
	InitLogger()

	r := gin.New()
	r.Use(GinLogger(logger), GinRecovery(logger, true)) // GinRecovery中的true为记录程序panic时的堆栈信息
	r.GET("/ping", func(c *gin.Context) {
		logger.Debug(
			"receive request",
			zap.String("host", c.Request.Host),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("ip", c.Request.RemoteAddr),
			zap.String("user-agent", c.Request.Header.Get("User-Agent")),
		)

		c.String(http.StatusOK, "pong")
	})

	r.Run(":8080")
}

func InitLogger() {
	encoder := getEncoder()
	writer := getLogWriter()
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)

	// 当我们不是直接使用初始化好的logger实例记录日志，而是将其包装成一个函数等，此时日录日志的函数调用链会增加，
	// 想要获得准确的调用信息就需要通过AddCallerSkip函数来跳过。
	//logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger = zap.New(core, zap.AddCaller())

	/*
		有时候我们除了将全量日志输出到xx.log文件中之外，还希望将ERROR级别的日志单独输出到一个名为xx.err.log的日志文件中。
		我们可以通过以下方式实现:

		encoder := getEncoder()
		// test.log记录全量日志
		logF, _ := os.Create("./test.log")
		c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.DebugLevel)
		// test.err.log记录ERROR级别的日志
		errF, _ := os.Create("./test.err.log")
		c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zap.ErrorLevel)
		// 使用NewTee将c1和c2合并到core
		core := zapcore.NewTee(c1, c2)
		logger = zap.New(core, zap.AddCaller())
	*/

}

func getEncoder() zapcore.Encoder {
	// 对日志时间戳做修改，改成人类可读格式, 方式一：
	//encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder     // 修改时间格式
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder  // 大写格式
	//return zapcore.NewJSONEncoder(encoderConfig)

	// 对日志时间戳做修改，改成人类可读格式, 方式一：直接将zap.NewProductionEncoderConfig的结构体copy过来修改
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zapcore.NewConsoleEncoder(encoderConfig)

	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()) // 使用JSON格式
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()) // 使用控制台格式
}

func getLogWriter() zapcore.WriteSyncer {
	// 打印日志到文件和终端
	//file, _ := os.OpenFile("./code/go_lib/zap/gin_logger/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0744)
	//mw := io.MultiWriter(os.Stdout, file) // 将日志只打印到日志文件时，将MultiWriter删除，直接将file封装成Sync并返回
	//return zapcore.AddSync(mw)

	// 根据日志文件大小进行切割， 线上可以预估下1h有多大量的日志从而实现按照小时切割的效果
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./code/go_lib/zap/gin_logger/app.log",
		MaxSize:    1,     // 单位MB
		MaxBackups: 3,     // 最多保留3个备份
		MaxAge:     28,    // 最多保留28天
		Compress:   false, // 是否压缩
	}

	return zapcore.AddSync(io.MultiWriter(os.Stdout, lumberJackLogger))

	/*
		官方的说法是为了添加日志切割归档功能，我们将使用第三方库Lumberjack来实现。
		目前只支持按文件大小切割，原因是按时间切割效率低且不能保证日志数据不被破坏。详情戳https://github.com/natefinch/lumberjack/issues/54。
		想按日期切割可以使用github.com/lestrrat-go/file-rotatelogs这个库，虽然目前不维护了，但也够用了:

		// 使用file-rotatelogs按天切割日志
		import rotatelogs "github.com/lestrrat-go/file-rotatelogs"

		l, _ := rotatelogs.New(
			filename+".%Y%m%d%H%M",
			rotatelogs.WithMaxAge(30*24*time.Hour),    // 最长保存30天
			rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
		)
		zapcore.AddSync(l)
	*/
}

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
