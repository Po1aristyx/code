package main
import (
   "fmt"
   "math"
)
func main() {
   //调用函数，对每个元素进行求平方根操作
   arr := []float64{1, 9, 16, 25, 30}
   visit(arr, func(v float64) {
      v = math.Sqrt(v)
      fmt.Printf("%.2f \n", v)
   })
   //调用函数，对每个元素进行求平方操作
   visit(arr, func(v float64) {
      v = math.Pow(v , 2)
      fmt.Printf("%.0f \n", v)
   })
}
//定义一个函数，遍历切片元素，对每个元素进行处理
func visit(list []float64, f func(float64)) {
   for _, value := range list {
      f(value)
   }
}
