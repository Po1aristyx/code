package main
import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
)
func main(){
	str := "1000phone"
	fmt.Println(str)
	str1 := HASH(str, "md5", false)
	fmt.Println(str1)
	str2 := HASH(str, "sha1", false)
	fmt.Println(str2)
	str3 := HASH(str, "sha256", false)
	fmt.Println(str3)
	arr := SHA256Double(str, false)
	fmt.Println(string(arr))
	str4 := SHA256DoubleString(str, false)
	fmt.Println(str4)
}
func HASH(text string, hashType string, isHex bool) string {
	var hashInstance hash.Hash
	switch hashType {
	case "md5":
		hashInstance = md5.New()
	case "sha1":
		hashInstance = sha1.New()
	case "sha256":
		hashInstance = sha256.New()
	case "sha512":
		hashInstance = sha512.New()
	}
	if isHex {
		arr , _ := hex.DecodeString(text)
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}
	cipherBytes := hashInstance.Sum(nil)
	return fmt.Sprintf("%x" , cipherBytes)
}

func SHA256Double(text string, isHex bool) []byte {
	hashInstance := sha256.New()
	if isHex {
		arr , _ := hex.DecodeString(text)
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}
	cipherBytes := hashInstance.Sum(nil)
	hashInstance.Reset()
	hashInstance.Write(cipherBytes)
	cipherBytes = hashInstance.Sum(nil)
	return cipherBytes
}


func SHA256DoubleString(text string, isHex bool) string {
	hashInstance := sha256.New()
	if isHex {
		arr , _ := hex.DecodeString(text)
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}
	cipherBytes := hashInstance.Sum(nil)
	hashInstance.Reset()
	hashInstance.Write(cipherBytes)
	cipherBytes = hashInstance.Sum(nil)
	return fmt.Sprintf("%x" , cipherBytes)
}

