package main
import (
	"fmt"
	"strings"
)
func main() {
	TestCompare()
	TestEqualFold()
	TestRepeat()
	TestReplace()
	TestJoin()
}
func TestCompare() {
	fmt.Println(strings.Compare("abc", "bcd"))//-1
	fmt.Println("abc" < "bcd")
}
func TestEqualFold() {
	fmt.Println(strings.EqualFold("Go", "go"))//true
}
func TestRepeat() {
	fmt.Println("g" + strings.Repeat("o", 8) + "le")//goooooooole
}
func TestReplace() {
	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "张", 2))//张老大 张老二 王老三
	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "张", -1))//张老大 张老二 张老三
}
func TestJoin() {
	s := []string{"abc", "ABC", "123"}
	fmt.Println(strings.Join(s, ","))//abc,ABC,123
	fmt.Println(strings.Join(s, ""))//abcABC123
}

