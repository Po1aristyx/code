package main
import "fmt"
func main() {
	s1 := []int{78, 100, 2, 400, 324}
	getLargest(s1)
}
func finished() {
	fmt.Println("结束！")
}
func getLargest(s []int) {
	defer finished()
	fmt.Println("开始寻找最大数值：")
	max := s[0]
	for _ , v := range s {
		if v > max {
			max = v
		}
	}
	fmt.Printf("%v中最大数为：%v \n" , s , max)
}
