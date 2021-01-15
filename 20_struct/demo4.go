package main

import "fmt"

// 结构体构造函数: 构造一个结构体实例的函数

type person3 struct {
	name, city string
	age int
}

// 构造函数
// 当构造函数返回的结构体体积较大，性能开销大时，通常返回结构体的指针类型，保证程序运行的性能
func newPerson3(name, city string, age int) *person3 {
	return &person3{
		name: name,
		city: city,
		age: age,
	}
}

func main() {
	p1 := newPerson3("刘德华", "芝加哥", 234)
	fmt.Printf("type: %T\n value: %#v\n", p1, p1)
}
