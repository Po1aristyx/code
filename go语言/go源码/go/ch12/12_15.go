package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(10 * time.Millisecond)
		data := <-ch1
		fmt.Println("ch1：", data)
	}()
	go func() {
		time.Sleep(2 * time.Second)
		data := <-ch2
		fmt.Println("ch2：", data)
	}()
	select {
	case ch1 <- 100: //阻塞
		close(ch1)
		fmt.Println("ch1中写入数据。。")
	case ch2 <- 200: //阻塞
		close(ch2)
		fmt.Println("ch2中写入数据。。")
	case <-time.After(2 * time.Millisecond): //阻塞
		fmt.Println("执行延时通道")
		//default:
		//	fmt.Println("default..")
	}

	time.Sleep(4 * time.Second)
	fmt.Printf("main  over ")
}
