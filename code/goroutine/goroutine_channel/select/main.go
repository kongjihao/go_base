package main

import "fmt"

// select

func main() {
	ch := make(chan int, 1) // 要是把缓冲区改为10，则for 循环第二轮之后，就是每次所有case就满足了，select就会随机执行case
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // 最终结果：0 2 4 6 8
		case ch <- i: // 首先先满足这个case
		default:
			fmt.Println("啥都不干")
		}
	}
}
