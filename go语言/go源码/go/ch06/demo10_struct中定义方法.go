package main

import "fmt"

type Employee struct {
	 name , currency string
	 salary float64
}
func main() {
	emp1 := Employee{"Daniel" , "$" , 2000}
	emp1.printSalary()
	printSalary(emp1)
}

//printSalary方法
func (e Employee) printSalary() {
	fmt.Printf("员工姓名：%s ，薪资：%s%.2f \n", e.name , e.currency , e.salary)
}

//printSalary函数
func printSalary(e Employee) {
	fmt.Printf("员工姓名：%s ，薪资：%s%.2f \n", e.name , e.currency , e.salary)
}
