package main
import (
	"fmt"
	"unicode/utf8"
)
func main() {
	s := "我爱Go语言"
	fmt.Println("字节长度", len(s))
	fmt.Println("--------------")
	//for ... range遍历字符串
	len := 0
	for i, ch := range s {
		fmt.Printf("%d:%c ", i, ch)
		len++
	}
	fmt.Println("\n字符串长度", len)
	fmt.Println("--------------")
	//遍历所有字节
	for i, ch := range []byte(s) {
		fmt.Printf("%d:%x ", i, ch)
	}
	fmt.Println()
	fmt.Println("--------------")
	//遍历所有字符
	count := 0
	for i, ch := range []rune(s) {
		fmt.Printf("%d:%c ", i, ch)
		count++
	}
	fmt.Println()
	fmt.Println("字符串长度" , count)
	fmt.Println("字符串长度" , utf8.RuneCountInString(s))
}