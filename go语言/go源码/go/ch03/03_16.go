package main
import "fmt"
func main(){
	count := 0;
	for i := 32.0; i >= 4; i -= 1.5 {
		count++;
	}
	fmt.Println(count)
}
