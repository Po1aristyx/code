package main

import "fmt"

func main(){
	i := 0
	for ; ; i++ {
		if (i > 20) {
			break
		}
		fmt.Printf("%d ", i)
	}

}
