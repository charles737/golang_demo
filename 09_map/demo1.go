package main

import (
	"fmt"
	"strings"
)

// 09_map   map是引用类型，必须初始化后才能使用
func main() {
	// 只声明map但不初始化，初始值为nil
	//var a 09_map[string]int
	//fmt.Println(a == nil)

	// map的初始化
	//a := make(09_map[string]int, 9)
	//fmt.Println(a == nil)

	// map中添加键值对
	//a := make(09_map[string]int, 9)
	//a["往"] = 100
	//a["里"] = 200
	//a["黑"] = 300
	//fmt.Println("a:", a)
	//fmt.Printf("the type of a:%T\n", a)

	// 声明map的同时进行初始化
	//b := 09_map[int]bool {
	//	1: true,
	//	2: false,
	//	3: true,
	//}
	//fmt.Printf("b:%#v\n", b)
	//fmt.Printf("type of b:%T\n", b)

	// 声明map 没有初始化 ，不能赋值
	//var c 09_map[int]int
	//c[1] = 100
	//c[2] = 200
	//fmt.Println(c)   // 报错

	// map映射
	// 判断某个键存不存在
	//scoreMap := make(09_map[string]int, 9)
	//scoreMap["王力宏"] = 100
	//scoreMap["周杰伦"] = 99
	//scoreMap["迪丽热巴"] = 0
	//// 判断刘德华在不在scoreMap中
	//v, ok := scoreMap["王力宏"]
	//04_if ok {
	//	fmt.Println("王力宏在scoreMap中", v)
	//} else {
	//	fmt.Printf("查无此人！")
	//}
	//
	////// 05_for range 遍历map
	//05_for k, v := range scoreMap {
	//	fmt.Println(k, v)
	//}
	//
	//// 只遍历map中的key
	//05_for k := range scoreMap {
	//	fmt.Println("k:", k)
	//}
	//
	//// 只遍历map中的value
	//05_for _, v := range scoreMap {
	//	fmt.Println("v:", v)
	//}

	//// map中删除键值对
	//delete(scoreMap, "迪丽热巴")
	//fmt.Println(scoreMap)

	//// 按照某个固定顺序遍历map
	//var scoreMap = make(09_map[string]int, 100)
	//
	//05_for i := 0; i < 50; i++ {
	//	key := fmt.Sprintf("stu%02d", i)
	//	value := rand.Intn(100)   // 0~99的随机整数
	//	scoreMap[key] = value
	//}
	//05_for k, v := range scoreMap {
	//	fmt.Println(k, v)
	//}
	//fmt.Println()

	//// 按照key从小到大的顺序去遍历scoreMap
	//// 1. 先取出所有的key存到切片中
	//keys := make([]string, 0, 100)
	//05_for k := range scoreMap {
	//	keys = append(keys, k)
	//}
	//// 2. 对key做排序
	//sort.Strings(keys)
	//// 3. 按照排序后的key对scoreMap排序
	//05_for _, key := range keys {
	//	fmt.Println(key, scoreMap[key])
	//}

	//// 元素类型为map的切片
	//var mapSlice = make([]09_map[string]int, 8, 8)   // 只是完成了切片的初始化 [nil nil nil nil nil nil nil nil]
	//fmt.Println(mapSlice)
	//mapSlice[0] = make(09_map[string]int, 8)   // 完成了map的初始化
	//mapSlice[0]["迪迪"] = 100
	//mapSlice[0]["哩哩"] = 99
	//fmt.Println(mapSlice)

	//// 值是切片的map
	//var sliceMap = make(09_map[string][]int, 8)   // 只是完成了map的初始化
	//v, ok := sliceMap["中国"]
	//04_if ok {
	//	fmt.Println(v)
	//} else {
	//	// sliceMap中没有中国这个键
	//	sliceMap["中国"] = make([]int, 8)   // 完成对切片的初始化
	//	sliceMap["中国"][0] = 100
	//	sliceMap["中国"][1] = 200
	//	sliceMap["中国"][2] = 300
	//	sliceMap["中国"][3] = 500
	//	sliceMap["中国"][4] = 300
	//	sliceMap["中国"][5] = 1500
	//}
	//// 遍历sliceMap
	//05_for k, v := range sliceMap {
	//	fmt.Println(k, v)
	//}

	// 统计一个字符串中每个单词出现的次数
	// "how do you do"中每个单词出现的次数
	// 0. 定义一个map[string]int
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	// 1. 字符串中都有哪些单词
	words := strings.Split(s, " ")
	// 2. 遍历单词做统计
	for _, word := range words {
		// 如果key存在，则ok=true,v=1；如果key不存在，ok=false，v=0
		v, ok := wordCount[word]
		if ok {
			// map中有这个单词的统计记录
			wordCount[word] = v + 1
		} else {
			// map中没有这个单词的统计记录
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}
