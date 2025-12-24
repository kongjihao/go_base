package main

import "fmt"

// 切片的赋值拷贝
func main() {
	// 切片赋值拷贝
	a := []int{1, 2, 3}
	b := a
	fmt.Println("a1 value: ", a, "a length:", len(a), "a capacity:", cap(a))
	fmt.Println("b1 value: ", b, "b length:", len(b), "b capacity:", cap(b))
	b[0] = 4 // 说明底层指针指向了同一地址，是c的拷贝，所以修改d的值会影响c，均指向c的内存地址
	fmt.Println("a2 value: ", a, "a length:", len(a), "a capacity:", cap(a))
	fmt.Println("b2 value: ", b, "b length:", len(b), "b capacity:", cap(b))

	fmt.Println()

	// 切片对比数组
	c := [3]int{1, 2, 3}
	d := c[:]
	fmt.Println("c1 value: ", c, "c length:", len(c), "c capacity:", cap(c))
	fmt.Println("d1 value: ", d, "d length:", len(d), "d capacity:", cap(d))
	d[0] = 4 // 说明底层指针指向了同一地址，是c的拷贝，所以修改d的值会影响c，均指向c的内存地址
	fmt.Println("c2 value: ", c, "c length:", len(c), "c capacity:", cap(c))
	fmt.Println("d2 value: ", d, "d length:", len(d), "d capacity:", cap(d))

}
