package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // When worker finishes, decrement the wait group counter
	for job := range jobs {
		fmt.Println("worker", id, "started  job", job)
		results <- job * 2
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		jobs <- i
	}

	close(jobs)

	// Start workers and increment wait group counter for each
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Start a goroutine to close results channel after all workers are done
	go func() {
		wg.Wait()      // Wait for all workers to finish
		close(results) // Close results after all workers finish
	}()

	// Consume the results
	for ret := range results {
		fmt.Println("result: ", ret)
	}
}

/*
	自己写的逻辑，代码健壮性不高，需要优化，尤其是自己写程序的时候没想明白results channel 会被关闭3此，导致运行时报错
package main

import (
	"fmt"
	"time"
)

// workPool

func worker(id int, jobs <-chan int, results chan<- int) {
	// 多goroutine运行次函数，此时就提现了channel的多路通信的特点了，channel的多路通信特性
	for job := range jobs {
		fmt.Println("worker", id, "started  job", job)
		results <- job * 2 // send result to results
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", job)
	}

	close(results) // close channel，这里三个go routine都close了，所以这里运行时会报错，因为channel被关闭了3次
}

func main() {
	// 初始化jobs和results
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// for循环构造jobs
	for i := 0; i < 5; i++ {
		jobs <- i
	}

	close(jobs) // close channel，别忘记!否则会导致死锁

	// for循环构造worker池子，开启3个goroutine，分别处理jobs
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	// 等待所有worker结束，for循环遍历results，打印结果
	// for j := 0; j < 5; j++ {
	// 	fmt.Println("result: ", <-results)
	// }

	for ret := range results {
		fmt.Println("result: ", ret)
	}
}



// 下面这段代码，更是有问题，因为close(results) 的时机没有想明白，所以导致输出结果更不符合预期
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for {
		tmp, ok := <- jobs
		if !ok {
			break
		}
		fmt.Println("worker", id, "started  job", tmp)
		results <- tmp * 2 // send result to results
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", tmp)
	}

	close(results)
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for i := 0; i < 5; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for ret := range results {
		fmt.Println("result: ", ret)
	}
}
*/
