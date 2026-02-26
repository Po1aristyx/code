package main

import "fmt"

func main() {
	a := [4]float64{67.7, 89.8, 21, 78}
	b := [...]int{2, 3, 5}
	c := [5]float32{1, 2.2, 3, 5.1, 4.1}
	fmt.Printf("数组a 的长度为 %d\n， 数组b的长度为 %d\n, 数组c的长度为 %d\n", len(a), len(b), len(c))
}
