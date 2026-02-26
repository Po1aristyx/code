package main
import (
	"os"
	"fmt"
)
func main() {
	file, err := os.OpenFile("./files/abc.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("打开文件异常", err.Error())
	} else {
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
