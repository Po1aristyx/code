package main

import (
	"fmt"
)

func main() {
	TestDes()
	fmt.Println()
	Test3Des()
}

func TestDes() {
	key := []byte("00000000") //只占8个字节
	fmt.Println("------------DES加密解密字节数组")
	resultArr, err := cryptotool.DesEncrypt([]byte("一"), key)
	//resultArr, err := cryptotool.DesEncrypt([]byte("Steven陪你学区块链开发技术"), key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后：%v\n", resultArr)
	fmt.Printf("加密后：%x\n", resultArr)

	resultArr, err = cryptotool.DesDecrypt(resultArr, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后：", string(resultArr))

	fmt.Println("------------DES加密解密字符串")
	cipherText, err := cryptotool.DesEncryptString("Steven陪你学区块链", key)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后：" , cipherText)

	originalText, err := cryptotool.DesDecryptString(cipherText, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后：", originalText)
}

func Test3Des() {
	//key := []byte("1234abcd1234abcd1234abcd")
	key := []byte("abcdefghijklmnopqrstuvwx") //占24个字节
	fmt.Println("------------3DES加密解密字节数组")
	result, err := cryptotool.TripleDesEncrypt([]byte("a"), key)
	//result, err := cryptotool.TripleDesEncrypt([]byte("Steven陪你学区块链"), key)
	fmt.Printf("加密后：%v\n", result)
	fmt.Printf("加密后：%x\n", result)
	if err != nil {
		panic(err)
	}
	origData, err := cryptotool.TripleDesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后：", string(origData))

	fmt.Println("------------3DES加密解密字符串")
	cipherText, err := cryptotool.TripleDesEncrypt2Str("Steven陪你学区块链", key)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后：", cipherText)

	originalText, err := cryptotool.TripleDesDecrypt2Str(cipherText, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后：", originalText)
}
