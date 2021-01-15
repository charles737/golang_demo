package main

import "fmt"

/*
两个goroutine, 两个channel
	1. 生成0～100的数字发送到 ch1
	2. 从 ch1 中取出数据并计算平方，把结果发送给 ch2
*/

// 单向通道
// <-chan 表示只能从通道中发送数据 send only
// chan<- 表示只能往通道中发送数据

// 生成0～100的数字发送到 ch1
func f1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

// 从 ch1 中取出数据并计算平方，把结果发送给 ch2
func f2(ch1 <-chan int, ch2 chan<- int) {
	// 从通道中取值的方式一
	for {
		tmp, ok := <- ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)

	// 从通道中取值的方式二
	for ret := range ch2 {
		fmt.Println(ret)
	}

}
