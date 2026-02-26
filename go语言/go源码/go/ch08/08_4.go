package main
import "fmt"
func main() {
	defer funcA()
	funcB()
	defer funcC()
	fmt.Println("main over...")
}
func funcA() {
	fmt.Println("这是funcA")
}
func funcB() {
	fmt.Println("这是funcB")
}
func funcC() {
	fmt.Println("这是funcC")
}