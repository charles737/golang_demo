package main

import "fmt"

// 通道是引用类型，必须使用 make 初始化后才能使用
// (channel, map, slice) 必须使用 make 初始化
// 有缓冲区与无缓冲区通道区别：
// 有缓冲区通道就像是快递代收点，可以帮我们存储数据
// 无缓冲区通道就像是没有快递代收点，必须是快递员等你，亲自把快递交到你手里（一直等待，死锁，必须等到有变量来取通道内的值）
// 无缓冲区通道又称为同步通道
// 有缓冲区通道又称为异步通道

func main() {
	// 有缓冲区通道
	//var ch1 chan int
	ch1 := make(chan int, 1)
	ch1 <- 10
	lenCh1 := len(ch1)
	x := <- ch1
	lenCh2 := len(ch1) // 取通道中元素的数量
	capCh1 := cap(ch1) // 取到通道的容量
	fmt.Printf("x:%v\n", x)
	fmt.Printf("len of ch1:%v\n", lenCh1)
	fmt.Printf("len of ch1:%v\n", lenCh2)
	fmt.Printf("cap of ch1:%v\n", capCh1)
	close(ch1)
}
