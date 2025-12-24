package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func A() {
	for i := 0; i < 10; i++ {
		fmt.Println("A ", i)
	}

	wg.Done()
}

func B() {
	for i := 0; i < 10; i++ {
		fmt.Println("B ", i)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(1) // only one cpu core
	wg.Add(2)

	go A()
	go B()
	wg.Wait() // wait for all goroutines to finish

	fmt.Println("main~~")
}
