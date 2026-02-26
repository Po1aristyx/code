package main

import "fmt"

func main(){
	fmt.Printf("%b \n", 123.123456)

	fmt.Printf("%f \n", 123.1)
	fmt.Printf("%.2f \n", 123.125456)

	fmt.Printf("%e \n", 123.123456)
	fmt.Printf("%E \n", 123.123456)
	fmt.Printf("%.1e \n", 123.125456)
	fmt.Printf("%F \n", 123.123456)
	fmt.Printf("%g \n", 123.123456)
	fmt.Printf("%G \n", 123.123456)
}
