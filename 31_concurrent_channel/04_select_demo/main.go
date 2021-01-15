package main

import "fmt"

// select

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <- ch:
			fmt.Printf("x:%d\n", x)
		case ch<- i:
		default:
			fmt.Println("default...")
		}
	}
	close(ch)
}
