package main

import "fmt"

func main() {
   //声明实际变量
   var a int = 120
   //声明指针变量
   var ip *int
      //给指针变量赋值，将变量a的地址赋值给ip
   ip = &a
   //打印a的类型和值
   fmt.Printf("a 的类型是%T，值是%v \n", a, a)
   //打印&a的类型和值
   fmt.Printf("&a 的类型是%T，值是%v \n", &a, &a)
   //打印ip的类型和值
   fmt.Printf("ip 的类型是%T，值是%v \n", ip, ip)
   //打印变量*ip的类型和值
   fmt.Printf("*ip 变量的类型是%T，值是%v \n", *ip, *ip)
   //打印变量*&a的类型和值
   fmt.Printf("*&a 变量的类型是%T，值是%v \n", *&a, *&a)
   fmt.Println(a, &a, *&a)
   fmt.Println(ip, &ip, *ip, *(&ip), &(*ip))
}

