package main
import (
	"fmt"
	"regexp"
)
func main() {
	testRegexp()
}
func testRegexp() {
	//1.Match(pattern string, b []byte) (matched bool, err error)匹配检查文本正则表达式是否与字节片匹配。
	flag, _ := regexp.Match("^\\d{6,7}$", []byte("123456789"))
	fmt.Println(flag)
	//2.MatchString(pattern string, s string) (matched bool, err error)
	flag, _ = regexp.MatchString("^\\d{6,7}$", "0123456789")
	fmt.Println(flag)
	//3.Compile(expr string) (*Regexp, error)
	RegExp, _ := regexp.Compile("^\\d{6}\\D{2}$")
	//4.MustCompile(str string) *Regexp
	RegExp2 := regexp.MustCompile("^[\u4e00-\u9fa5]+$")
	//5.Match(b []byte) bool
	flag = RegExp2.Match([]byte("一丁丂"))
	fmt.Println("xxx:" , flag)
	//6.MatchString(s string) bool
	flag = RegExp.MatchString("123456ab")
	fmt.Println(flag)
	//7.ReplaceAll(src, repl []byte) []byte
	text := "将字符串 123 按照正则表达式 3 4 5 分割成子字符串 56 78 组成的切片"
	RegExp3 := regexp.MustCompile("[\\d\\s]+")
	result := string(RegExp3.ReplaceAll([]byte(text), []byte("-")))
	fmt.Println("替换后的字符串为：" , result)
	text = "第一部分#第二部分##第三部分###第四部分#第五部分##第六部分"
	MyRegexp := regexp.MustCompile("#+")
	arr := MyRegexp.Split(text, 5)
	fmt.Println(arr)
}
