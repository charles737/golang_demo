package main

import "fmt"

// 函数进阶之变量作用域
// 定义全局变量num
var num = 10

// 定义函数
func testGlobal() {
	num := 100
	// 可以在函数中使用全局变量
	// 1. 先在自己的函数内部查找，找到了就用自己函数内部的变量
	// 2. 函数中找不到了在外层中寻找全局变量
	fmt.Println("全局变量", num)
}

// 函数作为参数传入
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calcu(x, y int, op func(int, int)int) int {
	return op(x, y)
}

func main() {
	// 函数可以作为变量
	a := testGlobal
	fmt.Printf("%T\n", a)
	a()
	ret1 := calcu(100, 200, add)
	fmt.Printf("ret1: %d\n", ret1)
	ret2 := calcu(200, 300, sub)
	fmt.Printf("ret2: %d\n", ret2)
}

