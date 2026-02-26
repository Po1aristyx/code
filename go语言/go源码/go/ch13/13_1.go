package main
import (
	"encoding/hex"
	"fmt"
)
func main(){
	arr := []byte{'1', '0', '0', '0','p', 'h', 'o' ,'n', 'e'}
	fmt.Println(string(arr))
	str :=BytesToHexString(arr)
	fmt.Println(str)
	str = ReverseHexString(str)
	arr,_ = HexStringToBytes(str)
	fmt.Printf("%x\n", arr)
	ReverseBytes(arr)
	fmt.Println(string(arr))
}
/**
 * 将字节数组转成16进制字符串： []byte -> string
 */
func BytesToHexString(arr []byte) string {
   return hex.EncodeToString(arr)
}
/**
 * 将16进制字符串转成字节数组： hex string ->  []byte
 */
func HexStringToBytes(s string) ([]byte, error) {
   arr, err := hex.DecodeString(s)
   return arr, err
}
/**
 * 16进制字符串大端和小端颠倒
 */
func ReverseHexString(hexStr string) string {
   arr, _ := hex.DecodeString(hexStr)
   ReverseBytes(arr)
   return hex.EncodeToString(arr)
}
/**
 * 字节数组大端和小端颠倒
 */
func ReverseBytes(data []byte) {
   for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
      data[i], data[j] = data[j], data[i]
   }
}



