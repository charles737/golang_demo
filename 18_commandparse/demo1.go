package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()   // 解析命令参数，把它们的值赋给相应的变量
	fmt.Printf("Hello, %s!\n", name)
}