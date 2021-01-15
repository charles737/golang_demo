package main

import (
	othername "demo/25_package/calc" // 给导入的包起别名
	"fmt"
)

// Go 语言中不允许导入包而不使用
// Go 语言中不允许循环引用包

// init 函数多用来做一些初始化的操作:日志初始化，加载配置文件
func init() {
	fmt.Println("main.init()")
}

func main() {
	retAdd := othername.Add(10, 20)
	fmt.Println("add:", retAdd)
	name := othername.Name
	fmt.Println("main():", name)
}

