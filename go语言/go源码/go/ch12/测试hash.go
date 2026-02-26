package main

import (
	"ch09/cryptotool"
	"fmt"
)

func main() {
	testString := "1"
	fmt.Println(cryptotool.HASH(testString , "md4" , false))
	fmt.Println(cryptotool.HASH(testString , "md5" , false))
	fmt.Println(cryptotool.HASH(testString , "sha1" , false))
	fmt.Println(cryptotool.HASH(testString , "ripemd160" , false))
	fmt.Println(cryptotool.HASH(testString , "sha256" , false))
	fmt.Println(cryptotool.HASH(testString , "sha512" , false))

}
