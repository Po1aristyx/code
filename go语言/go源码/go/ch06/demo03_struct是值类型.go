package main

import "fmt"

type Human struct {
	name string
	age  int8
	sex  byte
}

func main() {
	//1、初始化Human
	h1 := Human{"Steven", 35, 1}
	fmt.Printf("h1: %T , %v, %p \n", h1, h1, &h1)
	fmt.Println("--------------------")

	//将结构体对象拷贝
	h2 := h1
	h2.name = "David"
	h2.age = 30
	fmt.Printf("h2修改后: %T , %v, %p \n", h2, h2, &h2)
	fmt.Printf("h1: %T , %v, %p \n", h1, h1, &h1)
	fmt.Println("--------------------")

	//将结构体对象作为参数传递
	changeName(h1)
	fmt.Printf("h1: %T , %v, %p \n", h1, h1, &h1)
}

func changeName(h Human) {
	h.name = "Daniel"
	h.age = 13
	fmt.Printf("函数体内修改后：%T， %v ， %p \n" , h , h , &h)
}
