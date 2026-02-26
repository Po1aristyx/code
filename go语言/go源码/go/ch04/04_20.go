package main
import "fmt"
func main() {
   a := []int{1, 2, 3, 4}
   fmt.Printf("1、变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//[1,2,3,4]
   fmt.Printf("切片型变量a内存地址是：%p \n\n", a)//可以获取到地址，类似：0xc420018080
   //传值
   changeSliceVal(a)
   fmt.Printf("2、changeSliceVal函数调用后：变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//[1,2,3,4]
   //传引用
   changeSlicePtr(&a)
   fmt.Printf("3、changeSlicePtr函数调用后：变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//[250,2,3,4]
}
func changeSliceVal(a []int) {
   fmt.Printf("----------changeSliceVal函数内：值参数a的内存地址是：%p ，值为：%v \n", &a, a) //[1,2,3,4]
   fmt.Printf("----------changeSlicePtr函数内：值参数a的内存地址是：%p \n", a)
   a[0] = 99
}
func changeSlicePtr(a *[]int) {
   fmt.Printf("----------changeSlicePtr函数内：指针参数a的内存地址是：%p ，值为：%v \n", &a, a) //&[1,2,3,4]
   (*a)[1] = 250
}

