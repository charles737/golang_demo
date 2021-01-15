package main

import "fmt"

// 定义结构体

type person struct {
	name string
	age int
	city string
}

func main() {
	var p person
	p.name = "朱古力"
	p.city = "天京"
	p.age = 290
	fmt.Printf("p: %#v\n", p)
	fmt.Println(p.age)
	fmt.Println(p.city)
	fmt.Println(p.name)

	// 匿名结构体
	var user struct {
		name string
		married bool
	}
	user.name = "asdf"
	user.married = false
	fmt.Printf("user: %#v\n", user)
}