package main
import (
	"fmt"
	"strconv"
)
func main() {
	TestItoa()
	TestFormatInt()
	TestFormatUint()
	TestFormatFloat()
	TestFormatBool()
}
func TestItoa() {
	s := strconv.Itoa(199)
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))
	fmt.Println("----------------")
}
func TestFormatInt() {
	s := strconv.FormatInt(-19968, 16)//4e00
	s = strconv.FormatInt(-40869, 16)//9fa5
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))
	fmt.Println("----------------")
}
func TestFormatUint() {
	s := strconv.FormatUint(19968, 16)//4e00
	s = strconv.FormatUint(40869, 16)//9fa5
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))
	fmt.Println("----------------")
}
func TestFormatFloat() {
	s := strconv.FormatFloat(3.1415926 , 'g' , -1 , 64)
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))
	fmt.Println("----------------")
}
func TestFormatBool() {
	s := strconv.FormatBool(true)
	s = strconv.FormatBool(false)
	fmt.Printf("%T , %v  , 长度：%d \n", s, s, len(s))
	fmt.Println("----------------")
}

