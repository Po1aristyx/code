package main
import "fmt"
func main() {
   score := 88
   if score >= 90 {
      fmt.Println("优秀")
   } else if score >= 80 {
      fmt.Println("良好")
   } else if score >= 70 {
      fmt.Println("中等")
   } else if score >= 60 {
      fmt.Println("及格")
   } else if score < 60 {
      fmt.Println("不及格")
   }
}

