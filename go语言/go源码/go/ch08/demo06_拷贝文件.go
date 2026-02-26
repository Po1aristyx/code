package main

import (
	"os"
	"io"
	"fmt"
)

func main() {
	//srcFile := "./files/forbiddencity.docx"
	//destFile := "./files/xyz.docx"
	//total, err := copyFile1(srcFile, destFile)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println("复制ok：", total)
	//}

	srcFile := "./files/dsa.png"
	destFile := "./files/xyz.png"
	total, err := copyFile2(srcFile, destFile)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("复制ok：", total)
	}

}

func copyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	bs := make([]byte, 1024*8, 1024*8)
	n := -1
	total := 0
	for {
		n, err = file1.Read(bs)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		} else {
			total += n
			file2.Write(bs[:n])
		}
	}
	return total, nil
}

func copyFile2(srcFile, destFile string)(int64 , error) {
	file1 , err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2 , file1)
}
