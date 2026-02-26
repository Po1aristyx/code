package cryptotool

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

//DES加密字节数组，返回字节数组
func DesEncrypt(originalBytes, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	originalBytes = PKCS5Padding(originalBytes, block.BlockSize())
	//fmt.Println(block.BlockSize())
	//originalBytes = ZeroPadding(originalBytes, block.BlockSize())
	fmt.Println(originalBytes)
	blockMode := cipher.NewCBCEncrypter(block, key)
	cipherBytes := make([]byte, len(originalBytes))
	blockMode.CryptBlocks(cipherBytes, originalBytes)
	return cipherBytes, nil
}

//DES解密字节数组，返回字节数组
func DesDecrypt(cipherBytes, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	originalBytes := make([]byte, len(cipherBytes))
	blockMode.CryptBlocks(originalBytes, cipherBytes)
	originalBytes = PKCS5UnPadding(originalBytes)
	//originalBytes = ZeroUnPadding(originalBytes)
	return originalBytes, nil
}

//DES加密文本，返回对加密后字节数组进行base64处理后字符串
func DesEncryptString(originalText string, key []byte) (string, error) {
	cipherBytes, err := DesEncrypt([]byte(originalText), key)
	if err != nil {
		return "", err
	}
	base64str := base64.StdEncoding.EncodeToString(cipherBytes)
	return base64str, nil
}

//对Base64处理后的加密文本进行DES解密，返回解密后明文
func DesDecryptString(cipherText string, key []byte) (string, error) {
	cipherBytes, _ := base64.StdEncoding.DecodeString(cipherText)
	cipherBytes, err := DesDecrypt(cipherBytes, key)
	if err != nil {
		return "", err
	}
	return string(cipherBytes), nil
}

//----------------------------------
// 3DES加密字节数组，返回字节数组
func TripleDesEncrypt(originalBytes, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	//fmt.Println("------" , block.BlockSize())
	if err != nil {
		return nil, err
	}
	originalBytes = PKCS5Padding(originalBytes, block.BlockSize())
	fmt.Println(originalBytes)
	// originalBytes = ZeroPadding(originalBytes, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	cipherBytes := make([]byte, len(originalBytes))
	blockMode.CryptBlocks(cipherBytes, originalBytes)
	return cipherBytes, nil
}

// 3DES解密字节数组，返回字节数组
func TripleDesDecrypt(cipherBytes, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	originalBytes := make([]byte, len(cipherBytes))
	blockMode.CryptBlocks(originalBytes, cipherBytes)
	originalBytes = PKCS5UnPadding(originalBytes)
	// origData = ZeroUnPadding(origData)
	return originalBytes, nil
}

// 3DES加密字符串，返回base64处理后字符串
func TripleDesEncrypt2Str(originalText string, key []byte) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	originalData := PKCS5Padding([]byte(originalText), block.BlockSize())
	// originalData = ZeroPadding(originalData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	cipherBytes := make([]byte, len(originalData))
	blockMode.CryptBlocks(cipherBytes, originalData)
	cipherText := base64.StdEncoding.EncodeToString(cipherBytes)
	return cipherText, nil
}

// 3DES解密base64处理后的加密字符串，返回明文字符串
func TripleDesDecrypt2Str(cipherText string, key []byte) (string, error) {
	cipherArr, _ := base64.StdEncoding.DecodeString(cipherText)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	originalBytes := make([]byte, len(cipherArr))
	blockMode.CryptBlocks(originalBytes, cipherArr)
	originalBytes = PKCS5UnPadding(originalBytes)
	// origData = ZeroUnPadding(origData)
	return string(originalBytes), nil
}

//func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
//	padding := blockSize - len(ciphertext)%blockSize
//	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//	return append(ciphertext, padtext...)
//}
//
//func PKCS5UnPadding(origData []byte) []byte {
//	length := len(origData)
//	// 去掉最后一个字节 unpadding 次
//	unpadding := int(origData[length-1])
//	return origData[:(length - unpadding)]
//}

//func ZeroPadding(ciphertext []byte, blockSize int) []byte {
//	padding := blockSize - len(ciphertext)%blockSize
//	padtext := bytes.Repeat([]byte{0}, padding)
//	return append(ciphertext, padtext...)
//}
//
//func ZeroUnPadding(origData []byte) []byte {
//	return bytes.TrimRightFunc(origData, func(r rune) bool {
//		return r == rune(0)
//	})
//}

//func TestDes() {
//	key := []byte("12345678")
//	resultArr, err := DesEncrypt([]byte("steven@1000phone.com"), key)
//	if err != nil {
//		panic(err)
//	}
//	//base64str := base64.StdEncoding.EncodeToString(resultArr)
//	//resultArr, _ = base64.StdEncoding.DecodeString(base64str)
//	//fmt.Println(resultArr)
//
//	resultArr, err = DesDecrypt(resultArr, key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(resultArr))
//
//	fmt.Println("------------\nDES解密字符串")
//
//	cipherText, err := DesEncryptString("steven@1000phone.com", key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(cipherText)
//
//	originalText, err := DesDecryptString(cipherText, key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(originalText)
//}
//
//func Test3Des() {
//	key := []byte("123456781234567812345678")
//	result, err := TripleDesEncrypt([]byte("steven@1000phone.com"), key)
//	if err != nil {
//		panic(err)
//	}
//	origData, err := TripleDesDecrypt(result, key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(origData))
//
//	fmt.Println("------------\n3DES解密字符串")
//	cipherText, err := TripleDesEncrypt2Str("Steven陪你学区块链", key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(cipherText)
//
//	originalText, err := TripleDesDecrypt2Str(cipherText, key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(originalText)
//}
