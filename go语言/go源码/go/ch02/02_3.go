package main
import "fmt"
func main(){
	var flag bool
	fmt.Printf("%T , %t \n", flag, flag)
	flag = true
	fmt.Printf("%T , %t \n", flag, flag)
}
