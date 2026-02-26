package main

import (
	"os"
	"fmt"
)

func main() {
	//1、打开文件
	file, err := os.OpenFile("./files/abc.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	//3、关闭文件
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件异常", err.Error())
	} else {
		//2、写入文件
		n, err := file.Write([]byte("abcde12345"))
		if err != nil {
			fmt.Println("写入文件异常", err.Error())
		} else {
			fmt.Println("写入ok：" , n)
		}

		n, err = file.WriteString("中国字")
		if err != nil {
			fmt.Println("写入文件异常", err.Error())
		} else {
			fmt.Println("写入ok：" , n)
		}
	}
}
