package main

import "fmt"

// 运算符
func main() {
	// 1. 算术运算符
	// a := 10
	// b := 20
	// c := a + b
	// d := a - b
	// e := a * b
	// f := a / b   // 0
	// g := a % b   // 10
	// fmt.Println(a, b, c, d, e, f, g)
	// a++
	 //fmt.Println(a)

	 // 2. 关系运算符
	//fmt.Println(10 == 3)
	//fmt.Println(10 == 10)

	// 3. 逻辑运算符
	//fmt.Println(10 > 2 && 1 == 1)
	//fmt.Println(10 < 2 || 3 < 2)

	// 4. 位运算符
	a := 1   // 001
	b := 5   // 101
	fmt.Println(a & b)   // 001
	fmt.Println(a | b)   // 101
	fmt.Println(a ^ b)   // 100   同位不一样为1
	fmt.Println(1 << 2)   // 100
	fmt.Println(4 >> 2)   // 1
	fmt.Println(1 << 10)   // 1024
}
