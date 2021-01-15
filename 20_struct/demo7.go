package main

import "fmt"

// 嵌套结构体的冲突

type Address7 struct {
	Province string
	City string
	UpdateTime string
}

type Email struct {
	Adr string
	UpdateTime string
}

type Person7 struct {
	Name string
	Gender string
	Age int
	Address7  // 嵌套另外一个结构体(匿名嵌套的结构体)
	Email
}

func main() {
	p1 := Person7{
		Name: "周杰伦",
		Gender: "男",
		Age: 34,
		Address7: Address7{
			Province: "广东",
			City: "广州",
			UpdateTime: "2020-03-12",
		},
		Email: Email{
			Adr: "234324@q234.com",
			UpdateTime: "2020-09-09",
		},
	}
	fmt.Printf("%#v\n", p1)
	//fmt.Println(p1.UpdateTime)   // 冲突：嵌套结构体中包含了多个同名的UpdateTime字段
	fmt.Println("Email-UpdateTime:", p1.Email.UpdateTime)
	fmt.Println("Address7-UpdateTime:", p1.Address7.UpdateTime)
}