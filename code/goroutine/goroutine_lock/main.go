package main

import (
	"fmt"
	"sync"
)

// 对全局变量的并发修改，通过加互斥锁限制解决
/*
	互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源。
	Go 语言中使用sync包中提供的Mutex类型来实现互斥锁。


	互斥锁是完全互斥的，但是实际上有很多场景是读多写少的，
	当我们并发的去读取一个资源而不涉及资源修改的时候是没有必要加互斥锁的，
	这种场景下使用读写锁是更好的一种选择。读写锁在 Go 语言中使用sync包中的RWMutex类型。rwLock sync.RWMutex


	2、sync.WaitGroup
	3、sync.Once
	4、Go 语言中内置的 map 不是并发安全的，所以并发的时候得使用 sync.Map{}
	5、原子操作atomic包:针对整数数据类型（int32、uint32、int64、uint64）我们还可以使用原子操作来保证并发安全，通常直接使用原子操作比使用锁操作效率更高。Go语言中原子操作由内置的标准库`sync/atomic`提供。
	详见：https://www.liwenzhou.com/posts/Go/concurrence/#c-0-6-4
*/

var (
	x    int
	wg   sync.WaitGroup
	lock sync.Mutex // 声明一个互斥锁对象
)

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func main() {
	// 开启两个 go 写成并发执行add函数
	wg.Add(2)
	go add()
	go add()
	wg.Wait()

	fmt.Println(x)
}
