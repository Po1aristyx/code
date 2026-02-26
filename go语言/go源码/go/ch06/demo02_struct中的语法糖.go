package main

import "fmt"

//定义结构体 Emp
type Emp struct {
	name string
	age  int8
	sex  byte
}

func main() {
	//使用new()内置函数实例化struct
	emp1 := new(Emp)
	fmt.Printf("emp1: %T , %v , %p \n", emp1, emp1, emp1)

	(*emp1).name = "David"
	(*emp1).age = 30
	(*emp1).sex = 1

	//语法糖写法
	emp1.name = "David2"
	emp1.age = 31
	emp1.sex = 1

	fmt.Println(emp1)
	fmt.Println("-----------------")

	SyntacticSugar()
}

func SyntacticSugar() {
	//	数组中的语法糖
	arr := [4]int{10, 20, 30, 40}
	arr2 := &arr
	fmt.Println((*arr2)[len(arr)-1])
	fmt.Println(arr2[0])

	//	切片中的语法糖？
	arr3 := []int{100, 200, 300, 400}
	arr4 := &arr3
	fmt.Println((*arr4)[len(arr)-1])
	//fmt.Println(arr4[0])
}
