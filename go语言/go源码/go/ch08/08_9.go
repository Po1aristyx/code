package main
import "fmt"
func TestA() {
	fmt.Println("func TestA()")
}
func TestB() {
	panic("func TestB(): panic")
}
func TestC() {
	fmt.Println("func TestC()")
}
func main() {
	TestA()
	TestB()//TestB()发生异常，中断程序
	TestC()
}
