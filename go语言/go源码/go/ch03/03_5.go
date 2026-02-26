package main
import "fmt"
func main() {
   /* 定义局部变量 */
   grade := ""
   score := 78.5
   //思考：以下代码逻辑错误在哪里？
   //switch {
   //case score < 60:
   // grade = "E"
   //case score >= 60:
   // grade = "D"
   //case score >= 70:
   // grade = "C"
   //case score >= 80:
   // grade = "B"
   //case score >= 90:
   // grade = "A"
   //}
   switch { //switch后面省略不写，默认相当于：switch  true
   case score >= 90:
      grade = "A"
   case score >= 80:
      grade = "B"
   case score >= 70:
      grade = "C"
   case score >= 60:
      grade = "D"
   default:
      grade = "E"
   }
   fmt.Printf("你的等级是: %s\n", grade)
   fmt.Print("最终评价是: ")
   switch grade {
   case "A":
      fmt.Printf("优秀!\n")
   case "B":
      fmt.Printf("良好\n")
   case "C":
      fmt.Printf("中等\n")
   case "D":
      fmt.Printf("及格\n")
   default:
      fmt.Printf("差\n")
   }
}
