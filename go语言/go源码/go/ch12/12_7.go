package main
import (
	"fmt"
)
func main() {
	var ch1 chan int
	ch1 = make(chan int)
	fmt.Printf("%T\n", ch1)
	ch2 := make(chan bool)
	go func() {
		data, ok := <-ch1
		if ok {
			fmt.Println("子goroutine取到数值：", data)
		}
		ch2 <- true
	}()
	ch1 <- 10
	<-ch2 //目的就是为了 阻塞
	fmt.Println("main over...")
}