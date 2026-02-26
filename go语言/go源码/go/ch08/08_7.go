package main
import "fmt"
func main() {
	a := 5
	b := 6
	defer printAdd(a, b, true)
	a = 10
	b = 7
	printAdd(a, b, false)
}
func printAdd(a, b int, flag bool) {
	if flag {
		fmt.Printf("延迟执行函数printAdd() ，参数a，b分别为%d, %d , 两数之和为：%d\n", a, b, a+b)
	} else {
		fmt.Printf("未延迟执行函数printAdd() ，参数a，b分别为%d, %d , 两数之和为：%d\n", a, b, a+b)
	}
}