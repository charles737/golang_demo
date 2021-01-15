package main

import "fmt"

// 常量
const pi = 3.1415
const e = 2.7

const (
	n1 = iota
	n2
	n3
	n4
	n5 = 100
	n6
	_
	n8 = iota
	n9
)
const n10 = iota

func main() {
	res := 2 * pi * e
	fmt.Println(res)
	fmt.Println(n1, n2, n3, n4, n5, n6, n8, n9, n10)
}
