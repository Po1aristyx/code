package main

import (
	"fmt"
	"strings"
)

func main() {
	result := StringToLower("AbcdefGHijklMNOPqrstUVWxyz", processCase)
	fmt.Println(result)
	result = StringToLower2("AbcdefGHijklMNOPqrstUVWxyz", processCase)
	fmt.Println(result)

	result = StringToUpper("AbcdefGHijklMNOPqrstUVWxyz", processCase2)
	fmt.Println(result)
}

// 处理字符串，奇数偶数依次显示为大小写
func processCase(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}
func StringToLower(str string, f func(string) string) string {
	fmt.Printf("%T \n", f)
	return f(str)
}

type caseFunc func(string) string // 声明了一个函数类型。通过type关键字，caseFunc会形成一种新的类型。
func StringToLower2(str string, f caseFunc) string {
	fmt.Printf("%T \n", f)
	return f(str)
}

func processCase2(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToLower(string(value))
		} else {
			result += strings.ToUpper(string(value))
		}
	}
	return result
}
func StringToUpper(str string, f func(string) string) string {
	fmt.Printf("%T \n", f)
	return f(str)
}
