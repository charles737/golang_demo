package main

import "fmt"

// 嵌套结构体

type Address struct {
	Province string
	City string
}

type Person6 struct {
	Name string
	Gender string
	Age int
	Address  // 嵌套另外一个结构体(匿名嵌套的结构体)
}

func main() {
	p1 := Person6{
		Name: "刘德华",
		Gender: "男",
		Age: 23,
		Address: Address{
			Province: "广东",
			City: "广州",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Name, p1.Age, p1.Gender, p1.Address)
	fmt.Println(p1.Address.Province, p1.Address.City)   // 通过嵌套的匿名结构体字段访问其内部字段
	fmt.Println(p1.Province)   // 直接访问匿名结构体中的字段
}