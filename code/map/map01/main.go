package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// map学习
func main() {
	/*
		map 是一个无序的key-value集合，map为引用类型，必须初始化才能使用。map的长度是不固定的，
		map的长度为零时，称为nil map。其格式为: map[KeyType]ValueType
	*/
	// 1. 声明一个map
	var m map[string]int
	// 2. 初始化一个map
	m = make(map[string]int, 10)
	// 上面1、2两行代码等价于：m := make(map[string]int, 10)
	// 3. Map 如何添加kv
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	// 遍历map，key-value的顺序与添加的顺序无关，因为map是无序的
	for k, v := range m {
		println(k, v)
	}
	// 只遍历map的key
	for k := range m {
		println(k)
	}

	// 只遍历map的value
	for _, v := range m {
		println(v)
	}

	// 4. Map 如何删除key
	delete(m, "a")
	if _, ok := m["a"]; !ok {
		println("key \"a\" not exist!")
	}

	// 5.声明map的时候，初始化map
	a := map[int]bool{
		1: true,
		2: false,
	}
	for k, v := range a {
		println(k, v)
	}

	// 6.演示没有初始化map，直接使用，报错的情况
	// var b map[int]bool
	// b[1] = true // 报错
	// for k, v := range b {
	// 	println(k, v)
	// }

	// 7. 判断某个键是否存在
	value, ok := m["d"]
	fmt.Println("Is \"d\" key exist?", ok)
	if ok {
		fmt.Println("The value of d is ", value)
	} else {
		fmt.Println("The key \"d\" not exist! not key & value!")
	}

	// 8. 按照某个顺序区遍历map
	var scoreMap = make(map[string]int, 10)

	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("student%02d", i)
		value := rand.Intn(10) // 0~9

		scoreMap[key] = value
	}
	fmt.Println("soreMap1: ", scoreMap)

	// 按照key的序号从小到大遍历
	// 1. 先获取所有的key
	// 2. 对key进行排序
	// 3. 遍历key
	keys := make([]string, 0, len(scoreMap)) // 注意：初始化一个长度为0的切片，不要设置成10，否则会把soreMap的key都遍历一遍,就都为0了
	for key := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(scoreMap[v])
	}
}
