package main

import (
	"os"
	"bufio"
	"fmt"
	"io"
)

func main() {
	/*
		bufio.NewReader(file1) //构建带缓存的Reader对象：bufio.Reader
		reader.ReadString()
		reader.ReadBytes()

		bufio.NewWriter()
		writer.Write()
		writer.WriteString()
		writer.Flush()

		bufio.NewScanner(reader1)
		scanner.Split()
		scanner.Text()

	*/

	//testReader()

	testWriter()
}

//1、构建带缓存的Reader对象：bufio.Reader
func testReader() {
	fileName1 := "./files/blockchain.txt"
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

//
func testWriter() {
	fileName2 := "./files/yesterday.mp3"
	file2, _ := os.Open(fileName2)
	reader2 := bufio.NewReader(file2)

	fileName3 := "./files/abc.mp3"
	file3, _ := os.OpenFile(fileName3, os.O_WRONLY|os.O_CREATE, os.ModePerm)

	writer1 := bufio.NewWriter(file3)

	for {
		bs, err := reader2.ReadBytes(' ')
		writer1.Write(bs)
		writer1.Flush()
		if err == io.EOF {
			fmt.Println("文件读取完毕！")
			break
		}
	}
	file2.Close()
	file3.Close()
}
