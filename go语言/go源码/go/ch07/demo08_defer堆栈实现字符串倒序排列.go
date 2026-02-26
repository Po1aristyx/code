package main

import (
	"fmt"
	"time"
)

func main() {
	str := "Steven欢迎大家学习区块链，123开始吧"

	fmt.Println("defer实现")
	ReverseString(str)

	fmt.Println("-----------------")

	fmt.Println("普通for循环实现")
	ReverseStr(str)

}

//利用for循环实现字符串倒序排列
func ReverseStr(str string) {
	time1 := time.Now()
	temp := []rune{}
	sli := []rune(str)
	for i := len(sli) - 1; i >= 0; i-- {
		temp = append(temp , sli[i])
	}
	fmt.Println(string(temp))
	time2 := time.Now()
	fmt.Println("普通for循环实现：" , time2.Sub(time1).Nanoseconds())
}

//利用defer堆栈执行字符串倒序排列
func ReverseString(str string) {
	time1 := time.Now()
	for _, v := range []rune(str) {
		defer fmt.Printf("%c" , v)
	}
	time2 := time.Now()
	fmt.Println("defer实现：" , time2.Sub(time1).Nanoseconds())
}
