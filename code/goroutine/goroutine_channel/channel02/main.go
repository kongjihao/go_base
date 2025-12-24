package main

import "fmt"

// 通道02

/*
需求：
 1. 生成0~100的数字发送到ch1中
 2. 从ch1中接收数据，并计算结果的平方,把结果发送到ch2中
*/

func product(ch1 chan int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}

	close(ch1) // ch1干完活了，把0~99的数据发送给ch1了，并存储在ch1中了，关闭ch1
}

func squireSum(ch1, ch2 chan int) {
	// 从通道中取值方式1
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}

		ch2 <- tmp * tmp
	}

	close(ch2) // ch2干完活了，把0~99的平方发送给ch2了，并存储在ch2中了，关闭ch2
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	// 开启两个goroutine
	go product(ch1)        // 生成0~99的数字发送到ch1中
	go squireSum(ch1, ch2) // 从ch1中接收数据，并计算结果的平方,把结果发送到ch2中

	fmt.Printf("ch2 type: %T\n", ch2)   // ch2 type: chan int
	fmt.Printf("ch2 value: %+v\n", ch2) // ch2 value: 0x1400012e380

	// 接收ch2中的数据， 从通道中取值方式2
	for ret := range ch2 {
		fmt.Println("ch2中接收到的数据: ", ret)
	}

}
