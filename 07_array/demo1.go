package main

import "fmt"

// 07_array
func main() {
	//var a [3]int
	//var b [6]int
	//fmt.Println(a)
	//fmt.Println(b)

	//// 2. 数组的初始化
	//// 1) 定义时使用初始值列表的方式初始化
	//var cityArray = [4]string{"北京", "天京", "东京", "南京"}
	//fmt.Printf("cityArray: %s\n", cityArray)
	//fmt.Println(cityArray[0])
	//fmt.Println(cityArray[3])
	//// 2) 编译器推导数组的长度
	//var boolArray = [...]bool{true, false, false, true, true}
	//fmt.Println(boolArray)
	//// 3) 使用索引值的方式来初始化
	//var longArray = [...]string{1: "Python", 3: "Java", 8: "Golang", 2: "C++", 5: "PHP"}
	//fmt.Println(longArray)
	//fmt.Printf("%T\n", longArray)

	//// 3. 数组的遍历
	//var cityArray = [4]string{"北京", "天京", "东京", "南京"}
	//// 1) 使用for循环来遍历
	//05_for i := 0; i < len(cityArray); i++ {
	//	fmt.Printf("第%d个城市是: %s\n", i, cityArray[i])
	//}
	//// 2) 05_for range 遍历
	//05_for _, city := range cityArray {
	//	fmt.Println(city)
	//}

	//// 4. 二维数组
	//cityArray := [3][2]string{
	//	{"北京", "天京"},
	//	{"东京", "南京"},
	//	{"纽约", "墨尔本"},
	//}
	//fmt.Println(cityArray)
	//fmt.Println(cityArray[2][0])
	//// 二维数组的遍历
	//05_for _, v1 := range cityArray {
	//	05_for _, v2 := range v1 {
	//		fmt.Println(v2)
	//	}
	//}

	// 数组是值类型
	//x := [3]int{1, 2, 3}
	//fmt.Println(x)
	//f1(x)
	//fmt.Println(x)
	x := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
}

func f1(a [3][2]int) {
	a[0][0] = 100
	fmt.Println(a)
}