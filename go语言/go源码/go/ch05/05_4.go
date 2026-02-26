package main

import "fmt"

func main() {
	a := [...]string{"USA", "China", "India", "Germany", "France", "America"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a ：", a)
	fmt.Println("b ： ", b)
}
