package main
import (
	"fmt"
	"os"
)
func main(){
	fileName5 := "./test1"
	err := os.Remove(fileName5)
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
