package main
import "fmt"
func main(){
	i := 1
	sum := 0
	for i <= 40 {
		if i%3 == 0 {
			sum += i
			fmt.Print(i)
			if i < 39 {
				fmt.Print("+")
			} else {
				fmt.Printf(" = %d \n", sum)
			}
		}
		i++
	}
}
