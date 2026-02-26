package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	/*
	读取文件：
		1.打开文件
		2.读取文件
			file.Read([]byte)-->n,err
			从文件中开始读取数据，存入到byte切片中，返回值n是本次实际读取的数据量
				如果读取到文件末尾，n为0,err为EOF：end of file
		3.关闭文件
	 */
	//	 1、打开文件
	fileName := "./files/blockchain.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件错误", err.Error())
	} else {
		//2、读取文件
		bs := make([]byte, 1024*8, 1024*8)
		n := -1
		for {
			n , err = file.Read(bs)
			if n==0 || err == io.EOF {
				fmt.Println("读取文件结束！")
				break
			}
			fmt.Println(string(bs[:n]))
		}
	}

	//3、关闭文件
	file.Close()
}
