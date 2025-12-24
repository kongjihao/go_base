package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 声明一个有缓冲区通道和一个无缓冲区通道
	bufferedChan := make(chan int, 3)
	unbufferedChan := make(chan int)

	var wg sync.WaitGroup

	// 启动多个接收者 goroutine 来接收 bufferedChan 的值
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for val := range bufferedChan {
				fmt.Printf("Buffered Receiver %d received: %d\n", id, val)
				time.Sleep(500 * time.Millisecond) // 模拟处理延迟
			}
			fmt.Printf("Buffered Receiver %d finished.\n", id)
		}(i)
	}

	// 启动多个接收者 goroutine 来接收 unbufferedChan 的值
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for val := range unbufferedChan {
				fmt.Printf("Unbuffered Receiver %d received: %d\n", id, val)
				time.Sleep(500 * time.Millisecond) // 模拟处理延迟
			}
			fmt.Printf("Unbuffered Receiver %d finished.\n", id)
		}(i)
	}

	// 发送者 goroutine 向 bufferedChan 发送数据
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending to bufferedChan: %d\n", i)
			bufferedChan <- i
			time.Sleep(200 * time.Millisecond) // 模拟发送延迟
		}
		close(bufferedChan) // 发送完成，关闭通道
	}()

	// 发送者 goroutine 向 unbufferedChan 发送数据
	go func() {
		for i := 10; i < 16; i++ {
			fmt.Printf("Sending to unbufferedChan: %d\n", i)
			unbufferedChan <- i
			time.Sleep(200 * time.Millisecond) // 模拟发送延迟
		}
		close(unbufferedChan) // 发送完成，关闭通道
	}()

	// 等待所有接收者完成
	wg.Wait()

	fmt.Println("All receivers finished.")
}
