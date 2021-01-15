package main

import (
	"fmt"
	"sort"
)

// 切片   切片是引用类型，必须初始化后才能使用
func main() {
	//var a []string
	//var b []int
	//var c = []bool{true, false}
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)

	//// 基于数组定义切片
	//a := [5]int{55, 56, 57, 58, 59}
	//b := a[1:3]
	//fmt.Println(b)
	//fmt.Printf("%T\n", b)
	//// 切片再次切片
	//c := b[1:len(b)]
	//fmt.Println(c)

	//// make函数构造切片
	//d := make([]int, 5, 10)
	//fmt.Println(d)
	//fmt.Printf("%T\n", d)
	//fmt.Printf("the length of d: %d\n", len(d))
	//fmt.Printf("the capacity of d: %d\n", cap(d))

	////nil
	//var a []int   // 声明int类型的切片
	//b := []int{}   // 声明并且初始化
	//c := make([]int, 0)
	//04_if a == nil {
	//	fmt.Println("a是一个nil")   // a是一个nil
	//}
	//fmt.Println(a, len(a), cap(a))
	//04_if b == nil {
	//	fmt.Println("b是一个nil")   // b不是一个nil,初始化后会创建一个数组
	//}
	//fmt.Println(b)
	//04_if c == nil {
	//	fmt.Println("c是一个nil")
	//}
	//fmt.Println(c)

	//// 切片是赋值拷贝
	//a := make([]int, 3)
	//b := a
	//b[0] = 100
	//// a, b指向同一个底层数组
	//fmt.Println(a)
	//fmt.Println(b)

	//// 切片的遍历
	//a := []int{1, 2, 3, 4, 5, 6}
	////05_for i := 0; i < len(a); i++ {
	////	fmt.Println(a[i])
	////}
	////fmt.Println()
	//05_for _, v := range a {
	//	fmt.Println(v)
	//}

	//// 切片的扩容
	//var a []int
	//a[0] = 1
	//fmt.Println(a)   // 错误写法：切片声明后必须初始化才能使用

	//var a []int
	//a := []int{}
	//a := make([]int, 0)
	//05_for i := 0; i < 10; i++ {
	//	a = append(a, i)
	//	fmt.Printf("%v len: %d cap: %d ptr: %p\n", a, len(a), cap(a), a)
	//}
	//a = append(a, 1, 2, 3, 5, 6, 7)   // append 一次往切片里添加多个元素
	//fmt.Println(a)
	//b := []int{32, 45, 76}
	//a = append(a, b...)
	//fmt.Println(a)

	// 切片的copy
	//a := []int{1, 2, 3, 4, 5}
	//b := make([]int , 5, 5)
	//copy(b, a)
	//b[0] = 100
	//fmt.Println("a:", a)
	//fmt.Println("b:", b)

	// 切片删除元素
	// 删除切片中index的元素: append(a[:index], a[index+1:]...)
	//a := []string{"上海", "北京", "深圳", "广州"}
	//fmt.Println(a)
	//a = append(a[0:2], a[3:]...)
	//fmt.Println(a)

	// 使用sort内置包对数组进行排序
	a := [...]int{2, 5, 6, 34, 567, 12, 546, 23, 43}
	sort.Ints(a[:])
	fmt.Println("a:", a)
}
