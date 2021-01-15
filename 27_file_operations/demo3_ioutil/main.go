package main

import (
	"fmt"
	"io/ioutil"
)

// ioutil
// 不适用于文件特别大的场景，会撑爆内存

func readByIoutil()  {
	// 读取文件
	content, err := ioutil.ReadFile("./xxx.txt")
	if err != nil {
		fmt.Printf("read by ioutil filed, err:%v\n", err)
		return
	}
	fmt.Println(string(content))

}

func main() {
	readByIoutil()
}
