package main
import  "fmt"
func main() {
   //定义数组
   arr := [3]int{1, 2, 3}
   //根据数组截取切片
   nums1 := arr[:]
   nums2 := arr[:]
   fmt.Println("arr=", arr)//[1 2 3]
   nums1[0] = 100
   fmt.Println("arr=", arr)//[100 2 3]
   nums2[1] = 200
   fmt.Println("arr=", arr)//[100 200 3]
}
