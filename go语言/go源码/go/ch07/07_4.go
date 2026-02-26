package main
import (
	"fmt"
)
type Dog struct {
	name  string
	color string
	age   int8
	kind  string
}
func main() {
	//1、实现结构体的深拷贝
	//struct是值类型，默认的复制就是深拷贝
	d1 := Dog{"豆豆", "黑色", 2, "二哈"}
	fmt.Printf("d1: %T , %v , %p \n", d1, d1, &d1)
	d2 := d1 //深拷贝
	fmt.Printf("d2: %T , %v , %p \n", d2, d2, &d2)

	d2.name = "毛毛"
	fmt.Println("d2修改后：", d2)
	fmt.Println("d1：", d1)//d1没有因为d2的改变而改变，所以是深拷贝
	fmt.Println("------------------")

	//2、实现结构体浅拷贝：直接赋值指针地址
	d3 := &d1
	fmt.Printf("d3: %T , %v , %p \n", d3, d3, d3)
	d3.name = "球球"
	d3.color = "白色"
	d3.kind = "萨摩耶"
	fmt.Println("d3修改后：", d3)
	fmt.Println("d1：", d1)//d1因为d3的改变而改变，所以是浅拷贝
	fmt.Println("------------------")

	//3、实现结构体浅拷贝：通过new()函数来实例化对象
	d4 := new(Dog)
	d4.name = "多多"
	d4.color = "棕色"
	d4.age = 1
	d4.kind = "巴哥犬"
	d5 := d4
	fmt.Printf("d4: %T , %v , %p \n", d4, d4, d4)
	fmt.Printf("d5: %T , %v , %p \n", d5, d5, d5)

	fmt.Println("------------------")

	d5.color = "金色"
	d5.kind = "金毛"
	fmt.Println("d5修改后：", d5)
	fmt.Println("d4：", d4)
	fmt.Println("------------------")
}