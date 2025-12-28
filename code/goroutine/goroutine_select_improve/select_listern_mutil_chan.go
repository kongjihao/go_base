package main

import (
	"fmt"
	"time"
)

func foo(i int) chan int {
	c := make(chan int)
	go func() {
		c <- i
	}()
	return c
}

// foo 函数启动的 goroutine 向 ch1, ch2, ch3 中发送数据。
// 主函数中的 select 会监听 ch1, ch2, ch3，并将数据转发到 ch。
// 主函数从 ch 中读取数据并打印。
// 如果通道 ch 没有被关闭，for v := range ch 会导致死锁。
func main() {
	ch1, ch2, ch3 := foo(3), foo(6), foo(9)
	ch := make(chan int)

	// Start a goroutine to listen on multiple channels
	go func() {
		// 等待一会儿，确保 foo 函数的 goroutine 有时间发送数据，因为 foo 函数和 main 函数中的goroutine是并发执行的
		time.Sleep(1000 * time.Millisecond)
		for {
			// Use select to listen on multiple channels
			select {
			case v1 := <-ch1:
				ch <- v1
			case v2 := <-ch2:
				ch <- v2
			case v3 := <-ch3:
				ch <- v3
			default:
				// 如果所有通道都没有数据可读，default分支会被执行，防止阻塞。
				// 这里可以选择继续等待或者退出循环。
				close(ch)
				return
			}
		}
	}()

	// 报错死锁
	// for range 会一直尝试从通道 ch 中读取数据，直到通道被关闭。
	// 但是上面goroutine中的select语句会不断地等待从ch1、ch2、ch3中读取数据并发送到ch通道，
	// 而这些通道只会各发送一次数据后就不再发送更多数据了。
	// 且最后也没有关闭ch通道。所以在ch通道中的数据被读取完后，
	// for range循环会继续尝试从ch中读取数据，但此时没有更多的数据可读，导致程序进入死锁状态。
	for v := range ch {
		fmt.Println(v)
	}

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(<-ch)
	// }
}

// 为什么我第一次执行能读取到管道中的值，第一次执行后管道被关闭，但是我重新执行了整体函数为什么后面读取不到了呢

// 第一次运行时，程序能够正常输出，因为 goroutine 的调度顺序刚好满足了数据的发送和接收。
// 后续运行时，可能因为 default 分支过早关闭通道 ch，导致主函数无法读取到数据。
// 解决方法是通过计数器或去掉 default 分支，确保通道关闭的时机正确。
