package main
import (
	"fmt"
	"os"
)
func main(){
	fileName3 := "./test1/abc.txt"
	file1, err := os.Create(fileName3)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		fmt.Printf("%s 创建成功！%v \n", fileName3, file1)
	}
}