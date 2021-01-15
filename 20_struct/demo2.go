package main

import "fmt"

// 结构体指针
type person1 struct {
	name, city string
	age int
}

func main() {
	// 结构体指针
	var p2 = new(person1)
	fmt.Printf("p2: %T\n", p2)   // p2: *main.person1
	//(*p2).name = "鸿"
	//(*p2).city = "东京"
	//(*p2).age = 23
	//fmt.Printf("%#v\n",p2)

	// go 语法糖
	p2.name = "鸿"
	p2.city = "东京"
	p2.age = 23
	fmt.Printf("%#v\n",p2)

	fmt.Println()

	p3 := person1{
		name: "刘德华",
		age: 32,
		city: "东京",
	}
	fmt.Printf("p3: %T\n", p3)
	fmt.Printf("p3: %#v\n", p3)

	fmt.Println()

	// 取结构体的地址进行实例化
	p4 := &person1{}
	p4.age = 23
	p4.city = "东京"
	p4.name = "公司"
	fmt.Printf("p4: %T\n", p4)
	fmt.Printf("p4: %#v\n", p4)
}
