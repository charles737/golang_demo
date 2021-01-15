package main

import (
	"fmt"
	"time"
)

// work pool

func worker(id int, jobs <-chan int, results chan<- int)  {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500) // sleep500毫秒
		fmt.Printf("worker:%d finish job:%d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for j := 0; j < 3; j++ {
		go worker(j, jobs, results)
	}

	// 发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	// 输出结果
	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Printf("ret:%v\n", ret)
	}
	close(results)
}
