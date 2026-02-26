package main
import "fmt"
func main() {
   a := 10
   fmt.Printf("1、变量a的内存地址：%p ，值为：%v \n\n", &a, a) //10
   fmt.Printf("========int型变量a的内存地址：%p \n\n", a)  //？？？%!p
   //传值
   changeIntVal(a)
   fmt.Printf("2、changeIntVal函数调用之后：变量a的内存地址：%p ，值为：%v \n\n", &a, a) //10
   //传引用
   changeIntPtr(&a)
   fmt.Printf("3、changeIntPtr函数调用之后：变量a的内存地址：%p ，值为：%v \n\n", &a, a) //50
}
func changeIntVal(a int) {
   fmt.Printf("--------changeIntVal函数内：值参数a的内存地址：%p ，值为：%v \n", &a, a) //10
   a = 90
}
func changeIntPtr(a *int) {
   fmt.Printf("--------changeIntPtr函数内：指针参数a的内存地址：%p ，值为：%v \n", &a, a) //地址
   *a = 50
}

