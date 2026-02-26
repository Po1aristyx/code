package main

import (
	"ch09/cryptotool"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	//test()
	msg := "865f0da3066c1e88a55b275daabebbaf6381526424829f52493942d6ada34198"
	strSigR := "761762e45789077777e11b816d0d811b53fdffa1e009f10fe187ee9600db53f9"
	strSigS := "116748c67d9af7f25e4f215b5091919ec65f956df37d2b7d5b4633bdf2d400fe"
	pubKey := "04ba7e4e7eca0f67fa34106bdd2eb332a03d11dc38e5f23afca959c459b279ff3dc9b07f87d7f76d7a2b3fc616b48d9c46e889b907d5900679aaf83e8d7077513f"
	res := cryptotool.VerifySignature([]byte(pubKey), []byte(msg), strSigR, strSigS)
	fmt.Println("签名验证结果：", res)
}

func test() {
	//1、生成签名
	fmt.Println("1、生成签名-------------------------------")
	//调用函数生成私钥与公钥
	privKey, pubKey := cryptotool.NewKeyPair()
	//信息的哈希
	msg := sha256.Sum256([]byte("hello world"))

	//根据私钥和信息的哈希进行数字签名，产生r和s
	r, s, _ := ecdsa.Sign(rand.Reader, &privKey, msg[:])

	//生成r、s字符串
	fmt.Println("-------------------------------")
	strSigR := fmt.Sprintf("%x", r)
	strSigS := fmt.Sprintf("%x", s)
	fmt.Println("r、s的10进制分别为：", r, s)
	fmt.Println("r、s的16进制分别为：", strSigR, strSigS)

	//r和s拼接在一起，形成数字签名的der格式
	signatureDer := cryptotool.MakeSignatureDerString(strSigR, strSigS)
	//打印数字签名的16进制显示
	fmt.Println("数字签名DER格式为：", signatureDer)
	fmt.Println()

	//2、签名验证过程
	fmt.Println("2、签名验证过程-------------------------------")
	res := cryptotool.VerifySig(pubKey, msg[:], r, s)
	fmt.Println("签名验证结果：", res)

	res = cryptotool.VerifySignature(pubKey, msg[:], strSigR, strSigS)
	fmt.Println("签名验证结果：", res)
}
