package main
import "fmt"
func main() {
   myfunc := Counter()
   //fmt.Printf("%T\n", myfunc)
   fmt.Println("myfunc", myfunc)
   /* 调用 myfunc 函数，i 变量自增 1 并返回 */
   fmt.Println(myfunc())
   fmt.Println(myfunc())
   fmt.Println(myfunc())
   /* 创建新的函数 nextNumber1，并查看结果 */
   myfunc1 := Counter()
   fmt.Println("myfunc1", myfunc1)
   fmt.Println(myfunc1())
   fmt.Println(myfunc1())
}
//计数器.闭包函数
func Counter() func() int {
   i := 0
   res := func() int {
      i += 1
      return i
   }
   //fmt.Printf("%T , %v \n" , res , res)  //func() int , 0x1095af0
   fmt.Println("Counter中的内部函数:", res) //0x1095af0
   return res
}

