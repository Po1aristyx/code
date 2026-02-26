package main
import (
   "fmt"
)
func main() {
   //传进n个参数
   sum, avg, count := GetScore(90, 82.5, 73, 64.8)
   fmt.Printf("学员共有%d门成绩，总成绩为：%.2f，平均成绩为：%.2f", count, sum, avg)
   fmt.Println()
   // 传切片作为参数
   scores := []float64{92, 72.5, 93, 74.5, 89, 87, 74}
   sum, avg, count = GetScore(scores...)
   fmt.Printf("学员共有%d门成绩，总成绩为：%.2f，平均成绩为：%.2f", count, sum, avg)
}
//累加求和，参数个数不定，参数个数从0-n
func GetScore(scores ...float64) (sum, avg float64, count int) {
   for _, value := range scores {
      sum += value
      count++
   }
   avg = sum / float64(count)
   return
}

