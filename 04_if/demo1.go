package main

import "fmt"

// if判断
func main() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C ")
	}

	// if判断特殊写法
	if score := 100; score >= 90 {
		fmt.Println("A")
	} else if score >= 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C ")
	}
}
