package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "D:\\QQ数据\\go源码\\go\\ch09\\files\\blockchain.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件错误", err.Error())
	} else {
		bs := make([]byte, 1024*8, 1024*8)
		n := -1
		for {
			n, err = file.Read(bs)
			if n == 0 || err == io.EOF {
				fmt.Println("读取文件结束！")
				break
			}
			fmt.Println(string(bs[:n]))
		}
	}
	file.Close()
}
