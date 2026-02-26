package main
import "fmt"
func main() {
   pos := adder()
   for i := 0; i < 10; i++ {
      fmt.Printf("i=%d \t", i)
      fmt.Println(pos(i))
   }
   fmt.Println("---------------------")
   for i := 0; i < 10; i++ {
      fmt.Printf("i=%d \t", i)
      fmt.Println(pos(i))
   }
}
func adder() func(int) int {
   sum := 0
   return func(x int) int {
      fmt.Printf("sum1=%d \t", sum)
      sum += x
      fmt.Printf("sum2=%d \t", sum)
      return sum
   }
}

