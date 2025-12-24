package main

import "fmt"

// 切片学习，切片本质其实就是对底层数组的一个封装，它是一个引用类型 默认引用类型的0值(基础类型的0值为0)为nil
func main() {
	/*
		切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
		切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
	*/

	// 1.初始切片，切片是引用类型
	var a []string
	var b []int32
	var c = []bool{true, false}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 2.基于数组定义切片
	var d = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s := d[2:5] // 左闭右开区间
	fmt.Println("s:", s)

	// 3.通过make函数创建切片
	e := make([]int, 3, 5) // 长度为3，容量为5，当省略容量cap时，默认等于第二个参数长度len
	fmt.Println("e:", e)

	// 4.通过len获取切片长度和cap获取切片容量
	fmt.Printf("len:%d cap:%d slice:%v\n", len(e), cap(e), e)

	// 5.切片可以再切片
	g := e[:2]
	k := g[:len(e[:2])]
	fmt.Println("g:", g)
	fmt.Println("k:", k)

}
