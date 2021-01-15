package main

import "fmt"

// 一个类型可以同时实现多个接口
// 分别使用值接收者和指针接收者实现接口的区别

// 接口的嵌套
type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age int8
}

// 使用值接收者实现接口：类型的值和类型的指针都能够保存到接口变量中
//func (p person) move() {
//	fmt.Printf("%s在跑...\n", p.name)
//}

// 使用指针接收者实现接口：只有类型的指针能够保存到接口变量中
func (p *person) move() {
	fmt.Printf("%s在跑...\n", p.name)
}

func (p *person) say()  {
	fmt.Printf("%s在说...\n", p.name)
}

func main() {
	var m mover // 定义一个mover类型的变量
	var s sayer // 定义一个sayer类型的变量
	var a animal // 定义一个sayer类型的变量

	//p1 := person{ // p1 是person类型的值
	//	name: "刘德华",
	//	age: 18,
	//}
	//m = p1
	//m.move()
	//fmt.Println(m)

	fmt.Println()

	p2 := &person{ // p2 是person类型的指针
		name: "周星驰",
		age: 28,
	}
	m = p2
	s = p2
	m.move()
	s.say()
	fmt.Println(m)
	fmt.Println(s)

	fmt.Println()

	a = p2
	a.move()
	a.say()
	fmt.Println(a)
}