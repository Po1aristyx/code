package cryptotool

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"

	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
)

func SHA256HexString(text string) string {
	sha256Instance := sha256.New()
	arr, _ := hex.DecodeString(text)
	sha256Instance.Write(arr)
	ciphertext := sha256Instance.Sum(nil)
	return fmt.Sprintf("%x", ciphertext)
}

func SHA256Double(text string, isHex bool) string {
	sha256Instance := sha256.New()
	if isHex {
		arr, _ := hex.DecodeString(text)
		sha256Instance.Write(arr)
	} else {
		sha256Instance.Write([]byte(text))
	}
	hashBytes := sha256Instance.Sum(nil)
	sha256Instance.Reset()
	sha256Instance.Write(hashBytes)
	hashBytes = sha256Instance.Sum(nil)
	return fmt.Sprintf("%x", hashBytes)
}

func HASH(text string, hashType string, isHex bool) string {
	var hashInstance hash.Hash
	switch hashType {
	case "md4":
		hashInstance = md4.New()
	case "md5":
		hashInstance = md5.New()
	case "sha1":
		hashInstance = sha1.New()
	case "sha256":
		hashInstance = sha256.New()
	case "sha512":
		hashInstance = sha512.New()
	case "ripemd160":
		hashInstance = ripemd160.New()
	}
	if isHex {
		arr, _ := hex.DecodeString(text)
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}
	ciphertext := hashInstance.Sum(nil)
	return fmt.Sprintf("%x", ciphertext)
}

// 获取校验码
func GetCheckcode(hexString string) string {
	sha256DoubleString := SHA256Double(hexString, true)
	// 获取校验码
	rs := []byte(sha256DoubleString)
	return string(rs[:8])
}
