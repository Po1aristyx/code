package main

import (
	"errors"
	"fmt"
)

func main() {
	//1、创建error对象的方式1
	err1 := errors.New("自己创建的错误！")
	fmt.Println(err1.Error())
	fmt.Println(err1)
	fmt.Printf("err1的类型：%T\n", err1) //*errors.errorString
	fmt.Println("-----------------")

	//2、创建error对象的方式2
	err2 := fmt.Errorf("错误的类型%d", 10)
	fmt.Println(err2.Error())
	fmt.Println(err2)
	fmt.Printf("err2的类型：%T\n", err2) //*errors.errorString
	fmt.Println("-----------------")

	//创建error对象的方式2
	res , err3 := checkAge(-12)
	if err3 != nil  {
		fmt.Println(err3.Error())
		fmt.Println(err3)
	} else {
		fmt.Println(res)
	}
}

//设计一个函数：验证年龄。如果是负数，则返回error
func checkAge(age int) (string, error) {
	if age < 0 {
		err := fmt.Errorf("您的年龄输入是：%d ， 该数值为负数，有错误！", age)
		return "", err
	} else {
		return fmt.Sprintf("您的年龄输入是：%d ", age), nil
	}
}
