package main

// select 关键字学习，可以用于在多个通道操作中进行选择，从而实现并发控制。
// main函数所在的goroutine是程序的主goroutine，负责执行程序的入口逻辑。
func main() {
	/*
		select语句的特点：
		- select语句会阻塞，直到其中一个case满足条件。
		- 如果有多个case同时满足条件，Go语言会随机选择一个执行。
		- 如果没有case满足条件，并且有default子句，则会执行default子句。
		- 如果没有default子句，且没有任何case满足条件，select语句会一直阻塞，直到有case满足条件为止。
	*/

	// 1. 随机选择一个case执行
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() { ch1 <- 1 }()
	go func() { ch2 <- "我是字符串" }()
	// 此时会随机选择一个case执行
	select {
	case v := <-ch1:
		println("received from ch1", v)
	case v := <-ch2:
		println("received from ch2", v)
	}

	// 2.select第二个举例
	// ch1 := make(chan int)
	// ch2 := make(chan string)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch1 <- 1
	// }()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	ch2 <- "Hello"
	// }()

	// select {
	// case val := <-ch1:
	// 	fmt.Println("Received from ch1:", val)
	// case msg := <-ch2:
	// 	fmt.Println("Received from ch2:", msg)
	// }
	/*
		在上面代码中确实在main函数中启动了两个goroutine，分别执行了两个匿名函数，一个在1秒后向ch2通道发送数据，
		另一个在2秒后向ch1通道发送数据。select语句用于等待多个通道中的任意一个发送消息。
		在这个例子中，它等待ch1或ch2中的消息。哪个通道先发送消息，select就会执行相应的case分支。
		由于ch2的休眠时间更短（1秒），它通常会先发送消息，因此程序通常会输出“Received from ch2: Hello”。
	*/
}
