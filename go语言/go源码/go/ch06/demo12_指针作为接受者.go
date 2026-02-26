package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func main() {
	r1 := Rectangle{5 , 8}
	r2 := r1

	//打印对象的内存地址
	fmt.Printf("r1的地址：%p \n" , &r1)
	fmt.Printf("r2的地址：%p \n" , &r2)

	r1.setValue()
	fmt.Println("r1.height=" , r1.height)//8
	fmt.Println("r2.height=" , r2.height)//8
	fmt.Println("----------------")

	r1.setValue2()
	fmt.Println("r1.height=" , r1.height)//8?20?
	fmt.Println("r2.height=" , r2.height)//8
	fmt.Println("----------------")
}

func (r Rectangle) setValue() {
	fmt.Printf("setValue方法中r的地址：%p \n" , &r)
	r.height = 10
}

func (r *Rectangle) setValue2() {
	fmt.Printf("setValue2方法中r的地址：%p \n" , r)
	r.height = 20
}
