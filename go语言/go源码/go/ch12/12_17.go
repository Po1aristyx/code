package main
import (
	"fmt"
	"strings"
	"sync"
	"time"
	"strconv"
)
var tickets int = 20
var wg sync.WaitGroup
var mutex sync.Mutex
func main() {
	wg.Add(4)
	go saleTickets("1号窗口", &wg)
	go saleTickets("2号窗口", &wg)
	go saleTickets("3号窗口", &wg)
	go saleTickets("4号窗口", &wg)
	wg.Wait()
	defer fmt.Println("所有车票都售空！")
}
func saleTickets(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		//锁定
		mutex.Lock()
		if tickets > 0 {
			time.Sleep(1 * time.Second)
			//获取窗口的编号
			num, _ := strconv.Atoi(name[:1])
			pre := strings.Repeat("————————", num)
			fmt.Println(pre, name, tickets)
			tickets--
		} else {
			fmt.Printf("%s 结束售票 \n", name)
			mutex.Unlock()
			break
		}
		//解锁
		mutex.Unlock()
	}
}
