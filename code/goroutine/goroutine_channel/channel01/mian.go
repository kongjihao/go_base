package main

import "fmt"

// 通道01

func main() {
	// 声明一个通道，通道为引用类型，同slice、map一样，需要先使用make初始化再使用
	var ch1 chan int = make(chan int, 10) // 声明一个有缓冲区通道，容量为10
	// ch2 := make(chan int)                 // 无缓冲区通道
	/*
		go语言中有缓冲区和无缓冲区通道的区别:
			无缓冲区通道：发送数据时，必须有接收者正在接收数据.
			有缓冲区通道：发送数据时，如果没有接收者正在接收数据，则数据会先存入缓冲区中，如果缓冲区已满，则发送者阻塞，直到有接收者接收数据为止.
	*/

	ch1 <- 10
	len1 := len(ch1)   // 道的长度：1，因为还没有接收者，接收数据，所以通道的长度为1
	v := <-ch1         // 从通道中接收数据
	println("v = ", v) // 打印接收到的数据

	len2 := len(ch1) // 通道的长度：0，因为已经接收了数据
	cap1 := cap(ch1) // 通道的容量
	fmt.Printf("len1 = %d, len2 = %d, cap = %d\n", len1, len2, cap1)

	// 关闭通道
	close(ch1)
}
