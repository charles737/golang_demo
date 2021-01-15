package main

import "fmt"

func printBytes(s string) {
	fmt.Println("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func main() {
	name := "Hello World!"
	fmt.Printf("String: %s\n", name)
	printBytes(name)
}
