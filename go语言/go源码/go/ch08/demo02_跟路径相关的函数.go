package main

import (
	"path/filepath"
	"fmt"
	"path"
)

func main() {
	/*
	文件路径：
	1、绝对路径：absolute
		/Users/steven/Documents/go_project/files/dsa.png
	2、相对路径：relative 都是相当于当前的工程
		.当前目录
		..上一层
	*/

	/*（1）、判断是否是绝对路径filepath.IsAbs()
	（2）、获取相对路径filepath.Rel()
	（3）、获取绝对路径filepath.Abs()
	（4）、拼接路径path.Join(）*/
	filePath1 := "/Users/steven/Documents/goyard/files/dsa.png"
	filePath2 := "files/blockchain.txt"
	fmt.Println(filepath.IsAbs(filePath1)) //true
	fmt.Println(filepath.IsAbs(filePath2)) //false

	fmt.Println(filepath.Rel("/Users/steven/Documents/goyard", filePath1))
	fmt.Println(filepath.Abs(filePath1))
	fmt.Println(filepath.Abs(filePath2))

	fmt.Println(path.Join(filePath1, ".."))
	fmt.Println(path.Join(filePath1, "."))
	fmt.Println(path.Join("/Steven/Daniel", filePath2))

}
