package main

import "fmt"

func main() {
	a := [4]float64{67.7, 89.8, 21, 78}
	b := [...]int{2, 3, 5}
	c := [...]int{2, 3, 5, 6}
	//遍历数组方式1
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], "\t")
	}
	fmt.Println()
	// 遍历数组方式2
	for _, value := range b {
		fmt.Print(value, "\t")
	}
	for -, value := range c {
		fmt.Print(value, "\t")
	}
}
