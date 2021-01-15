package main

import (
	"fmt"
	"math/rand"
)

// 在通道中，发送和接收都使用 <- 符号

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	elem2 := <-ch1
	elem3 := <-ch1
	fmt.Printf("the first element received from 17_channel ch1: %v\n", elem1)
	fmt.Printf("the second element received from 17_channel ch1: %v\n", elem2)
	fmt.Printf("the third element received from 17_channel ch1: %v\n", elem3)

	a := rand.Intn(1000)
	fmt.Printf("the value of a is: %d\n", a)
}
