package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client demo

/*
一个TCP客户端进行TCP通信的流程如下：
建立与服务端的链接
进行数据收发
关闭链接
使用Go语言的net包实现的TCP客户端代码如下：

需要注意TCP 黏包的问题，与编程语言无关，粘包问题并不是语言本身引起的，而是由 TCP 协议的传输机制 决定的。
TCP 是基于字节流的协议，它并不提供消息边界的概念(需要应用层自己去解决消息边界的识别问题)，传输的数据可能会被拆分成多个包发送，
也可能会将多个包合并为一个包来传输。
- 解决 TCP 粘包问题的常见方案:

  - 1.固定长度消息
    发送和接收方约定每条消息的长度是固定的。接收方可以按固定长度读取数据。适合小消息或固定格式消息的传输。

  - 2. 使用特殊分隔符
    在每条消息后面附加一个特殊字符（如 \n 或自定义分隔符），接收方通过分隔符识别消息的边界。这种方式比较灵活，常用于文本数据的传输。

  - 3. 自定义消息头+消息体 (自定义协议)
    通过在每条消息前附加一个消息头，消息头中包含消息体的长度信息。接收方首先读取消息头的长度，然后根据该长度读取完整的消息体。这种方式适合于需要处理不同长度消息的场景，常用于高效的二进制协议中。
*/
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000") // 用于客户端发起网络连接，成功后可以通过返回的 conn 对象进行双向数据通信。
	if err != nil {
		fmt.Println("with Server make conn error! err: ", err)
		return
	}
	defer conn.Close() // 关闭链接

	inputReader := bufio.NewReader(os.Stdin) // 从终端输入，发送给tcp server
	for {
		input, err := inputReader.ReadString('\n') // 读取用户输入
		if err != nil {
			fmt.Println("user input error! err: ", err)
			return
		}
		inputInfo := strings.Trim(input, "\r\n") // 这个函数用于去除字符串两端指定的字符。在这里，去除了输入字符串两端可能存在的回车符'\r'和换行符'\n'，以确保输入内容的干净。
		if inputInfo == "" {
			continue
		}
		if strings.ToUpper(inputInfo) == "Q" { // 输入Q就退出
			return
		}

		// 发送数据
		_, err = conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("tcp conn write date error! err: ", err)
		}

		// 接收数据
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("Recv conn data error! err: ", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}
