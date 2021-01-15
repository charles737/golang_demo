package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 准备几个通道
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	// 随机选择一个通道，并向它发送元素值
	// 必须使用rand.Seed()初始化rand.Intn()和rand包的其他函数使用的全局Source
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行
	select {
	case <-intChannels[0]:
		fmt.Println("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Println("The second candidate case is selected.")
	case <-intChannels[2]:
		fmt.Println("The third candidate case is selected.")
	default:
		fmt.Println("No candidate case is selected!")
	}
}