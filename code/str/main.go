package main

import (
	"fmt"
	"strings"
)

// 字符串处理
func main() {
	// 字符串拼接1
	s := "hello world"
	s2 := "hello kongjihao"
	fmt.Println(s + s2)

	// 字符串拼接2
	s3 := fmt.Sprintf("%s - %s", s, s2)
	fmt.Println(s3)

	//字符串分割
	s4 := "how do you do?"
	// 使用 fmt.Println 正常输出字符串结果
	fmt.Println(strings.Split(s4, " ")) // [how do you do?]
	// 使用 fmt.Printf 格式化输出字符串结果
	fmt.Printf("%q\n", strings.Split(s4, " ")) // ["how" "do" "you" "do?"]
	fmt.Printf("%T\n", strings.Split(s4, " ")) // []string
	fmt.Printf("%s\n", strings.Split(s4, " ")) // [how do you do?]

}
