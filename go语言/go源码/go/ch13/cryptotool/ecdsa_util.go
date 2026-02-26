package cryptotool

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)

//生成私钥和公钥，生成的私钥为结构体ecdsa.PrivateKey的指针
func NewKeyPair() (ecdsa.PrivateKey, []byte) {

	//生成secp256椭圆曲线
	curve := elliptic.P256()
	//产生的是一个结构体指针，结构体类型为ecdsa.PrivateKey
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("私钥：%x\n", private)
	fmt.Printf("私钥X：%x\n", private.X.Bytes())
	fmt.Printf("私钥Y：%x\n", private.Y.Bytes())
	fmt.Printf("私钥D：%x\n", private.D.Bytes())

	//x坐标与y坐标拼接在一起，生成公钥
	pubKey := append(private.X.Bytes(), private.Y.Bytes()...)
	//打印公钥，公钥用16进制打印出来长度为128，包含了x轴坐标与y轴坐标。
	fmt.Printf("公钥：%x \n", pubKey)

	return *private, pubKey
}

//生成签名的DER格式
func MakeSignatureDerString(r, s string) string {
	// 获取R和S的长度
	lenSigR := len(r) / 2
	lenSigS := len(s) / 2

	// 计算DER序列的总长度
	lenSequence := lenSigR + lenSigS + 4

	// 将10进制长度转16进制字符串
	strLenSigR := DecimalToHex(int64(lenSigR))
	strLenSigS := DecimalToHex(int64(lenSigS))
	strLenSequence := DecimalToHex(int64(lenSequence))

	// 拼凑DER编码
	derString := "30" + strLenSequence
	derString = derString + "02" + strLenSigR + r
	derString = derString + "02" + strLenSigS + s
	derString = derString + "01"
	return derString
}

//验证签名1
func VerifySig(pubKey, message []byte, r, s *big.Int) bool {
	curve := elliptic.P256()

	//公钥的长度
	keyLen := len(pubKey)

	//前一半为x轴坐标，后一半为y轴坐标
	x := big.Int{}
	y := big.Int{}
	x.SetBytes(pubKey[:(keyLen / 2)])
	y.SetBytes(pubKey[(keyLen / 2):])
	rawPubKey := ecdsa.PublicKey{curve, &x, &y}

	//根据交易哈希、公钥、数字签名验证成功。ecdsa.Verify func Verify(pub *PublicKey, hash []byte, r *big.Int, s *big.Int) bool
	res := ecdsa.Verify(&rawPubKey, message, r, s)
	return res
}

//验证签名2
func VerifySignature(pubKey, message []byte, r, s string) bool {
	curve := elliptic.P256()

	//公钥的长度
	keyLen := len(pubKey)

	//前一半为x轴坐标，后一半为y轴坐标
	x := big.Int{}
	y := big.Int{}
	x.SetBytes(pubKey[:(keyLen / 2)])
	y.SetBytes(pubKey[(keyLen / 2):])
	rawPubKey := ecdsa.PublicKey{curve, &x, &y}

	//根据交易哈希、公钥、数字签名验证成功。ecdsa.Verify func Verify(pub *PublicKey, hash []byte, r *big.Int, s *big.Int) bool

	rint := big.Int{}
	sint := big.Int{}

	rByte, _ := hex.DecodeString(r)
	sByte, _ := hex.DecodeString(s)
	rint.SetBytes(rByte)
	sint.SetBytes(sByte)

	//fmt.Println("------", rint.SetBytes(rByte))
	//fmt.Println("------", sint.SetBytes(sByte))

	res := ecdsa.Verify(&rawPubKey, message, &rint, &sint)
	return res
}
