package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// newwriter

func writeByBufio() {
	// 打开文件
	fileObj, err := os.OpenFile("xxx.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file filed, err:%v\n", err)
		return
	}

	// 关闭文件
	defer fileObj.Close()

	// 写入文件
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("呵呵") // 将内容写入缓冲区
	writer.Flush() // 将缓冲区的内容写入磁盘
}

func writeByIoutil() {
	str := "嘿嘿，呵呵，嗯嗯"
	err := ioutil.WriteFile("./xxx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
}

func main() {
	//writeByBufio()
	writeByIoutil()
}
