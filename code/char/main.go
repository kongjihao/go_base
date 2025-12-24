package main

import "fmt"

func main() {
	// 1. 字符类型
	var a byte = 'a'                 // uint8
	var b rune = 'a'                 // int32
	fmt.Println(a, b)                // 97 97
	fmt.Printf("a:%T\tb:%T\n", a, b) // a:uint8 b:int32

	// 2. 字符类型转换，观察这两种遍历的区别
	// 字符串索引遍历
	s1 := "hello七米"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c ", s1[i]) // h e l l o ä ¸  ç ± ³
	}

	// for range遍历字符串,处理中文字符的时候最好用for range遍历
	for _, r := range s1 {
		fmt.Printf("%c ", r) // h e l l o 七 米
	}
}
