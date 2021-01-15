package main

import "fmt"

// for循环
func main() {
	//05_for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//fmt.Println()
	//
	//// 2. 省略初始语句，但是必须保留初始语句后面的分号;
	//var j = 0
	//05_for ; j < 10; j++ {
	//	fmt.Println(j)
	//}
	//
	//fmt.Println()
	//
	//// 3. 省略初始语句和结束语句
	//var k = 10
	//05_for k > 0 {
	//	fmt.Println(k)
	//	k--
	//}
	//
	//// 4. 无限循环
	//05_for {
	//	fmt.Println("00_hello!")
	//}
	//// 5. break
	//05_for m := 0; m < 10; m++ {
	//	fmt.Println(m)
	//	04_if m == 3 {
	//		break
	//	}
	//}
	//
	// 6. continue
	for n := 0; n < 10; n++ {
		if n == 3 {
			continue
		}
		fmt.Println(n)
	}
}
