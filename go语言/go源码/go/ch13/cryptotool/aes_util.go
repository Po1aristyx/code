package cryptotool

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//AES加密字节数组，返回字节数组
func AesEncrypt(originalBytes, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	originalBytes = PKCS5Padding(originalBytes, blockSize)
	// originalBytes = ZeroPadding(originalBytes, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	cipherBytes := make([]byte, len(originalBytes))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := originalBytes
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
	// origData := cipherBytes
	blockMode.CryptBlocks(originalBytes, cipherBytes)
	originalBytes = PKCS5UnPadding(originalBytes)
	// origData = ZeroUnPadding(origData)
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

//func TestAes() {
//	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
//	key := []byte("stevenwang@12345")
//	result, err := AesEncrypt([]byte("Steven陪你学区块链"), key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(base64.StdEncoding.EncodeToString(result))
//	origData, err := AesDecrypt(result, key)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(origData))
//}
