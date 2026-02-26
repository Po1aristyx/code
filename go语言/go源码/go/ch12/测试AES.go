package main

import (
	"ch09/cryptotool"
	"fmt"
)

func main() {
	TestAes()
}

func TestAes() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("1234567890abcdefghijklmnopqrstuv")
	fmt.Println("------------AES加密解密字节数组")
	resultArr, err := cryptotool.AesEncrypt([]byte("Steven陪你学区块链"), key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后：%x\n", resultArr)

	resultArr, err = cryptotool.AesDecrypt(resultArr, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后：", string(resultArr))

	fmt.Println("------------AES加密解密字符串")
	cipherText, err := cryptotool.AesEncryptString("Steven陪你学区块链", key)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后：", cipherText)

	originalText, err := cryptotool.AesDecryptString(cipherText, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后：", originalText)
}
