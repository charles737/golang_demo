package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文件操作

// read by bufio
func readByBufio() {
	// 打开文件
	fileObj, err := os.Open("./xxx.txt")
	if err != nil {
		fmt.Printf("file open filed, err:%v\n", err)
		return
	}

	// 关闭文件
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			// 把当前读了多少字节的数据打印出来，然后退出
			fmt.Printf(line)
			return
		}
		if err != nil {
			fmt.Printf("read file by bufio filed, err:%v\n", err)
			return
		}
		fmt.Printf(line)
	}
}

func main() {
	readByBufio()

}