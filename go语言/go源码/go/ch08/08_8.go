package main
import (
	"fmt"
)
func main() {
	str := "Steven欢迎大家学习区块链，123开始吧"
	fmt.Printf("原始字符串:\n %s\n", str)
	fmt.Println("翻转后字符串: ")
	ReverseString(str)

}
//利用defer堆栈执行字符串倒序排列
func ReverseString(str string) {
	for _, v := range []rune(str) {
		defer fmt.Printf("%c" , v)
	}
}
