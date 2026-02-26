package main
import "fmt"
func main() {
   fmt.Println("1、------------------")
   //var numbers []int
   numbers := make([]int , 0 , 20)
   printSlices("numbers:", numbers)
   numbers = append(numbers, 0) //[0]
   printSlices("numbers:", numbers)
   /* 向切片添加一个元素 */
   numbers = append(numbers, 1) //[0 1]
   printSlices("numbers:", numbers)
   /* 同时添加多个元素 */
   numbers = append(numbers, 2, 3, 4, 5, 6, 7) //[0 1 2 3 4 5 6 7]
   printSlices("numbers:", numbers)
   fmt.Println("2、------------------")
   //追加一个切片
   s1 := []int{100, 200, 300, 400, 500, 600, 700}
   numbers = append(numbers, s1...)
   printSlices("numbers:", numbers)
   fmt.Println("3、------------------")
   //切片删除元素
   //删除第一个元素
   numbers = numbers[1:]
   printSlices("numbers:", numbers)
   //删除最后一个元素
   numbers = numbers[:len(numbers)-1]
   printSlices("numbers:", numbers)
   //删除中间一个元素
   a := int(len(numbers)/2)
   fmt.Println("中间数：" , a)
   numbers = append(numbers[:a] , numbers[a+1:]...)
   printSlices("numbers:", numbers)
   fmt.Println("4、========================")
   /* 创建切片 numbers1 是之前切片的两倍容量*/
   //numbers1 := make([]int, 0, (cap(numbers))*2)
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)
   /* 拷贝 numbers 的内容到 numbers1 */
   count := copy(numbers1, numbers)
   fmt.Println("拷贝个数：", count)
   printSlices("numbers1:", numbers1)
   numbers[len(numbers)-1] = 99
   numbers1[0] = 100
   /*numbers1与numbers两者不存在联系，numbers发生变化时，
   numbers1是不会随着变化的。也就是说copy方法是不会建立两个切片的联系的
    */
   printSlices("numbers1:", numbers1)
   printSlices("numbers:", numbers)
}
func printSlices(name string, x []int) {
   fmt.Print(name, "\t")
   fmt.Printf("addr:%p \t len=%d \t cap=%d \t slice=%v\n", x , len(x),cap(x),x)
}
