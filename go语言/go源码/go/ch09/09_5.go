package main
import (
	"fmt"
	"os"
)
func main(){
	fileName3 := "./test1/abc.txt"
	file2, err := os.Open(fileName3)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s 打开成功！%v \n", fileName3, file2)
	}
	fileName4 := "./test1/abc2.txt"
	file4, err := os.OpenFile(fileName4, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s 打开成功！%v \n", fileName4, file4)
	}
	file2.Close()
	file4.Close()
}
