package main
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)
func main() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("1234567890abcdefghijklmnopqrstuv")
	str := "区块链很有趣"
	fmt.Println("------------AES加密解密字节数组")
	resultArr, _ := AesEncrypt([]byte(str), key)
	fmt.Printf("加密后：%x\n", resultArr)
	resultArr, _ = AesDecrypt(resultArr, key)
	fmt.Println("解密后：", string(resultArr))
	fmt.Println("------------AES加密解密字符串")
	cipherText, _ := AesEncryptString(str, key)
	fmt.Println("加密后：", cipherText)
	originalText, _ := AesDecryptString(cipherText, key)
	fmt.Println("解密后：", originalText)
}
//AES加密字节数组，返回字节数组
func AesEncrypt(originalBytes, key []byte) ([]byte, error) {
   block, err := aes.NewCipher(key)
   if err != nil {
      return nil, err
   }
   blockSize := block.BlockSize()
   originalBytes = PKCS5Padding(originalBytes, blockSize)
   blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
   cipherBytes := make([]byte, len(originalBytes))
   blockMode.CryptBlocks(cipherBytes, originalBytes)
   return cipherBytes, nil
}
//AES解密字节数组，返回字节数组
func AesDecrypt(cipherBytes, key []byte) ([]byte, error) {
   block, err := aes.NewCipher(key)
   if err != nil {
      return nil, err
   }
   blockSize := block.BlockSize()
   blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
   originalBytes := make([]byte, len(cipherBytes))
   blockMode.CryptBlocks(originalBytes, cipherBytes)
   originalBytes = PKCS5UnPadding(originalBytes)
   return originalBytes, nil
}
//AES加密文本，返回对加密后字节数组进行base64处理后字符串
func AesEncryptString(originalText string, key []byte) (string, error) {
   cipherBytes, err := AesEncrypt([]byte(originalText), key)
   if err != nil {
      return "", err
   }
   base64str := base64.StdEncoding.EncodeToString(cipherBytes)
   return base64str, nil
}
//对Base64处理后的加密文本进行DES解密，返回解密后明文
func AesDecryptString(cipherText string, key []byte) (string, error) {
   cipherBytes, _ := base64.StdEncoding.DecodeString(cipherText)
   cipherBytes, err := AesDecrypt(cipherBytes, key)
   if err != nil {
      return "", err
   }
   return string(cipherBytes), nil
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
