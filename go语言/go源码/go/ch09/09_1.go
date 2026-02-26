package main
import (
	"os"
	"fmt"
)
func main() {
	//绝对路径
	path := "D:/go/ch09/picture/soldier.jpg"
	printMessage(path)
	//相对路径
	path = "./picture/timg.jpg"
	printMessage(path)
}
func printMessage(filePath string){
	fileInfo , err := os.Stat(filePath)
	if err !=nil {
		fmt.Println("err:" , err.Error())
	} else {
		fmt.Printf("数据类型是：%T \n" , fileInfo)
		fmt.Println("文件名：",fileInfo.Name())
		fmt.Println("是否为目录：",fileInfo.IsDir())
		fmt.Println("文件大小：",fileInfo.Size())
		fmt.Println("文件权限：",fileInfo.Mode())
		fmt.Println("文件最后修改时间：",fileInfo.ModTime())
	}
}

