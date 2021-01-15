package main

import "fmt"

// new and make
func main() {
	// new
	// 错误示例
	//var a *int
	//*a = 100
	//fmt.Println("*a: ", *a)   // a = nil 报错
	//
	//var b 09_map[string]int
	//b["朱古力"] = 100
	//fmt.Println("b: ", b)   // 报错 map是引用类型，必须通过make来初始化，在内存中占有一块内存空间

	// 正确示例
	var a *int
	a = new(int)
	*a = 100
	fmt.Println("*a: ", *a)

	// make
	// make用于内存分配，只用于slice, 09_map, chan的内存创建，返回的类型就是这三个类型本身，而不是它们的指针类型
	var b map[string]int
	b = make(map[string]int, 10)
	b["朱古力"] = 100
	fmt.Println("b: ", b)

	// new 和 make 的区别：
	// 1. 二者都是用来做内存分配的
	// 2. make只用于slice,09_map,chan的初始化，返回的还是这三个引用类型本身
	// 3. new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
}
