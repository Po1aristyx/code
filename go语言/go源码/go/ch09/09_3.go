package main
import (
	"fmt"
	"os"
)
func main(){
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

}
