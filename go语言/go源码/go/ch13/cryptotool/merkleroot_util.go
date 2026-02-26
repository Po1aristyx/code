package cryptotool

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	GetMerkleRoot()

	////以0为例：有两笔交易。coinbase交易及真实交易
	//testString := "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b"
	//testString2 := "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b"
	//
	////将两笔交易进行小端排序
	//fmt.Println(ReverseStr2ByteStr(testString))
	//fmt.Println(ReverseStr2ByteStr(testString2))
	////3ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a
	////3ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a
	//
	////将交易拼接在一起
	//a := "3ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a3ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a"
	////字符串转字节数组
	//arr, _ := hex.DecodeString(a)
	//
	////进行两次hash
	//res := sha256.Sum256(arr)
	//res = sha256.Sum256(res[:])
	//
	////将字节数组转字符串
	//str := hex.EncodeToString(res[:])
	//fmt.Println(str)
	//
	////转成大端排序格式
	//fmt.Println(ReverseStr2ByteStr(str))

}

func GetMerkleRoot() {
	//以以98901为例：有两笔交易。coinbase交易及真实交易
	testString := "16f0eb42cb4d9c2374b2cb1de4008162c06fdd8f1c18357f0c849eb423672f5f"
	testString2 := "cce2f95fc282b3f2bc956f61d6924f73d658a1fdbc71027dd40b06c15822e061"

	//将两笔交易进行小端排序
	fmt.Println(ReverseHexString(testString))
	fmt.Println(ReverseHexString(testString2))
	//5f2f6723b49e840c7f35181c8fdd6fc0628100e41dcbb274239c4dcb42ebf016
	//61e02258c1060bd47d0271bcfda158d6734f92d6616f95bcf2b382c25ff9e2cc

	//将交易拼接在一起
	a := "5f2f6723b49e840c7f35181c8fdd6fc0628100e41dcbb274239c4dcb42ebf01661e02258c1060bd47d0271bcfda158d6734f92d6616f95bcf2b382c25ff9e2cc"
	//字符串转字节数组
	arr, _ := hex.DecodeString(a)

	//进行两次hash
	res := sha256.Sum256(arr)
	res = sha256.Sum256(res[:])

	//将字节数组转字符串
	str := hex.EncodeToString(res[:])
	fmt.Println(str)

	//转成大端排序格式
	fmt.Println(ReverseHexString(str))

}
