package main
import (
	"encoding/base64"
	"fmt"
)
func main() {
	str := "心怀不惧，方能翱翔于天际"
	cipherText := Base64EncodeString(str)
	fmt.Println("base64 编码后：",cipherText)
	fmt.Println("base64 解码后：",Base64DecodeString(cipherText))
}
func Base64EncodeString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func Base64DecodeString(str string) string {
	result, _ := base64.StdEncoding.DecodeString(str)
	return string(result)
}

