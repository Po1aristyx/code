package main
import (
	"fmt"
	"strings"
	"unicode"
)
func main() {
	TestFields()
	TestFieldsFunc()
	TestSplitAfterN()
	TestSplit()
}
func TestFields() {
	fmt.Println(strings.Fields("  abc 123 ABC xyz XYZ")) //[abc 123 ABC xyz XYZ]
}
func TestFieldsFunc() {
	f := func(c rune) bool {
		//return c == '='
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.FieldsFunc("abc@123*ABC&xyz%XYZ" , f)) //[abc 123 ABC xyz XYZ]
}
func TestSplit() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))//[a b c]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))//["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))//[" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))//[""]
}
func TestSplitN() {
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))//["a"   "b,c"]
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 1))//["a,b,c"]
}
func TestSplitAfter() {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))//["a," "b," "c"]
}
func TestSplitAfterN() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))//["a,"   "b,c"]
}