package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	//fileName1 := "./files/blockchain.txt"
	//file1, _ := os.Open(fileName1)
	reader1 := bufio.NewReader(strings.NewReader("abcdefg 1000phone blockchain ready go"),)
	scanner := bufio.NewScanner(reader1)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "q!" {
			break
		}
	}
}