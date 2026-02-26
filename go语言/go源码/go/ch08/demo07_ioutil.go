package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	/*
	ioutil包：
		ReadFile()   //读取文件中的所有的数据，返回读取的字节数组
		WriteFile()  //向指定文件写入数据，如果文件不存在，则创建文件，写入数据之前清空文件
		ReadDir()    //读取一个目录下的子内容：子文件和子目录，但是仅有一层
		TempDir(）   //在当前目录下，创建一个以指定字符串为前缀的临时文件夹，并返回文件夹路径
		TempFile()   //在当前目录下，创建一个以指定字符串为前缀的文件，并以读写模式打开文件，并返回os.File指针对象
	 */
	//1、ReadFile()   //读取文件中的所有的数据，返回读取的字节数组
	fileName1 := "./files/blockchain.txt"
	data, err := ioutil.ReadFile(fileName1)
	if err != nil {
		fmt.Println("文件打开异常", err.Error())
	} else {
		fmt.Println(string(data))
	}

	//2、WriteFile()  //向指定文件写入数据，如果文件不存在，则创建文件，写入数据之前清空文件
	fileName2 := "./files/xyz.txt"
	s1 := "Steven陪你学区块链"
	err = ioutil.WriteFile(fileName2, []byte(s1), 0777)
	if err != nil {
		fmt.Println("写入文件异常", err.Error())
	} else {
		fmt.Println("文件写入OK!")
	}

	//3、文件拷贝
	err = ioutil.WriteFile(fileName2 , data , os.ModePerm)
	if err != nil {
		fmt.Println("文件拷贝异常", err.Error())
	} else {
		fmt.Println("文件拷贝成功!")
	}

	//4、ReadDir()    //读取一个目录下的子内容：子文件和子目录，但是仅有一层
	dirName := "./"
	fileInfos , err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("目录遍历异常", err.Error())
	} else {
		for i , v := range fileInfos {
			fmt.Println(i , v.Name() , v.IsDir() , v.Size() , v.ModTime())
		}
	}

	//5、TempDir(）   //在当前目录下，创建一个以指定字符串为前缀的临时文件夹，并返回文件夹路径
	filename , err := ioutil.TempDir("./" , "temp")
	if err != nil {
		fmt.Println("创建目录失败" , err.Error())
	} else {
		fmt.Println(filename)
	}

	//6、TempFile()   //在当前目录下，创建一个以指定字符串为前缀的文件，并以读写模式打开文件，并返回os.File指针对象
	file1 , err := ioutil.TempFile(filename , "temp")
	if err != nil {
		fmt.Println("创建文件失败" , err.Error())
	} else {
		file1.WriteString("写入内容:" + file1.Name())
	}
	file1.Close()
}
