package main

import "fmt"

func getIntChan() <-chan int {
	num := 50
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func main() {
	intchan2 := getIntChan()
	for elem := range intchan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}
}
