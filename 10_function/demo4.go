package main

import (
	"fmt"
	"strings"
)

// 闭包 + 匿名函数
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	// 闭包 + 函数 + 外部变量引用
	r := makeSuffixFunc(".txt")
	ret := r("朱古力")
	fmt.Println(ret)
}