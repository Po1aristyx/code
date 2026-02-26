package main

import "fmt"

func main() {
	var i int
	for i = 10; ; {
		fmt.Print(i)
		i++
		if i > 20 {
			break
		}
	}

}
