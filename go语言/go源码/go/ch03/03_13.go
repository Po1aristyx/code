package main

import "fmt"

func main(){
	str := "123ABCabc一丁丂"
	for i, value := range str {
		fmt.Printf("第 %d 位的ASCII值=%d ，字符是%c \n", i, value ,value)
	}
}
