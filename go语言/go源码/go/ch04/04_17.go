package main
import "fmt"
const COUNT int = 4
func main() {
   a := [COUNT]string{"abc", "ABC", "123", "一二三"}
   i := 0
   //定义指针数组
   var ptr [COUNT]*string
   fmt.Printf("%T , %v \n", ptr, ptr)
   for i = 0; i < COUNT; i++ {
      //将数组中每个元素的地址赋值给指针数组
      ptr[i] = &a[i]
   }
   fmt.Printf("%T , %v \n", ptr, ptr)
   //获取指针数组中第一个值，其实就是一个地址
   fmt.Println(ptr[0])
   //根据数组元素的每个地址获取该地址所指向的元素的数值
   for i = 0; i < COUNT; i++ {
      fmt.Printf("a[%d] = %s \n", i, *ptr[i])
   }
}

