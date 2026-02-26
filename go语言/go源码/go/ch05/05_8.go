package main
import "fmt"
func main() {
   a := [4]float64{67.7, 89.8, 21, 78}
   b := []int{2, 3, 5}
   fmt.Printf("变量a —— 地址：%p ， 类型：%T，数值：%v，长度：%d \n", &a, a, a, len(a))
   fmt.Printf("变量b —— 地址：%p ， 类型：%T，数值：%v，长度：%d \n", &b, b, b, len(b))
   c := a
   d := b
   fmt.Printf("变量c —— 地址：%p ， 类型：%T，数值：%v，长度：%d \n", &c, c, c, len(c))
   fmt.Printf("变量d —— 地址：%p ， 类型：%T，数值：%v，长度：%d \n", &d, d, d, len(d))
   a[1] = 200
   fmt.Println("a=", a, "c=", c)
   d[0] = 100
   fmt.Println("b=", b, "d=", d)
}

