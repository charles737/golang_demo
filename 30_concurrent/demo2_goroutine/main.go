package main

import (
	"fmt"
	"sync"
)

// go routine demo2_mutex

var wg sync.WaitGroup

func main() { // 开启主goroutine去执行main函数

	wg.Add(10000) // 计数器+1
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println("hello", i)
			wg.Done()
		}(i) // 既是匿名函数也是闭包
	}
	fmt.Println("main")
	wg.Wait() // 阻塞，等所有派出去的线程完成后收工（计数器为0）
}
