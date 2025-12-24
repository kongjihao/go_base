package main

// 单向通道，限制通道中只能发送数据，或者只能从通道中接收数据

// product函数中传参就是一个只发送数据的单向通道
func product(ch1 chan<- int) {
	for i := 0; i < 100; i++ {
		ch1 <- i // ch1只能发送，不能被接收
	}

	close(ch1) // 将数据向ch1通道中发送完成后，关闭通道
}

// Consumer函数中传入的两个参数，分别是一个只接收数据的单向通道ch2和一个可以发送数据的通道ch1
func Consumer(ch1 <-chan int, ch2 chan<- int) {
	// 循环从ch1通道中接收数据，并将接收到的数据做平方后发送到ch2通道中
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}

		ch2 <- tmp * tmp
	}
	close(ch2) // 关闭ch2通道
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go product(ch1)
	go Consumer(ch1, ch2)

	for ret := range ch2 {
		println("ch2 = ", ret)
	}
}
