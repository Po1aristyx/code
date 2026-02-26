package cryptotool

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"log"
)

const PrivKeyBytesLen = 32

type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
}

func NewKeyPair2() *KeyPair {
	private, public := newKeyPair2()
	keyPair := KeyPair{private, public}
	return &keyPair
}

func newKeyPair2() ([]byte, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	d := private.D.Bytes()
	b := make([]byte, 0, PrivKeyBytesLen)
	priKey := paddedAppend(PrivKeyBytesLen, b, d)
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	pubKey, _ = hex.DecodeString("04" + hex.EncodeToString(pubKey))
	//xstring := fmt.Sprintf("%x", private.PublicKey.X.Bytes())
	//ystring := fmt.Sprintf("%x", private.PublicKey.Y.Bytes())
	//fmt.Println("X=", xstring)
	//fmt.Println("Y=", ystring)
	return priKey, pubKey
}

// paddedAppend appends the src byte slice to dst, returning the new slice.
// If the length of the source is smaller than the passed size, leading zero
// bytes are appended to the dst slice before appending src.
func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}
