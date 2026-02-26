package main
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)
func main() {
	str := "一篇诗，一斗酒，一曲长歌，一剑天涯"
	fmt.Println("加密前：", str)
	data, _ := RsaEncryptString(str)
	fmt.Println("加密后", data)
	origData, _ := RsaDecryptString(data)
	fmt.Println("解密后：", string(origData))
}
var decrypted string
var privateKey, publicKey []byte
func init() {
   var err error
   flag.StringVar(&decrypted, "d", "", "加密过的数据")
   flag.Parse()
   publicKey, err = ioutil.ReadFile("public.pem")
   if err != nil {
      os.Exit(-1)
   }
   privateKey, err = ioutil.ReadFile("private.pem")
   if err != nil {
      os.Exit(-1)
   }
}
// 加密字节数组，返回字节数组
func RsaEncrypt(origData []byte) ([]byte, error) {
   block, _ := pem.Decode(publicKey)
   if block == nil {
      return nil, errors.New("public key error")
   }
   pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
   if err != nil {
      return nil, err
   }
   pub := pubInterface.(*rsa.PublicKey)
   return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}
// 解密字节数组，返回字节数组
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
   block, _ := pem.Decode(privateKey)
   if block == nil {
      return nil, errors.New("private key error!")
   }
   priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      return nil, err
   }
   return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
// 加密字符串，返回base64处理的字符串
func RsaEncryptString(origData string) (string, error) {
   block, _ := pem.Decode(publicKey)
   if block == nil {
      return "", errors.New("public key error")
   }
   pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
   if err != nil {
      return "", err
   }
   pub := pubInterface.(*rsa.PublicKey)
   cipherArr, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData))
   if err != nil {
      return "", err
   } else {
      return base64.StdEncoding.EncodeToString(cipherArr), nil
   }
}
// 解密经过base64处理的加密字符串，返回加密前的明文
func RsaDecryptString(cipherText string) (string, error) {
   block, _ := pem.Decode(privateKey)
   if block == nil {
      return "", errors.New("private key error!")
   }
   priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
   if err != nil {
      return "", err
   }
   cipherArr, _ := base64.StdEncoding.DecodeString(cipherText)
   originalArr, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cipherArr)
   if err != nil {
      return "", err
   } else {
      return string(originalArr), nil
   }
}

