package main

import "fmt"

// 为什么需要接口

// 接口不管你是什么类型，只管你要实现什么方法
// 定义一个类型，一个抽象的类型，只要实现了say()方法的类型都称为sayer类型
type sayer interface {
	say()
}

type dog struct {}

// 这里是方法，不是函数
// 函数和方法的区别在于：方法前面有接收者(d dog)
func (d dog) say() {
	fmt.Println("汪汪汪~")
}

type cat struct {}

func (c cat) say() {
	fmt.Println("喵喵喵~")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("呵呵呵~")
}

// touch 的函数
func touch(arg sayer) {
	// 不管传进来的是什么，都要执行touch，而touch就会执行say方法
	arg.say()
}

func main() {
	c1 := cat{}
	touch(c1)

	d1 := dog{}
	touch(d1)

	p1 := person{
		name: "周星驰",
	}
	touch(p1)

	// 用一个接口变量存储所有实现了这个接口的对象
	var s sayer
	p2 := person{
		name: "周杰伦",
	}
	s = p2
	fmt.Println(s)
}