package main
import (
	"fmt"
	"time"
)
func main() {
	//1.双向通道
	ch1 := make(chan string)
	go fun1(ch1)
	data := <-ch1
	fmt.Println("main，接受到数据：", data)//我是Steven老师
	ch1 <- "Go语言好学么？"
	ch1 <- "Go语言好学么???？"
	go fun2(ch1)
	go fun3(ch1)
	time.Sleep(1 * time.Second)
	fmt.Println("main over!")
}
func fun1(ch1 chan string) {
	ch1 <- "我是Steven老师"
	data := <-ch1
	data2 := <-ch1
	fmt.Println("回应：", data, data2)
}
//功能：只有写入数据
func fun2(ch1 chan<- string) {
	//只能写入
	ch1 <- "How are you?"
	//data := <- ch1
	//<- ch1 //invalid operation: <-ch1 (receive from send-only type chan<- string)
}
//功能：只有读取数据
func fun3(ch1 <-chan string) {
	data := <-ch1
	fmt.Println("只读：", data)
	//ch1 <- "hello" //invalid operation: ch1 <- "hello" (send to receive-only type <-chan string)
}
