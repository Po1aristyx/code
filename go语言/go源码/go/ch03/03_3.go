package main
import "fmt"
func main() {
   if num := 10; num%2 == 0 {
      fmt.Println(num, "偶数")
   } else {
      fmt.Println(num, "奇数")
   }
}

