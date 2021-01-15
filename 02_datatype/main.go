package main

import (
	"fmt"
	"math"
	"strings"
)

// 基本数据类型

func main() {
	// 十进制打印为二进制 %b
	n := 10
	fmt.Printf("n: %d\n", n)
	fmt.Printf("n: %b\n", n)
	fmt.Println()

	// 八进制打印为二进制 %o
	m := 077
	fmt.Printf("m: %d\n", m)
	fmt.Printf("m: %o\n", m)
	fmt.Println()

	// 十六进制打印为二进制
	l := 0xff
	fmt.Printf("l: %d\n", l)
	fmt.Printf("l: %x\n", l)
	fmt.Println()

	// uint8 不能超过255
	var age uint8
	age = 255
	fmt.Println("age: ", age)
	fmt.Println()

	// float
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println()

	// bool
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	fmt.Println()

	// 字符串
	s1 := "00_hello"
	s2 := "hihi"
	s4 := "啦啦啦啦啦"
	s3 := s1 + s2 + s4
	fmt.Println(s3)
	s6 := fmt.Sprintf("%s - %s", s2, s4)
	fmt.Printf("s8: %s\n", s6)
	fmt.Println()

	// 多行字符串
	s5 := `
	多行字符串
	两个反引号之间的内容
	会原样
	输出
	`
	fmt.Println(s5)
	fmt.Printf("the length of s5: %d\n", len(s5))
	fmt.Println()

	// 分割字符串
	s7 := "how are you ? man do i do what you want"
	fmt.Println(strings.Split(s7, " "))
	fmt.Printf("%T\n", strings.Split(s7, " "))
	fmt.Println()

	// 判断是否包含
	fmt.Println(strings.Contains(s7, "men"))

	// 判断前缀
	fmt.Println(strings.HasPrefix(s7, "how"))

	// 判断后缀
	fmt.Println(strings.HasSuffix(s7, "man"))

	// 判断子串的位置
	fmt.Println(strings.Index(s7, "you"))

	// 判断最后子串出现的位置
	fmt.Println(strings.LastIndex(s7, "you"))

	//join
	s8 := []string{"are", "you", "ok", "?"}
	fmt.Printf("s8: %s\n", s8)
	fmt.Println(strings.Join(s8, " "))
}

