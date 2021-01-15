package main

import "fmt"

// 为任意类型添加方法
// 不能够为不是你这个包里定义的类型添加新的方法

// 基于内置的基本类型造一个我们自己的类型
type myInt int

func (m myInt) sayHi() {
	fmt.Println("hi")
}

func main() {
	var m1 myInt
	m1 = 100
	m1.sayHi()
}
