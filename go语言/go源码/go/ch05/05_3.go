package main
import "fmt"
func main() {
   /* 数组 - 5 行 2 列*/
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
   fmt.Println(len(a))
   fmt.Println(len(a[0]))
   /* 输出数组元素 */
   for  i := 0; i < len(a); i++ {
      for j := 0; j < len(a[0]); j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }
}

