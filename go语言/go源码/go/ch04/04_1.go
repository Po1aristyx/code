package main
import "fmt"
/* 声明全局变量 */
var a1 int = 7
var b1 int = 9
func main() {
   /* main 函数中声明局部变量 */
   a1, b1, c1 := 10, 20, 0
   fmt.Printf("main()函数中 a1 = %d\n", a1) //10
   fmt.Printf("main()函数中 b1 = %d\n", b1) //20
   fmt.Printf("main()函数中 c1 = %d\n", c1) //0
   c1 = sum(a1, b1)
   fmt.Printf("main()函数中 c1 = %d\n", c1) //33
}
/* 函数定义-两数相加 */
func sum(a1, b1 int) (c1 int) {
   a1++
   b1 += 2
   c1 = a1 + b1
   fmt.Printf("sum() 函数中 a1 = %d\n", a1) //11
   fmt.Printf("sum() 函数中 b1 = %d\n", b1) //22
   fmt.Printf("sum() 函数中 c1 = %d\n", c1) //33
   return c1
}

