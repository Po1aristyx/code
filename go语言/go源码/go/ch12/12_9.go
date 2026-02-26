package main
import (
	"fmt"
	"time"
)
func main() {
	//1.非缓冲通道
	ch1 := make(chan int)
	fmt.Println("非缓冲通道" ,len(ch1) , cap(ch1))
	go func(){
		data := <- ch1
		fmt.Println("获得数据", data)
	}()
	//
	ch1 <- 100
	time.Sleep(time.Second)
	fmt.Println("赋值ok" , "main over...")

	//2、非缓冲通道
	ch2 := make(chan string)
	go sendData(ch2)

	for data := range ch2 {
		fmt.Println("\t 读取数据", data)
	}
	fmt.Println("main over...")

	//3.缓冲通道,缓冲区满了才会阻塞
	ch3 := make(chan string , 6)
	go sendData(ch3)

	for data := range ch3 {
		fmt.Println("\t 读取数据", data)
	}
	fmt.Println("main over...")
}
func sendData(ch chan string) {
	for i := 1; i <= 3; i++ {
		ch <- fmt.Sprintf("data%d", i)
		fmt.Println("往通道放入数据：", i)
	}
	defer close(ch)
}
