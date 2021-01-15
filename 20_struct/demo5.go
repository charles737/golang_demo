package main

import "fmt"

// 结构体的匿名字段

type Person5 struct {
	string
	int
}

func main() {
	p1 := Person5{
		"刘德华",
		23,
	}

	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}