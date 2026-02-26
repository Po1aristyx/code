package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	//time.Sleep(50 * time.Microsecond)
	fmt.Println("第五人格启动")
}
