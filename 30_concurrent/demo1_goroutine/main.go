package main

import (
	"fmt"
	"sync"
)

// go routine demo

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello!", i)
	wg.Done() // 通知wg把计数器-1
}

func main() { // 开启主goroutine去执行main函数

	wg.Add(10000) // 计数器+1
	for i := 0; i < 10000; i++ {
		go hello(i) // 开启一个goroutine去执行hello函数
	}
	fmt.Println("main")
	//time.Sleep(time.Second) // 一般仅做示例，不在实际项目中使用
	wg.Wait() // 阻塞，等所有派出去的线程完成后收工（计数器为0）
}
