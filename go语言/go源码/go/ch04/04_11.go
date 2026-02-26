package main
import "fmt"
func main() {
   fmt.Println(factorial(5))
}
//通过递归实现阶乘
func factorial(n int) int {
   if n == 0 {
      return 1
   }
   return n * factorial(n-1)
}
//通过循环实现阶乘
func getMultiple(num int) (result int) {
   result = 1
   for i:=1; i<= num; i++ {
      result *= i
   }
   return
}

