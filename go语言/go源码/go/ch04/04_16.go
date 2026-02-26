package main
import (
   "fmt"
)
func main() {
   a := 58
   fmt.Println("函数调用之前a的值：", a)
   fmt.Printf("%T \n", a)
   fmt.Printf("%x \n", &a)
   //b := &a
   var b *int = &a
   change(b)
   fmt.Println("函数调用之后的a的值：", a)
}
func change(val *int) {
   *val = 15
}

