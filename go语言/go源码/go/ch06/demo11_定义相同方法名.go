package main

import (
	"math"
	"fmt"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func main() {
	r1:=Rectangle{10,4}
	r2:=Rectangle{12,5}
	c1:=Circle{1}
	c2:=Circle{10}

	fmt.Println("r1的面积" , r1.Area())
	fmt.Println("r2的面积" , r2.Area())
	fmt.Println("c1的面积" , c1.Area())
	fmt.Println("c2的面积" , c2.Area())
}

//定义Rectangle的方法
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

//定义Circle的方法
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
