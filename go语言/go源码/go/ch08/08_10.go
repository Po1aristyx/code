package main
import "fmt"
func TestA() {
	fmt.Println("func TestA()")
}
func TestB(x int) {
	var a [100]int
	a[x] = 1000 //x值为11时，数组越界
}
func TestC() {
	fmt.Println("func TestC()")
}
func main() {
	TestA()
	TestB(101)//TestB()发生异常，中断程序
	TestC()
}
