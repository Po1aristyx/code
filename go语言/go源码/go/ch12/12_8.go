package main
import "fmt"
func main() {
	//channel关闭后是否可以写入和读取呢？
	ch1 := make(chan int)
	go func() {
		ch1 <- 100
		ch1 <- 200
		close(ch1)
		ch1 <- 10 //关闭的channel，无法写入数据
	}()
	data, ok := <-ch1
	fmt.Println("main读取数据：", data, ok)
	data, ok = <-ch1
	fmt.Println("main读取数据：", data, ok)
	data, ok = <-ch1
	fmt.Println("main读取数据：", data, ok)
	data, ok = <-ch1
	fmt.Println("main读取数据：", data, ok)
	data, ok = <-ch1
	fmt.Println("main读取数据：", data, ok)
}

