package main
import (
	"strconv"
	"fmt"
)
func main() {
	TestAtoi()
	TestParseInt()
	TestParseUint()
	TestParseFloat()
	TestParseBool()
}
func TestAtoi() {
	a, _ := strconv.Atoi("100")
	fmt.Printf("%T , %v \n", a, a+2)
	fmt.Println("---------------")
}
func TestParseInt() {
	num, _ := strconv.ParseInt("-4e00", 16, 64)
	fmt.Printf("%T , %v \n", num, num)
	num, _ = strconv.ParseInt("01100001", 2, 64)
	fmt.Printf("%T , %v\n", num, num)
	num, _ = strconv.ParseInt("-01100001", 10, 64)
	fmt.Printf("%T , %v\n", num, num)
	num, _ = strconv.ParseInt("4e00", 10, 64)
	fmt.Printf("%T , %v\n", num, num)
	fmt.Println("---------------")
}
func TestParseUint() {
	num, _ := strconv.ParseUint("4e00", 16, 64)
	fmt.Printf("%T , %v \n", num, num)
	num, _ = strconv.ParseUint("01100001", 2, 64)
	fmt.Printf("%T , %v\n", num, num)
	num, _ = strconv.ParseUint("-1100001", 10, 64)
	fmt.Printf("%T , %v\n", num, num)
	num, _ = strconv.ParseUint("4e00", 10, 64)
	fmt.Printf("%T , %v\n", num, num)
	fmt.Println("---------------")
}
func TestParseFloat() {
	pi := "3.1415926"
	num , _ := strconv.ParseFloat(pi , 64)
	fmt.Printf("%T , %v\n", num, num*2)
	fmt.Println("---------------")
}
func TestParseBool() {
	flag , _:=strconv.ParseBool("steven")
	fmt.Printf("%T , %v\n", flag, flag)
	fmt.Println("---------------")
}
