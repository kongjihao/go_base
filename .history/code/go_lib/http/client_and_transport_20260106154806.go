package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// -------------------------- 1. 自定义 Transport --------------------------
// CustomTransport 自定义 HTTP Transport，实现 http.RoundTripper 接口
// 嵌套 http.Transport 以复用默认连接池、TLS 等底层能力
type CustomTransport struct {
	BaseTransport http.RoundTripper // 底层 Transport（默认用 http.DefaultTransport）
	AuthToken     string             // 自定义鉴权 Token（示例：用于全局添加鉴权 Header）
}

// RoundTrip 实现 http.RoundTripper 接口的核心方法，处理单次 HTTP 传输
func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// 1. 前置处理：添加自定义逻辑（如鉴权 Header、请求日志）
	// 自动添加鉴权 Header（适用于需要全局鉴权的场景）
	if t.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+t.AuthToken)
	}
	// 打印请求详情（开发环境调试用，生产环境可关闭或输出到日志系统）
	reqDump, err := httputil.DumpRequestOut(req, false) // false 表示不打印请求体
	if err != nil {
		log.Printf("Failed to dump request: %v", err)
	} else {
		log.Printf("\n=== Request Start ===")
		log.Printf("%s", string(reqDump))
		log.Printf("=== Request End ===")
	}

	// 2. 调用底层 Transport 执行实际请求（若未设置 BaseTransport，默认用 http.DefaultTransport）
	if t.BaseTransport == nil {
		t.BaseTransport = http.DefaultTransport
	}
	resp, err := t.BaseTransport.RoundTrip(req)
	if err != nil {
		return nil, fmt.Errorf("transport request failed: %w", err) // 包装错误，便于排查
	}

	// 3. 后置处理：添加响应逻辑（如打印响应状态、错误重试）
	log.Printf("\n=== Response Start ===")
	log.Printf("Status Code: %d", resp.StatusCode)
	log.Printf("Response Header: %v", resp.Header.Get("Content-Type"))
	log.Printf("=== Response End ===")

	return resp, nil
}

// -------------------------- 2. 自定义 Client --------------------------
// NewCustomClient 创建自定义 HTTP Client，支持传入鉴权 Token 和默认超时
func NewCustomClient(authToken string, defaultTimeout time.Duration) *http.Client {
	// 初始化自定义 Transport，传入鉴权 Token
	customTransport := &CustomTransport{
		AuthToken: authToken,
		// 可选：自定义底层 Transport 配置（如调整连接池大小、禁用 Keep-Alive）
		BaseTransport: &http.Transport{
			MaxIdleConns:        10,          // 最大空闲连接数
			IdleConnTimeout:     30 * time.Second, // 空闲连接超时时间
			DisableCompression:  false,       // 启用压缩
			DisableKeepAlives:   false,       // 启用 Keep-Alive
		},
	}

	// 返回自定义 Client，绑定 Transport 和默认超时
	return &http.Client{
		Transport: customTransport,
		Timeout:   defaultTimeout, // 客户端默认超时（Context 超时优先级更高）
	}
}

// -------------------------- 3. 实际使用示例 --------------------------
func main() {
	// 1. 初始化自定义 Client（传入鉴权 Token、默认超时 10 秒）
	authToken := "your-actual-auth-token-123" // 替换为真实鉴权 Token
	client := NewCustomClient(authToken, 10*time.Second)

	// 2. 创建带超时的 Context（5 秒超时，优先级高于 Client 默认超时）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保函数退出时释放 Context 资源，避免内存泄漏

	// 3. 监听系统信号（如 Ctrl+C），支持主动取消请求
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // 监听中断、终止信号
	go func() {
		<-sigChan
		log.Println("\nReceived exit signal, canceling request...")
		cancel() // 触发 Context 取消，终止正在进行的请求
	}()

	// 4. 构建 HTTP 请求（绑定 Context，确保超时/取消生效）
	reqURL := "https://httpbin.org/delay/3" // 测试接口：模拟 3 秒延迟响应
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,  // 请求方法
		reqURL,          // 请求 URL
		nil,             // 请求体（GET 方法为空）
	)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	// 可选：添加自定义请求头（如 User-Agent、链路追踪 ID）
	req.Header.Set("User-Agent", "Go-Custom-Client/1.0")
	req.Header.Set("X-Trace-ID", fmt.Sprintf("trace-%d", time.Now().UnixNano())) // 唯一追踪 ID

	// 5. 发送请求并处理响应
	log.Printf("Sending request to: %s (timeout: 5s)", reqURL)
	resp, err := client.Do(req)
	if err != nil {
		// 判断错误是否由 Context 超时/取消引起
		select {
		case <-ctx.Done():
			log.Fatalf("Request failed: %v (reason: %v)", err, ctx.Err())
		default:
			log.Fatalf("Request failed: %v", err)
		}
	}
	defer resp.Body.Close() // 确保响应体关闭，避免资源泄漏

	// 6. 读取响应体（示例：读取前 200 字节，实际场景需用 io.ReadAll 完整读取）
	log.Println("Request succeeded! Response body (first 200 bytes):")
	buf := make([]byte, 200)
	n, err := resp.Body.Read(buf)
	if err != nil && err.Error() != "EOF" { // 忽略正常的 EOF 错误（读取到末尾）
		log.Fatalf("Failed to read response body: %v", err)
	}
	log.Printf("%s", string(buf[:n]))
}