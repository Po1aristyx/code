package main
import "fmt"
type Teacher struct {
   name    string
   age     int
   married bool
   sex     int8
}
func main() {
   a := Teacher{"Steven", 35, true, 1}
   fmt.Printf("1、变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//{Steven 35 true 1}
   fmt.Printf("struct型变量a内存地址是：%p \n\n", a)//可以获取到地址？
   //传值
   changeStructVal(a)
   fmt.Printf("2、changeArrayVal函数调用后：变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//{Steven 35 true 1}
   //传引用
   changeStructPtr(&a)
   fmt.Printf("3、changeArrayPtr函数调用后：变量a的内存地址是：%p ，值为：%v \n\n", &a, a)//
}
func changeStructVal(a Teacher) {
   fmt.Printf("----------changeArrayVal函数内：值参数a的内存地址是：%p ，值为：%v \n", &a, a) //
   fmt.Printf("----------changeArrayPtr函数内：值参数a的内存地址是：%p \n", a) //获取不到地址?
   a.name = "Josh"
   a.age = 29
   a.married = false
}
func changeStructPtr(a *Teacher) {
   fmt.Printf("----------changeArrayPtr函数内：指针参数a的内存地址是：%p ，值为：%v \n", &a, a) //{Daniel 20 false 1}
   (*a).name = "Daniel"
   (*a).age = 20
   (*a).married = false
}

