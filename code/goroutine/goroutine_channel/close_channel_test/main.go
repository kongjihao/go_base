package main

import "fmt"

// 通道关闭后，仍然可以接收通道中已经发送但尚未接收的值。当所有的值都被取出后，for range 自动检测到通道已关闭，然后退出循环。
// 但是当一个通道（channel）被关闭后，任何对该通道的进一步 发送 操作都会导致运行时错误（panic）。
func main() {
	ch := make(chan int, 5)

	// 向通道发送一些数据
	ch <- 1
	ch <- 2
	ch <- 3

	// 关闭通道
	close(ch)

	// 使用 for range 读取通道中的值
	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("通道已关闭，所有值都被取出")
}
