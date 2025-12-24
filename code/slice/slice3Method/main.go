package main

import "fmt"

func main() {
	// append方法，切片的扩容机制,查看源码在：GOROOT/src/runtime/slice.go
	var a []int
	// a = append(a, 1) // 一定要有接收值s，否则会报错，因为append方法返回的是新的切片
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("slice: %v len: %d cap: %d ptr: %p\n", a, len(a), cap(a), a)
	}
	/*
		slice: [0] len: 1 cap: 1 ptr: 0x140000a6018
		slice: [0 1] len: 2 cap: 2 ptr: 0x140000a6040
		slice: [0 1 2] len: 3 cap: 4 ptr: 0x140000be020
		slice: [0 1 2 3] len: 4 cap: 4 ptr: 0x140000be020
		slice: [0 1 2 3 4] len: 5 cap: 8 ptr: 0x140000a80c0
		slice: [0 1 2 3 4 5] len: 6 cap: 8 ptr: 0x140000a80c0
		slice: [0 1 2 3 4 5 6] len: 7 cap: 8 ptr: 0x140000a80c0
		slice: [0 1 2 3 4 5 6 7] len: 8 cap: 8 ptr: 0x140000a80c0
		slice: [0 1 2 3 4 5 6 7 8] len: 9 cap: 16 ptr: 0x140000c0000
		slice: [0 1 2 3 4 5 6 7 8 9] len: 10 cap: 16 ptr: 0x140000c0000
	*/

	// 2. append一次追加多个元素
	var b []int
	c := []int{7, 8, 9}
	b = append(b, 4, 5, 6)
	fmt.Println("b =", b)
	d := append(b, c...)
	fmt.Println("d =", d)

	// 3. 切片的拷贝copy, 和引用的区别在于，拷贝后，两个切片的地址是不一样的
	e := make([]int, 3)
	copy(e, b)
	f := e
	fmt.Println("b = ", b)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	e[0] = 100
	fmt.Println("b2 = ", b)
	fmt.Println("e2 = ", e)
	fmt.Println("f2 = ", f)
	/*
		b =  [4 5 6]
		e =  [4 5 6]
		f =  [4 5 6]

		b2 =  [4 5 6]
		e2 =  [100 5 6]
		f2 =  [100 5 6]
	*/

	// 4. 使用append实现切片删除元素, 删除元素后，后面的元素会往前移动
	g := []string{"a", "b", "c", "d"}
	h := g[:2]
	i := g[2:]
	j := append(h, i...)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	/*
		g =  [a b c d]
		h =  [a b]
		i =  [c d]
		j =  [a b c d]
	*/

	// 5.使用copy实现切片删除元素
	k := []string{"a", "b", "c", "d"}
	l := k[:2]
	m := k[2:]
	fmt.Println("k1 = ", k)
	fmt.Println("l1 = ", l)
	fmt.Println("m1 = ", m)
	copy(m, l)
	fmt.Println("k2 = ", k)
	fmt.Println("l2 = ", l)
	fmt.Println("m2 = ", m)
	/*
		k1 =  [a b c d]
		l1 =  [a b]
		m1 =  [c d]
		k2 =  [a b a b]
		l2 =  [a b]
		m2 =  [a b]
	*/

}
