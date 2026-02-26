package main

import "fmt"

func main() {
	/* 定义局部变量 */
	fmt.Print("1-50的素数：")
	var a, b, c, d int
	for a = 2; a <= 50; a++ {
		for b = 2; b <= (a / b); b++ {
			if a%b == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if b > (a / b) {
			fmt.Printf("%d\t", a)
			c = c + a
			d = d + 1
		}
		fmt.Printf("\n")
		fmt.Printf("%d\n %d\n", c, d)
	}
}
