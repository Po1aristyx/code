package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	testReader()
}
func testReader() {
	fileName1 := "D:\\QQ数据\\go源码\\go\\ch09\\files\\blockchain.txt"
	file1, _ := os.Open(fileName1)
	reader1 := bufio.NewReader(file1)
	fmt.Printf("%T\n", reader1)

	for {
		s1, err := reader1.ReadString('\n')
		fmt.Print(s1)
		if err == io.EOF {
			fmt.Println("读取完毕！")
			break
		}
	}

	file1.Close()
}
