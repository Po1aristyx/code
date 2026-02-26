package main
import "fmt"
func main() {
   var numbers = make([]int,3,5)
   fmt.Printf("%T\n" , numbers)
   fmt.Printf("len=%d cap=%d slice=%v\n",len(numbers),cap(numbers),numbers)
}

