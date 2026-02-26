package main

import "fmt"

func main() {
	for i := 9; i >= 1; i-- { // i 控制行数，是乘法中的第二个数。
		for j := 9; j >= i; j-- { // j 控制每行的列数，是乘法中的第一个数。
			fmt.Printf("%d*%d=%-2d ", j, i, i*j)
		}
		fmt.Println()
	}
}
