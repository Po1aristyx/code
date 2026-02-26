package cryptotool

import (
	"encoding/base64"
)

//const BASE64CODE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
//const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Base64EncodeString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64DecodeString(str string) string {
	result, _ := base64.StdEncoding.DecodeString(str)
	return string(result)
}
