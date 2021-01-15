package main

import "fmt"

// 函数
// 定义一个不需要参数也没有返回值的函数:sayHello()
func sayHello() {
	fmt.Println("hahahello!")
}

// 定义一个接受string类型的name参数
func sayHello2(name string) {
	fmt.Println("Hello ", name)
}

// 定义接收多个参数的函数并且有一个返回值(没有声明返回值)
func intSum(a int, b int) int {
	ret := a + b
	return ret
}

// 定义接收多个参数的函数并且有一个返回值(声明了返回值)
func intSum2(a int, b int) (ret int) {
	ret = a + b
	return
}

// 函数接收可变参数
// 可变函数在函数体中是切片类型
func intSum3(a ...int) int {
	ret := 0
	for _, v := range a {
		ret = ret + v
	}
	return ret
}

// 当既出现固定参数，又出现可变参数，可变参数放在最后
func intSum4(a int, b ...int) int {
	ret := a
	for _, v := range b {
		ret = ret + v
	}
	return ret
}

// 定义具有多个返回值的函数
func calc(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	// 函数调用
	//sayHello()
	//name := "扎啤"
	//sayHello2(name)
	//sayHello2("小扎啤")
	//r := intSum2(10, 20)
	//fmt.Println(r)
	//r := intSum3(1, 3, 65, 6, 32, 6, 234, 2134)
	//fmt.Println(r)
	//r4 := intSum4(1, 3, 65, 6, 32, 6, 234, 2134)
	//fmt.Println(r4)
	// go语言的函数中没有默认参数
	sum, sub := calc(10, 32)
	fmt.Println(sum, sub)
}