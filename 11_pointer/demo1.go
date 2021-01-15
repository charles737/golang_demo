package main

import "fmt"

// 指针
func main() {
	a := 10   // int类型
	b := &a   // *int类型(int指针) 变量b是一个int类型的指针(*int)，它存储的是变量a的内存地址
	c := *b
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
