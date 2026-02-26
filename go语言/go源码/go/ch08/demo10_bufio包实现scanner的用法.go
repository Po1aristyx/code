package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	//s := ""
	//fmt.Scanln(&s)
	//fmt.Println(s)

	reader1 := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader1)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "q!" {
			break
		}
	}
}
