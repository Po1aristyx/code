package main
import (
	"fmt"
	"time"
)
func main() {
	//2.使用After(),返回值<- chan Time,同Timer.C
	ch1 := time.After(5 * time.Second)
	fmt.Println(time.Now())
	data := <-ch1
	fmt.Printf("%T\n",data) //time.Time
	fmt.Println(data)
}



