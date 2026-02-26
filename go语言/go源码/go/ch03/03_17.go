package main

import "fmt"

func main() {
	// 定义行数
	lines := 8
	for i := lines; i >= 0; i-- {
		for n := 0; n < 2*i+1; n++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
