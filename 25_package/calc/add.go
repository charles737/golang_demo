package calc

import (
	"demo/25_package/snow"
	"fmt"
)

// 标识符首字母大写表示对外可见
// 通常不对外可见的标识符都是首字母小写

// Name 定义一个全局变量
var Name = "清迈"

// Add 定义一个计算两个int类型数据和的函数
func Add(x, y int) int {
	snow.Snow()
	retSub := Sub(x, y)
	fmt.Println("sub:", retSub)
	return x + y
}

// init函数在包导入的时候自动执行
// init函数没有参数，返回值
func init() {
	fmt.Println("calc.init()")
	fmt.Println("init():", Name) // init 是在全局声明之后执行的
}