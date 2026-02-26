package main
import (
	"fmt"
)
func main() {
	ch1 := make(chan string)
	go sendData(ch1)
	//1、循环接收数据方式1
	for {
		data := <-ch1
		//如果通道关闭，通道中传输的数据则为各数据类型的默认值。chan int 默认值为0，chan string默认值为"" 等。
		if data == "" {
			break
		}
		fmt.Println("从通道中读取数据方式1：", data)
	}
	//2、循环接收数据方式2
	//for {
	//	data, ok := <-ch1
	//	fmt.Println(ok)
	//	//通过多个返回值的形式来判断通道是否关闭，如果通道关闭，则ok值为false。
	//	if !ok {
	//		break
	//	}
	//	fmt.Println("从通道中读取数据方式2：", data)
	//}
	//3、循环接收数据方式3
	//for range循环会自动判断通道是否关闭，自动break循环。
	for value := range ch1 {
		fmt.Println("从通道中读取数据方式3：", value)
	}

}
func sendData(ch1 chan string) {
	defer close(ch1)
	for i := 0; i < 3; i++ {
		ch1 <- fmt.Sprintf("发送数据%d\n", i)
	}
	fmt.Println("发送数据完毕。。")
	//显式调用close()实现关闭通道
}