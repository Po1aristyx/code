package main
import (
	"fmt"
)
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		//time.Sleep(1 * time.Second)
		ch1 <- 100
	}()
	go func() {
		//time.Sleep(1 * time.Second)
		ch2 <- 200
	}()
	select {
	case data := <-ch1:
		fmt.Println("ch1中读取数据了:", data)
	case data := <-ch2:
		fmt.Println("ch2中读取数据了：", data)
	default:
		fmt.Println("执行了default。。。")
	}
}

