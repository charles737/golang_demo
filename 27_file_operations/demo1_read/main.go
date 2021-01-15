package main

import (
	"fmt"
	"io"
	"os"
)

// 文件操作

func readFromFile() {
	
}

func readAll() {
	// 打开文件
	fileObj, err := os.Open("./xxx.txt")
	if err != nil {
		fmt.Printf("file open filed, err:%v\n", err)
		return
	}

	// 关闭文件
	defer fileObj.Close()

	for {
		// 读取文件内容
		tmp := make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF { // EOF: end of file
			// 把当前读了多少字节的数据打印出来，然后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read from file filed, err:%v\n", err)
			return
		}
		fmt.Printf("read %d bytes from file.\n", n)
		fmt.Println(string(tmp[:n]))
	}
}

func main() {
	readAll()

}