package main

import (
	"fmt"
)

// 匿名函数和闭包
// 定义一个函数，它的返回值是一个函数
func a() func() {
	return func() {
		fmt.Println("hehe")
	}
}

func b() func() {
	name := "啦啦啦"
	return func() {
		fmt.Println("hehe", name)
	}
}

func main() {
	// 1. 将匿名函数赋值给一个变量
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()

	// 2. 在匿名函数花括号后面加括号，表示立即执行匿名函数
	func() {
		fmt.Println("匿名函数立即执行")
	}()

	// 闭包：闭包=函数+外层变量的引用
	r := b()   // r此时就是一个闭包
	r()   // 相当于执行了b函数内部的匿名函数
}