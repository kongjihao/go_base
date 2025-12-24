package main

import (
	"fmt"
	"sync"
)

// 等待 go 协程执行完毕。再结束主程序。
var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("Hello, ", i)
	wg.Done() // 计数牌-1,告诉wg,一个go routine执行完毕。
}

// go routine 初识
func main() {
	wg.Add(1000) // 计数牌+1

	for i := 0; i < 1000; i++ {
		go hello(i) // 开启1000个goroutine,去执行hello函数。
	}
	// go hello() // 开启一个goroutine,去执行hello函数。
	fmt.Println("Hello, main!")
	// time.Sleep(time.Second) // 等待1s等待hello函数执行完毕。不推荐，推荐用waitgroup

	wg.Wait() // 等待计数牌归零。等待所有小弟都执行完毕。
}
