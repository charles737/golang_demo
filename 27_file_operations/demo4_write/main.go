package main

import (
	"fmt"
	"os"
)

// 文件写

func write() {
	// 打开文件
	fileObj, err := os.OpenFile("xxx.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file filed, err:%v\n", err)
		return
	}

	// 关闭文件
	defer fileObj.Close()

	// 写入文件
	str := "abcde"
	fileObj.Write([]byte(str)) // []byte 写入字节切片数据
	fileObj.WriteString("呵呵") // string 直接写入字符串数据
}

func main() {
	write()
}
