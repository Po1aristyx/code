package main
import (
	"os"
	"bufio"
	"fmt"
	"io"
)
func main() {
	testWriter()
}
func testWriter() {
	fileName2 := "./files/onlyMyRailgun.mp3"
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
