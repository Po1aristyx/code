package main
import "fmt"
func main() {
   a := [4]int{1, 2, 3, 4}
   fmt.Printf("1、变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//[1,2,3,4]
   fmt.Printf("数组型变量a内存地址是：%p \n\n", a)//可以获取到地址？  ❌
   //传值
   changeArrayVal(a)
   fmt.Printf("2、changeArrayVal函数调用后：变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//[99,2,3,4]  ❌
   //传引用
   changeArrayPtr(&a)
   fmt.Printf("3、changeArrayPtr函数调用后：变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//[99,250,3,4]  ❌
}
func changeArrayVal(a [4]int) {
   fmt.Printf("----------changeArrayVal函数内：值参数a的内存地址是：%p ，值为：%v \n", &a, a) //[1,2,3,4]
   fmt.Printf("----------changeArrayPtr函数内：值参数a的内存地址是：%p \n", a) //获取不到地址
   a[0] = 99
}
func changeArrayPtr(a *[4]int) {
   fmt.Printf("----------changeArrayPtr函数内：指针参数a的内存地址是：%p ，值为：%v \n", &a, a) //&[1,2,3,4]
   (*a)[1] = 250
}
