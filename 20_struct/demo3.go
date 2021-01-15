package main

import "fmt"

//结构体的初始化
// 1. 键值对的初始化
// 2. 值的列表进行初始化

type person2 struct {
	name, city string
	age int
}

func main() {
	// 1. 键值对初始化
	p5 := person2{
		name: "刘德华",
		age: 23,
		city: "纽约",
	}
	fmt.Printf("p5: %#v\n", p5)

	fmt.Println()

	// 2. 值的列表进行初始化
	p6 := person2{
		"王力宏",
		"曼哈顿",
		89,
	}
	fmt.Printf("p6: %#v\n", p6)
}
