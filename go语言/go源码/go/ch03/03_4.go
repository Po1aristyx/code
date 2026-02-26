package main
import "fmt"
func main() {
   if score := 98; score >= 60 {
      if score >= 70 {
         if score >= 80 {
            if score >= 90 {
               fmt.Println("优")
            } else {
               fmt.Println("良")
            }
         } else {
            fmt.Println("中等")
         }
      } else {
         fmt.Println("及格")
      }
   } else {
      fmt.Println("不及格")
   }
}
