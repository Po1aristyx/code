package main

import (
	"os"
	"fmt"
)

func main() {
	/*
	文件操作：
		1.创建文件夹,如果文件夹存在，创建失败
			os.MKdir()
			os.MKdirAll()

		2.创建文件：如果文件存在，会覆盖
			os.Create()-->*file

		3.打开文件：
			os.Open(filename)
			os.OpenFile(filename,mode,perm)
		4.关闭文件：
			file.Close()
		5.删除：
			os.Remove()
			os.RemoveAll()
	 */
	//	 1、创建目录
	fileName1 := "./test1"
	err := os.Mkdir(fileName1, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s 目录创建成功！\n", fileName1)
	}

	fileName2 := "./test2/abc/xyz"
	err = os.MkdirAll(fileName2, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s 目录创建成功！\n", fileName2)
	}

	//2、创建文件.如果文件已经存在，则文件被覆盖
	fileName3 := "./test1/abc.txt"
	file1, err := os.Create(fileName3)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s 创建成功！%v \n", fileName3, file1)
	}

	//	3、打开文件
	file2, err := os.Open(fileName3)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s 打开成功！%v \n", fileName3, file2)
	}

	/*
	第一个参数：文件名称
	第二个参数：文件的打开方式
		O_RDONLY：只读模式(read-only)
		O_WRONLY：只写模式(write-only)
		O_RDWR：读写模式(read-write)
		O_APPEND：追加模式(append)
		O_CREATE：文件不存在就创建(create a new file if none exists.)
	第三个参数：文件的权限：文件不存在创建文件，需要指定权限
	 */
	fileName4 := "./test1/abc2.txt"
	file4, err := os.OpenFile(fileName4, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s 打开成功！%v \n", fileName4, file4)
	}

	//4、关闭文件，程序与文件的联系断开
	file4.Close()

	//	5、删除文件和目录
	fileName5 := "./test1"
	err = os.Remove(fileName5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s 删除成功！" , fileName5)
	}

	err = os.RemoveAll(fileName5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s 删除成功！" , fileName5)
	}
}
