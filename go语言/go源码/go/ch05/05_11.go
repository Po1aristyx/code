package main
import (
   "fmt"
   "strconv"
)
func main() {
   //思考：使用哪种初始化切片的方式更高效？
   var sa []string
   //sa := make([]string , 0 , 20)
   printSliceMsg(sa)
   //当使用append追加元素到切片时，如果容量不够，go就会创建一个新的切片变量来储存元素。
   for i := 0; i < 15; i++ {
      sa = append(sa, strconv.Itoa(i))
      printSliceMsg(sa)
   }
   printSliceMsg(sa)
}
//打印输出格式化信息
func printSliceMsg(sa []string) {
   fmt.Printf("addr:%p \t len:%v \t cap:%d  \t  value:%v\n", sa, len(sa), cap(sa), sa)
}

