package main

import (
	"fmt"
	"sync"
)

// 多个goroutine并发操作全局变量x
// 错误示例

var (
	x int64
	wg sync.WaitGroup
)

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Printf("x:%d\n", x) // x:5968 不一定会得到10000
}
