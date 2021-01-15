package main

import "fmt"

// 指针

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 100
}

func main() {
	a := 1
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)
}
