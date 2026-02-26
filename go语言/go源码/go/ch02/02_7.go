package main

import "fmt"


func main(){
	arr := []byte{'x', 'y', 'z', 'Z'}
	//字符串
	fmt.Printf("%s \n", "欢迎大家学习区块链")
	fmt.Printf("%q \n", "欢迎大家学习区块链")
	fmt.Printf("%x\n", "欢迎大家学习区块链")
	fmt.Printf("%X \n", "欢迎大家学习区块链")
	fmt.Printf("%T , %s \n", arr, arr)
	fmt.Printf("%T , %q \n", arr, arr)
	fmt.Printf("%T , %x \n", arr, arr)
	fmt.Printf("%T , %X \n", arr, arr)

}
