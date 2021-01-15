package main

import "fmt"

// 空接口
// 接口中没有定义任何需要实现的方法时，该接口就是一个空接口
// 任意类型都实现了空接口 --> 空接口变量可以存储任意值

// 空接口一般不需要提前定义
type xxx interface {
}

// 空接口的应用
// 1. 空接口类型作为函数的参数
// 2. 空接口可以作为map的value类型

func main() {
	//var x xxx // 写法一：定义一个空接口变量
	var x interface{} // 写法二：空接口一般不需要提前定义
	x = 100
	ret1 := x.(int)
	fmt.Println(ret1)
	//fmt.Println(x)
	x = "world"
	ret2 := x.(string)
	fmt.Println(ret2)
	//fmt.Println(x)
	x = true
	ret3, ok := x.(int) // 类型断言，猜的类型不对时，ok=false,ret3=string类型的零值
	if !ok {
		fmt.Println("不是布尔类型")
	} else {
		fmt.Println("是布尔类型", ret3)
	}
	//fmt.Println(x)

	// 使用switch进行类型断言
	switch v := x.(type) {
	case int:
		fmt.Printf("是int类型，value:%v\n", v)
	case bool:
		fmt.Printf("是bool类型，value:%v\n", v)
	case string:
		fmt.Printf("是字符串类型，value:%v\n", v)
	default:
		fmt.Printf("猜不到，value:%v\n", v)
	}

	fmt.Println()

	//var m = make(map[string]interface{}, 16)
	//m["name"] = "捷克"
	//m["age"] = 17
	//m["hobby"] = []string{"钢琴", "大提琴", "小提琴"}
	//fmt.Println(m)
}
