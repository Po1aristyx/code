package main

import (
	"fmt"
	"time"
)

func main() {
	//创建计时器
	timer1 := time.NewTimer(5 * time.Second)
	timer2 := time.NewTimer(10 * time.Second)
	//fmt.Printf("%T\n", timer1) //*time.Timer
	fmt.Println(time.Now())
	data := <-timer1.C //<-chan time.Time
	//timer1.C <- 2018-11-07 21:32:15.135125 +0800 CST m=+0.000275697
	fmt.Printf("%T\n", timer1.C) //<-chan time.Time
	fmt.Printf("%T\n", data)     //time.Time
	fmt.Println(data)

	data1 := <-timer2.C          //<-chan time.Time
	fmt.Printf("%T\n", timer2.C) //<-chan time.Time
	fmt.Printf("%T\n", data)     //time.Time
	fmt.Println(data1)
}
