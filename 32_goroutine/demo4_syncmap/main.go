package main

import (
	"fmt"
	"sync"
)

// sync.Map 是并发安全的Map
// Go语言中的Map是非并发安全的，如下示例

var wg sync.WaitGroup
var m = make(map[int]int)
var m2 = sync.Map{}

func get(key int) int {
	return m[key]
}

func set(key, value int) {
	m[key] = value
}

//func main() {
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go func(i int) {
//			set(i, i + 100) // 设置键值对
//			fmt.Printf("k:%v v:%v\n", i, get(i)) // 打印键值对
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

func main() {
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i + 100) // 设置键值对
			value, _ := m2.Load(i)
			fmt.Printf("k:%v v:%v\n", i, value) // 打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}
