package main

import "fmt"

var n int

func main() {
	n = 200
	fmt.Println("Hello!")
	fmt.Printf("The value of n is %d\n", n)

	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)

	var name = "mi"
	var age = 23
	fmt.Println(name, age)

	name1 := "sfdjasl"
	age1 := 345
	fmt.Println(name1, age1)
}
