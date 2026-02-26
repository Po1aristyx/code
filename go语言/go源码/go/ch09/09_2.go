package main
import (
	"path/filepath"
	"fmt"
	"path"
)
func main() {
	filePath1 := "D:/go/ch09/picture/soldier.jpg"
	filePath2 := "./picture/timg.jpg"
	fmt.Println(filepath.IsAbs(filePath1)) //true
	fmt.Println(filepath.IsAbs(filePath2)) //false
	fmt.Println(filepath.Rel("D:/go/ch09", filePath1))
	fmt.Println(filepath.Abs(filePath1))
	fmt.Println(filepath.Abs(filePath2))
	fmt.Println(path.Join(filePath1, ".."))
	fmt.Println(path.Join(filePath1, "."))
	fmt.Println(path.Join("D:/blockChain", filePath2))
}
