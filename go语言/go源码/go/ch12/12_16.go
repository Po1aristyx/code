package main
import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)
func main() {
	var wg sync.WaitGroup
	fmt.Printf("%T\n", wg) //sync.WaitGroup
	fmt.Println(wg)        //{{} [0 0 0 0 0 0 0 0 0 0 0 0] 0}
	wg.Add(3)
	rand.Seed(time.Now().UnixNano())
	go printNum(&wg, 1)
	go printNum(&wg, 2)
	go printNum(&wg, 3)

	wg.Wait() //main,进入阻塞状态，当计数器为0后解除阻塞
	defer fmt.Println("main over...")
}
func printNum(wg *sync.WaitGroup, num int) {
	for i := 1; i <= 3; i++ {
		pre := strings.Repeat("\t", num-1)
		fmt.Printf("%s 第%d号子goroutine , %d \n", pre, num, i)
		//time.Sleep(time.Duration(rand.Intn(1000)))
		time.Sleep(time.Second)
	}
	wg.Done() //计数器减1
}
