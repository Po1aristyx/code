package main

import "fmt"

func main(){

	var value complex64 = 2.1 + 12i
	value2 :=complex(2.1,12)
	fmt.Println(real(value))
	fmt.Println(imag(value))
	fmt.Println(value2)
}