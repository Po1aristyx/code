package main
import "fmt"
func main() {
   for i := 0; i < 5; i++ {
      fmt.Printf("i=%d \t", i)
      fmt.Println(add2(i))
   }
}
func add2(x int) int {
   sum := 0
   sum += x
   return sum
}

