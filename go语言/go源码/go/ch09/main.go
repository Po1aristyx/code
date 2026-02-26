package main

import (
	"ch09/cryptotool"
	"fmt"
)



func main() {
	str := "e2af004f2f8f75706dbeead105ce1dccd206ffb3f3e825a37659c92025fee7c6"
	x := cryptotool.HexStringToDecimal(str)
	fmt.Println(str)
	fmt.Println(x)
	//102531904150998900948772833063385126650172866314226367788803311503960837121990

	fmt.Println("=========")

}
