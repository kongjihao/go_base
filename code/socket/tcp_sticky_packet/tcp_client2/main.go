package main

import (
	"fmt"
	"learnProject02/src/code/socket/tcp_sticky_packet/proto"
	"net"
)

// 接下来在服务端和客户端分别使用上面定义的proto包的Decode和Encode函数处理数据。

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
