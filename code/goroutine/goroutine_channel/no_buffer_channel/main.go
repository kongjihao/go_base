package main

import (
	"fmt"
	"time"
)

// 无缓冲区channel教学
/*
1、无缓冲通道强制发送者和接收者在同一时刻进行同步。发送者会等待接收者准备好接收数据，接收者会等待发送者发送数据。
2、无缓冲通道通常用于在 goroutine 之间进行同步操作，因为它可以确保 goroutine 之间的通信是同步的。
3、如果没有接收者等待，发送操作会永远阻塞，反之亦然。这种特性在高并发程序中可以帮助控制 goroutine 的执行顺序。
*/

func main() {
	// 声明一个无缓冲通道
	ch := make(chan int)

	// 启动一个 goroutine，用于接收数据
	go func() {
		val := <-ch // 接收数据，发送者会阻塞直到接收者接收
		fmt.Println("接收到数据:", val)
	}()

	// 启动一个 goroutine，用于发送数据
	go func() {
		fmt.Println("发送数据: 42")
		ch <- 42 // 发送数据，直到接收者接收数据之前，这里会阻塞
		fmt.Println("数据已发送")
	}()

	// 等待一段时间，确保程序不会提前退出
	time.Sleep(2 * time.Second)
}
