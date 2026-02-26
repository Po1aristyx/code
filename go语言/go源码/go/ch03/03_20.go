package main
import "fmt"
func main(){
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break // 如果i > 5，则循环终止（loop is terminated ）
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
}
