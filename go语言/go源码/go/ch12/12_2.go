package main
import (
	"fmt"
	"time"
)
func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	time.Sleep(50 * time.Microsecond)
	fmt.Println("main function")
}
