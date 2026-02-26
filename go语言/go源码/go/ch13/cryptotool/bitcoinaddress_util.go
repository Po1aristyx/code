package cryptotool

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

/*
根据公钥生成比特币地址
参数：公钥、网络类型
网络id号 ● 比特币主网id号：0x00 ● testnet测试网络id号：0x6f
返回值：比特币地址字符串，error
*/
func GenerateAddress(pubkey string, nettype int) (string, error) {
	//1、判断公钥有效性
	if len(pubkey) != 130 && len(pubkey) != 66 {
		err := errors.New("公钥输入错误！")
		return "", err
	}

	//2、计算公钥sha256
	hashString := SHA256HexString(pubkey)
	//fmt.Println("sha256:", hashString)

	// 3、RipeMD160
	ripemd160String := HASH(hashString, "ripemd160" , true)
	//ripemd160String = fmt.Sprintf("%s", ripemd160String)
	//fmt.Println("ripemd160:", ripemd160String)

	// 4、增加网络id号 ● 比特币主网id号：0x00 ● testnet测试网络id号：0x6f
	prefix := ""
	switch nettype {
	case 0:
		prefix = "00"
	case 1:
		prefix = "6f"
	case 2:
		prefix = "34"
	}
	versionString := prefix + ripemd160String
	//fmt.Println("加前缀：", versionString)

	// 5、计算两次hash
	sha256DoubleString := SHA256Double(versionString, true)
	//fmt.Println(sha256DoubleString)

	// 6、获取校验码
	rs := []rune(sha256DoubleString)
	checknum := string(rs[:8])
	//fmt.Println(checknum)

	// 7、形成16进制比特币地址
	addrHex := versionString + checknum
	//fmt.Println(addrHex)

	// 8、对16进制地址进行Base58编码
	arr, _ := hex.DecodeString(addrHex)
	addrBase58 := Base58Encode(arr)
	result := fmt.Sprintf("%s \n", addrBase58)
	//fmt.Println(result)
	return result, nil
}

/*
* 验证WIF格式私钥是否有效
*
* 参数：String，wif格式私钥
 */
func ValidateWifPrivateKey(wifPrivateKey string) bool {
	arrPivateKey := Base58Decode([]byte(wifPrivateKey))
	hexString := fmt.Sprintf("%x\n", arrPivateKey)
	//fmt.Println(hexString)
	rs := []rune(hexString)
	checknum := string(rs[len(rs)-9 : len(rs)-1])
	//fmt.Println("checknum", checknum, len(checknum))

	versionStr := string(rs[:len(rs)-9])
	//fmt.Println("versionStr", versionStr)

	//生成新的校验码
	checkcode := GetCheckcode(versionStr)
	//fmt.Println(strings.Compare(checknum, checkcode))
	return strings.EqualFold(checknum, checkcode)
}

/*
* 将16进制私钥转成WIF格式或WIF-compressed格式
*
* 参数1：String，16进制私钥
*
* 参数2：boolean ,是否采用压缩格式
 */
func GeneratePrivateKeyWIF(hexPrivKey string, compressed bool) string {
	versionStr := ""
	// 获取校验码
	if compressed {
		versionStr = "80" + hexPrivKey + "01"
	} else {
		versionStr = "80" + hexPrivKey
	}
	//fmt.Println(versionStr)
	hashDouble := SHA256Double(versionStr, true)
	//fmt.Println(hashDouble)
	checknum := Substr(hashDouble, 0, 8)
	// 生成16进制私钥
	strPrivKey := versionStr + checknum
	//fmt.Println(strPrivKey)
	arr, _ := hex.DecodeString(strPrivKey)
	return string(Base58Encode(arr))
}

/*
 * 将WIF格式私钥返回到16进制格式
 *
 * 参数：String，WIF格式私钥
 *
 */
func GetPrivateKeyHexString(wifPrivateKey string) string {
	if ValidateWifPrivateKey(wifPrivateKey) {
		flag := true
		if len(wifPrivateKey) == 51 && strings.HasPrefix(wifPrivateKey, "5") {
			flag = false
		}
		//arr, _ := hex.DecodeString(wifPrivateKey)
		arr := Base58Decode([]byte(wifPrivateKey))
		hexString := hex.EncodeToString(arr)

		//fmt.Println("hexString:::" , hexString)

		result := ""
		if flag {
			result = Substring(hexString, 2, len(hexString)-10)
		} else {
			result = Substring(hexString, 2, len(hexString)-8)
		}
		return result
	}
	return ""
}
