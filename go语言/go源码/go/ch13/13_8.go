package main
import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)
func main() {
	fmt.Println("生成签名-------------------------------")
	privKey, pubKey := NewKeyPair()
	msg := sha256.Sum256([]byte("hello world"))
	r, s, _ := ecdsa.Sign(rand.Reader, &privKey, msg[:])
	strSigR := fmt.Sprintf("%x", r)
	strSigS := fmt.Sprintf("%x", s)
	fmt.Println("r、s的10进制分别为：", r, s)
	fmt.Println("r、s的16进制分别为：", strSigR, strSigS)
	signatureDer := MakeSignatureDerString(strSigR, strSigS)
	fmt.Println("数字签名DER格式为：", signatureDer)
	res := VerifySig(pubKey, msg[:], r, s)
	fmt.Println("签名验证结果：", res)
	res = VerifySignature(pubKey, msg[:], strSigR, strSigS)
	fmt.Println("签名验证结果：", res)
}
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.X.Bytes(), private.Y.Bytes()...)
	return *private, pubKey
}
func MakeSignatureDerString(r, s string) string {
	lenSigR := len(r) / 2
	lenSigS := len(s) / 2
	lenSequence := lenSigR + lenSigS + 4
	strLenSigR := DecimalToHex(int64(lenSigR))
	strLenSigS := DecimalToHex(int64(lenSigS))
	strLenSequence := DecimalToHex(int64(lenSequence))
	derString := "30" + strLenSequence
	derString = derString + "02" + strLenSigR + r
	derString = derString + "02" + strLenSigS + s
	derString = derString + "01"
	return derString
}
func VerifySig(pubKey, message []byte, r, s *big.Int) bool {
	curve := elliptic.P256()
	keyLen := len(pubKey)
	x := big.Int{}
	y := big.Int{}
	x.SetBytes(pubKey[:(keyLen / 2)])
	y.SetBytes(pubKey[(keyLen / 2):])
	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
	res := ecdsa.Verify(&rawPubKey, message, r, s)
	return res
}
func VerifySignature(pubKey, message []byte, r, s string) bool {
	curve := elliptic.P256()
	keyLen := len(pubKey)
	x := big.Int{}
	y := big.Int{}
	x.SetBytes(pubKey[:(keyLen / 2)])
	y.SetBytes(pubKey[(keyLen / 2):])
	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
	rint := big.Int{}
	sint := big.Int{}
	rByte, _ := hex.DecodeString(r)
	sByte, _ := hex.DecodeString(s)
	rint.SetBytes(rByte)
	sint.SetBytes(sByte)
	res := ecdsa.Verify(&rawPubKey, message, &rint, &sint)
	return res
}
func DecimalToHex(n int64) string {
	if n < 0 {
		log.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}