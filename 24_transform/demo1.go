package main

import "fmt"

func main() {
	//var srcInt = int16(-255)
	//dstInt := int8(srcInt)
	//fmt.Printf(string(dstInt))

	//var str = [3]string{"a", "b", "c"}
	//fmt.Print(str)
	//
	//s1 := make([]int, 5)
	//fmt.Printf("The length of s1: %d\n", len(s1))
	//fmt.Printf("The capacity of s1: %d\n", cap(s1))
	//fmt.Printf("The value of s1: %d\n", s1)
	//s2 := make([]int, 5, 8)
	//fmt.Printf("The length of s2: %d\n", len(s2))
	//fmt.Printf("The capacity of s2: %d\n", cap(s2))
	//fmt.Printf("The value of s2: %d\n", s2)

	aMap := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	k := "three"
	v, ok := aMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}
}
