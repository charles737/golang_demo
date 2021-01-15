package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁
// 互斥锁的执行效率：1.407701478s
// 读写锁的执行效率：124.019653ms

var (
	x int64
	wg sync.WaitGroup // 实现并发任务的同步
	lock sync.Mutex
	rwLock sync.RWMutex
)

func read() {
	//lock.Lock()
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock.RUnlock()
	wg.Done()
}

func write() {
	//lock.Lock()
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 10)
	//lock.Unlock()
	rwLock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start))
}
