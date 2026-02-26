package main
import "fmt"
func main(){
	fmt.Printf("%T , %d \n", 123, 123)
	fmt.Printf("%T , %5d \n", 123, 123)
	fmt.Printf("%T , %05d \n", 123, 123)
	fmt.Printf("%T , %b \n", 123, 123)
	fmt.Printf("%T , %o \n", 123, 123)
	fmt.Printf("%T , %c \n", 97, 97)
	fmt.Printf("%T , %q \n", 97, 97)

	fmt.Printf("%T ,%x \n", 123, 123)
	fmt.Printf("%T ,%X \n", 123,123)
	fmt.Printf("%T ,%U \n",'一' ,'一')
}