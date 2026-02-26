package main

import (
	"ch09/cryptotool"
	"fmt"
)

func main() {
	data, err := cryptotool.RsaEncryptString("Steven陪你学区块链+Go语言！！！")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println("------------------------")
	origData, err := cryptotool.RsaDecryptString(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}
