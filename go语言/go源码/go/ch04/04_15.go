package main
import (
   "fmt"
)
func main() {
   b := 3158
   a := &b
   fmt.Println("b 的地址：", a) //0xc420014050
   fmt.Println("*a 的值：", *a) //3158
   *a++
   fmt.Println("b 的新值：", b)//3159
}

