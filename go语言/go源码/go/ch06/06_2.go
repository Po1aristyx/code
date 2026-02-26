package main
import (
	"fmt"
	"strings"
	"unicode"
)
func main() {
	TestContains()
	TestCount()
	TestIndex()
	TestIndexFunc()
	TestLastIndex()
	TestLastIndexFunc()
	res := GetFileSuffix("abc.xyz.lmn.jpg")
	fmt.Println(res)
}
func TestContains() {
	fmt.Println(strings.Contains("seafood", "foo"))   //true
	fmt.Println(strings.Contains("seafood", "bar"))   //false
	fmt.Println(strings.Contains("seafood", ""))      //true
	fmt.Println(strings.Contains("", ""))             //true
	fmt.Println(strings.Contains("steven王2008", "王")) //true
}
func TestContainsAny() {
	fmt.Println(strings.ContainsAny("team", "i"))        //false
	fmt.Println(strings.ContainsAny("failure", "u & i")) //true
	fmt.Println(strings.ContainsAny("foo", ""))          //false
	fmt.Println(strings.ContainsAny("", ""))             //false
}
func TestContainsRune() {
	fmt.Println(strings.ContainsRune("一丁丂", '丁'))   //true
	fmt.Println(strings.ContainsRune("一丁丂", 19969)) //true
}
func TestCount() {
	fmt.Println(strings.Count("cheese", "e")) //3
	fmt.Println(strings.Count("one", ""))     //4
}
func TestHasPrefix() {
	fmt.Println(strings.HasPrefix("1000phone news", "1000"))  //true
	fmt.Println(strings.HasPrefix("1000phone news", "1000a")) //false
}
func TestHasSuffix() {
	fmt.Println(strings.HasSuffix("1000phone news", "news")) //true
	fmt.Println(strings.HasSuffix("1000phone news", "new"))  //false
}
func TestIndex() {
	fmt.Println(strings.Index("chicken", "ken")) //4
	fmt.Println(strings.Index("chicken", "dmr")) //-1
}
func TestIndexAny() {
	fmt.Println(strings.IndexAny("abcABC120", "教育基地A")) //3
}
func TestIndexByte() {
	fmt.Println(strings.IndexByte("123abc", 'a')) //3
}
func TestIndexRune() {
	fmt.Println(strings.IndexRune("abcABC120", 'C')) //5
	fmt.Println(strings.IndexRune("It培训教育", '教'))    //8
}
func TestIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello123,中国", f)) //9
}
func TestLastIndex() {
	fmt.Println(strings.LastIndex("Steven learn english", "e")) //13
	fmt.Println(strings.Index("go gopher", "go"))               //0
	fmt.Println(strings.LastIndex("go gopher", "go"))           //3
	fmt.Println(strings.LastIndex("go gopher", "rodent"))       //-1
}
func TestLastIndexAny() {
	fmt.Println(strings.LastIndexAny("chicken", "aeiouy")) //5
	fmt.Println(strings.LastIndexAny("crwth", "aeiouy"))   //-1
}
func TestLastIndexByte() {
	fmt.Println(strings.LastIndexByte("abcABCA123", 'A')) //6
}
func TestLastIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.LastIndexFunc("Hello,世界", f))       //9
	fmt.Println(strings.LastIndexFunc("Hello,world中国人", f)) //17
}
func GetFileSuffix(str string) string {
	arr := strings.Split(str, ".")
	return arr[len(arr)-1]
}
