package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1.练习题一，更好的了解切片append
	a := make([]int, 5, 10)
	var c []int

	for i := 0; i < 10; i++ {
		a = append(a, i) // 等价于 a = append(a[:], i) 等价于 a = append(a, fmt.Sprintf("%d", i))
		c = append(a, i) // 将a的内容复制给切片c，这样c和a会共享相同的底层数组，导致对c的修改也会影响到a
	}
	// a slice: [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9], a len: 15, a cap: 20
	fmt.Printf("a slice: %v, a len: %d, a cap: %d\n", a, len(a), cap(a))
	fmt.Printf("c slice: %v, c len: %d, c cap: %d\n", c, len(c), cap(c))

	// 2.练习题二，利用切片能力对指定数组排序
	b := [...]int{3, 7, 8, 9, 1}
	sort.Ints(b[:]) // 此处[:]表示切片，即b[0:len(b)]，此时b[:]切片指针指向的是b[0]，也就是b数组的首地址
	fmt.Println(b)  //[1 3 7 8 9]
}
