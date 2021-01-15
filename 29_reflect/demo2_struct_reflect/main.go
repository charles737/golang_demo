package main

import (
	"fmt"
	"reflect"
)

// 结构体反射

type student struct {
	Name string `json:"name" ini:"s_name"`
	Score int `json:"score" ini:"s_score"`
}

func (s student) Study() string {
	msg := "it's time to study!"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "it's time to sleep!"
	fmt.Println(msg)
	return msg
}

// 根据反射去获取结构中方法的函数
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod()) // 取到变量的方法数量
	// 因为下面需要拿到具体的方法调用，所以使用的是值信息
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args) // 调用方法
	}
	// 通过方法名获取结构体的方法
	t.MethodByName("Sleep") // 返回 Method, bool
	methodObj := v.MethodByName("Sleep") // 返回 Value
	fmt.Println(methodObj.Type()) // func() string
}

func main() {
	stu1 := student{
		Name: "刘德华",
		Score: 89,
	}
	//// 通过反射获取结构体中的所有字段信息
	//t := reflect.TypeOf(stu1)
	//fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())
	//// 遍历结构体变量的所有字段
	//for i := 0; i < t.NumField(); i++ {
	//	// i : 结构体字段的索引
	//	// 根据结构体字段的索引去取字段
	//	fileObj := t.Field(i)
	//	fmt.Printf("name:%v type:%v tag:%v\n", fileObj.Name, fileObj.Type, fileObj.Tag)
	//	fmt.Printf("json:%v ini:%v\n", fileObj.Tag.Get("json"), fileObj.Tag.Get("ini"))
	//}
	//
	//// 根据名字去取结构体中的字段
	//fileObj2, ok := t.FieldByName("Name")
	//if ok{
	//	fmt.Printf("name:%v type:%v tag:%v\n", fileObj2.Name, fileObj2.Type, fileObj2.Tag)
	//}

	fmt.Println()

	printMethod(stu1)
}
