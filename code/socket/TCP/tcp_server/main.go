package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// tcp 网络通信编程学习
// tcp/server/main.go

/*
TCP/IP(Transmission Control Protocol/Internet Protocol) 即传输控制协议/网间协议，
是一种面向连接（连接导向）的、可靠的、基于字节流的传输层（Transport layer）通信协议，
因为是面向连接的协议，数据像水流一样传输，会存在黏包问题。
*/

/*
一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。
因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次链接就创建一个goroutine去处理。

TCP服务端程序的处理流程：
监听端口
接收客户端请求建立链接
创建goroutine处理链接。
*/

// 链接处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭链接

	reader := bufio.NewReader(conn) // 创建 bufio.reader 对象, tcp 数据流接收的标准定义
	for {
		var buf [512]byte             // buffer 读缓冲区
		n, err := reader.Read(buf[:]) // 基于字节流读取数据，返回（n int, err error）
		// 服务端在读取数据时，如果客户端关闭了连接，可能会收到 EOF
		if err != nil {
			if err == io.EOF {
				fmt.Println("client closed connection")
			} else {
				fmt.Println("tcp conn reader error! err: ", err)
			}
			return
		}

		recvStr := string(buf[:n]) // 将字节流强转为字符串类型
		fmt.Println("收到来自clinet发来的数据: ", recvStr)
		conn.Write([]byte(recvStr)) // 客户端发来什么数据，返回什么数据
	}
}
func main() {
	/*
		net.Listen 用于服务端监听指定的网络地址和端口，等待客户端发起连接。
		服务端通过 listener.Accept() 来接受客户端的连接请求，之后进行数据通信。
	*/
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("get tcp listen error! err: ", err)
		return
	}
	// for 死循环不断监听 client 的 tcp 请求
	for {
		conn, err := listen.Accept() // 接受客户端的连接请求。
		if err != nil {
			fmt.Println("make tcp conn error!, err: ", err)
			return
		}

		go process(conn) // 启动一个goroutine处理连接
	}
}
