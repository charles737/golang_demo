package main

import (
	"fmt"
	"sync"
)

// 多个goroutine访问全局变量x
// 正确示例：互斥锁

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
)

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 释放锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Printf("x:%d\n", x) // x:10000
}
