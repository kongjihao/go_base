package main

import "fmt"

// 结构体继承实现

type Animal struct {
	Name string
}

func (a *Animal) Move() {
	fmt.Printf("%s Moving\n", a.Name)
}

type Dog struct {
	Feet    int8 // 腿
	*Animal      // 即可以嵌套结构体类型，又可以嵌套结构体指针类型，此处是匿名嵌套的是结构体指针类型
}

func (d *Dog) Wang() {
	fmt.Printf("%s WangWang\n", d.Name)
}

func main() {
	dog := &Dog{
		Feet: 4,
		Animal: &Animal{
			Name: "Tom",
		}, // 此处的Name是Dog结构体中嵌套的Animal结构体的Name
	}

	// 此处调用的Move方法是嵌套的Animal结构体指针类型的方法，此时体现了了继承
	dog.Move()
	dog.Wang()

}
