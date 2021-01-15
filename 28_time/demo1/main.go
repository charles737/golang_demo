package main

import (
	"fmt"
	"time"
)

// time

func main() {
	now := time.Now() // 时间对象
	fmt.Println(now)

	year := now.Year()
	fmt.Println(year)

	month := now.Month()
	fmt.Println(month)

	day := now.Day()
	fmt.Println(day)

	hour := now.Hour()
	fmt.Println(hour)

	minute := now.Minute()
	fmt.Println(minute)

	second := now.Second()
	fmt.Println(second)

	fmt.Println()

	// 时间戳
	// 从1970年1月1日到现在的秒数
	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()
	fmt.Println(timeStamp1)
	fmt.Println(timeStamp2)

	fmt.Println()

	// 将时间戳转换为具体的时间格式
	// 1609308484 + 3600 = 1609312084 当前时间戳 + 一个小时
	t := time.Unix(1609312084, 0)
	fmt.Println(t)

	fmt.Println()

	//// 时间间隔
	//n := 5
	//// type Duration int64
	//// Duration 是基于int64的自定义类型
	//time.Sleep(time.Duration(n) * time.Second)
	//fmt.Println("over!")

	// Add
	// now + 1hour
	fmt.Println(now)
	t1 := now.Add(time.Hour)
	fmt.Println(t1)

	fmt.Println()

	// Sub
	t2 := t1.Sub(now)
	fmt.Println(t2)

	fmt.Println()

	//// 定时器
	//for tmp := range time.Tick(time.Second) {
	//	fmt.Println(tmp)
	//}

	// 时间格式化
	// 其他语言： Y   m  d  H  M  S
	// Golang：2006 01 02 15 04 05
	t3 := now.Format("2006-01-02 15:04:05")
	fmt.Println(t3)
	t4 := now.Format("2006-01-02")
	fmt.Println(t4)
	t5 := now.Format("2006-01-02 03:04:05 PM")
	fmt.Println(t5)

	fmt.Println()

	// 解析字符串类型的时间
	timeStr := "2020/12/12 15:00:00"
	timeObj1, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil {
		fmt.Printf("parse timeStr failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj1)
	// 1. 拿到时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2. 根据时区去解析一个字符串格式的时间
	timeObj2, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Printf("parse timeStr failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj2) // 东8区

}
