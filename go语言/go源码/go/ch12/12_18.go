package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwm sync.RWMutex

	// 创建两个通道
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 启动第二个goroutine，向ch2发送数据
	go func() {
		fmt.Println("goroutine 2，开始向ch2发送数据")
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	// 启动第一个goroutine，向ch1发送数据
	go func() {
		fmt.Println("goroutine 1，开始向ch1发送数据")
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	// 使用for循环替代select
	for {
		// 尝试从ch1和ch2接收数据
		select {
		case v1 := <-ch1:
			fmt.Printf("从ch1接收到数据: %d\n", v1)
			rwm.RLock()
			fmt.Println("对ch1数据进行读操作")
			rwm.RUnlock()
			// 如果接收到数据则退出循环
			break
		case v2 := <-ch2:
			fmt.Printf("从ch2接收到数据: %d\n", v2)
			rwm.RLock()
			fmt.Println("对ch2数据进行读操作")
			rwm.RUnlock()
			// 如果接收到数据则退出循环
			break
		default:
			// 如果两个通道都没有数据可接收，则继续循环
			time.Sleep(100 * time.Millisecond)
		}
		break
	}

	fmt.Println("主程序结束")
}
