package main

import "fmt"

// panic and recover

func a() {
	fmt.Println("func a")
}

func b() {
	defer func() {
		// 1. recover必须搭配defer使用
		// 2. defer一定要在可能引发panic的语句之前定义
		 err := recover()
		 if err != nil {
		 	fmt.Println("func b error!")
		 }
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("func c")
}

func main() {
	a()
	b()
	c()
}
