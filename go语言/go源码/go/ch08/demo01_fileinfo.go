package main

import (
	"os"
	"fmt"
)

func main() {
	//绝对路径
	fileInfo , err := os.Stat("/Users/steven/Documents/goyard/files/dsa.png")
	fileInfo , err = os.Stat("/Users/steven/Documents/goyard/files")

	//相对路径
	fileInfo , err = os.Stat("./files/forbiddencity.docx")
	fileInfo , err = os.Stat("../node_demo/n01_helloworld.js")


	if err !=nil {
		fmt.Println("err:" , err.Error())
	} else {
		fmt.Printf("%T \n" , fileInfo)
		fmt.Printf("%v \n" , fileInfo)
		//文件名
		fmt.Println(fileInfo.Name())
		//是否是目录
		fmt.Println(fileInfo.IsDir())
		//文件尺寸大小
		fmt.Println(fileInfo.Size())
		//mode 权限
		fmt.Println(fileInfo.Mode())
		//文件最后修改时间
		fmt.Println(fileInfo.ModTime())
	}

}
