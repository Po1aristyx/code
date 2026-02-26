package main
import (
	"io/ioutil"
	"fmt"
	"os"
)
func main() {
	fileName1 := "./files/blockchain.txt"
	data, err := ioutil.ReadFile(fileName1)
	if err != nil {
		fmt.Println("文件打开异常", err.Error())
	} else {
		fmt.Println(string(data))
	}
	fileName2 := "./files/xyz.txt"
	s1 := "Steven陪你学区块链"
	err = ioutil.WriteFile(fileName2, []byte(s1), 0777)
	if err != nil {
		fmt.Println("写入文件异常", err.Error())
	} else {
		fmt.Println("文件写入OK!")
	}
	err = ioutil.WriteFile(fileName2 , data , os.ModePerm)
	if err != nil {
		fmt.Println("文件拷贝异常", err.Error())
	} else {
		fmt.Println("文件拷贝成功!")
	}
	dirName := "./"
	fileInfos , err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println("目录遍历异常", err.Error())
	} else {
		for i , v := range fileInfos {
			fmt.Println(i , v.Name() , v.IsDir() , v.Size() , v.ModTime())
		}
	}
	filename , err := ioutil.TempDir("./" , "temp")
	if err != nil {
		fmt.Println("创建目录失败" , err.Error())
	} else {
		fmt.Println(filename)
	}
	file1 , err := ioutil.TempFile(filename , "temp")
	if err != nil {
		fmt.Println("创建文件失败" , err.Error())
	} else {
		file1.WriteString("写入内容:" + file1.Name())
	}
	file1.Close()
}
