package main

import (
	"fmt"
	"reflect"
)

// reflect

type Cat struct {}

type Dog struct {}

func reflectType(x interface{}) {
	// 我是不知道别人调用我这个函数的时候会传递进来什么类型的变量
	// 方式一：通过类型断言
	// 方式二：借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
	fmt.Printf("%T\n", obj) // *reflect.rtype
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v, %T\n", v, v)
	// 如何得到一个传入时候的类型的变量
	// Kind() 拿到值对应的类型的种类
	k := v.Kind()
	fmt.Printf("k:%v\n", k)
	switch k {
	case reflect.Float32:
		// 把反射取到的值转换成float32类型的变量
		ret := float32(v.Float())
		fmt.Printf("%v %T\n", ret, ret)
	case reflect.Int32:
		// 把反射取到的值转换成int32类型的变量
		ret := int32(v.Int())
		fmt.Printf("%v %T\n", ret, ret)
	}
}

func reflectSetValue(x interface{})  {
	v := reflect.ValueOf(x)
	// Elem() 用来根据指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.SetFloat(2.123)
	}
}

func main() {
	//// 基本数据类型
	//var a float32 = 1.234
	//reflectType(a) // float32 float32 float32
	//var b int8 = 10
	//reflectType(b) // int8 int8 int8
	//
	//fmt.Println()
	//
	//// 结构体类型
	//var c Cat
	//reflectType(c) // main.Cat Cat struct
	//var d Dog
	//reflectType(d) // main.Dog Dog struct
	//
	//fmt.Println()
	//
	//// slice 获取不到obj.Name，返回空
	//var e []int
	//reflectType(e) // []int  slice
	//var f []string
	//reflectType(f) // []string  slice
	//
	//fmt.Println()
	//
	// valueOf
	var aa int32 = 100
	reflectValue(aa)
	fmt.Println()

	var bb float32 = 1.234
	reflectValue(bb)
	//
	//// set value
	//var aaa int32 = 10
	//fmt.Println(aaa)
	//reflectSetValue(&aaa)
	//fmt.Println(aaa)
}
