package main

// for 循环
func main() {
	// 1. 初始化语句, 普通for循环
	for i := 0; i < 10; i++ {
		println(i)
	}

	// 2.省略初始语句, 但必须保留初始化语句的分号
	j := 0
	for ; j < 10; j++ {
		println(j)
	}

	// 3.省略初始和结束语句
	k := 0
	for k < 10 {
		println(k)
		k++
	}

	// 4.无限循环, 死循环
	/*for {
		fmt.Println("hello")
	}*/

	// 5.break跳出循环
	for l := 0; l < 10; l++ {
		if l == 3 {
			break
		}
		println(l)
	}

	// 6.continue继续下一次循环
	for m := 0; m < 10; m++ {
		if m%2 == 0 {
			continue
		}
		println(m)
	}

	// 7.嵌套循环
	for n := 0; n < 3; n++ {
		for o := 'A'; o < 'A'+3; o++ {
			println(o)
		}
	}

	// 8.range关键字
	s := "abc"
	for p, v := range s {
		println(p, v)
	}
	s = "中华人民共和国"
	for q, v := range s {
		println(q, string(v))
	}

	// 9.goto语句, label标签
	i := 0
LABEL1:
	println(i)
	i++
	if i < 10 {
		goto LABEL1
	}
	return

}
