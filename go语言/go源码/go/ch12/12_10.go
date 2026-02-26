package main
import (
	"fmt"
	"time"
	"math/rand"
	"strings"
)
func main() {
	//用channel来传递数据, 不再需要自己去加锁维护一个全局的阻塞队列
	ch1 := make(chan int)
	ch_bool1 := make(chan bool) //判断结束
	ch_bool2 := make(chan bool) //判断结束
	ch_bool3 := make(chan bool) //判断结束
	rand.Seed(time.Now().UnixNano())
	//生产者
	go producer(ch1)
	//消费者
	go consumer(1 , ch1 , ch_bool1)
	go consumer(2 , ch1 , ch_bool2)
	go consumer(3 , ch1 , ch_bool3)
	<-ch_bool1
	<-ch_bool2
	<-ch_bool3

	defer fmt.Println("main...over...")
}
//生产者
func producer(ch1 chan int) {
	for i:=1; i<=10 ;i++  {
		ch1 <- i
		fmt.Println("生产蛋糕，编号为：" , i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
	defer close(ch1)
}
//消费者
func consumer(num int , ch1 chan int, ch chan bool) {
	for data := range ch1 {
		pre := strings.Repeat("————————" , num)
		fmt.Printf("%s %d号购买%d号蛋糕 \n" , pre , num , data)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
	ch <- true
	defer close(ch)
}
