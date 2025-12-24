package main

import "fmt"

// 使用值接收者实现接口 和 使用指针接收者实现接口的区别

// 接口嵌套
type animal interface {
	mover // mover接口
	sayer // sayer接口
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int
}

// 使用值接受者实现接口: 类型的值 和 类型的指针都可以保存到接口变量m中
func (p person) move() {
	println("person move")
}

func (p person) say() {
	println("person say")
}

// 当使用指针类型的接受者实现接口时，只有指针类型的变量才可以保存到接口变量m中
// func (p *person) move() {
// 	println("person move")
// }

type cat struct{}

func (c cat) move() {
	println("cat move")
}

func (c cat) say() {
	println("cat say")
}

func main() {
	var m mover // 定义一个mover接口类型的变量 m
	var s sayer
	p := person{name: "tom", age: 18} // 定义一个person类型的变量 p
	m = p
	m.move()
	s = p
	s.say()
	fmt.Printf("m1: %v\n", m)

	p2 := &person{name: "jerry", age: 20} // 定义一个person类型的指针变量 p2
	m = p2
	m.move()
	fmt.Printf("m1: %v\n", m)

	c := cat{}
	m = c
	m.move()
	fmt.Printf("m2: %v\n", m)

	var a animal // 定义一个animal接口类型的变量 a，接口嵌套类型
	c2 := cat{}
	a = c2
	a.move()
	a.say()
}
