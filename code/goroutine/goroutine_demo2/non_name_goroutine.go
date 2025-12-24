package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		// 注意这里的匿名函数，要用匿名函数内部的i 值，不要用for 循环外面的i 值，
		// 否则都打印1000，因为形成了闭包，导致i值都是1000
		// 此处解决了闭包问题
		go func(i int) {
			fmt.Println("Hello ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
