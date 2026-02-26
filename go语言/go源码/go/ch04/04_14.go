package main
import "fmt"
type Student struct {
   name    string
   age     int
   married bool
   sex     int8
}
func main() {
   var s1 = Student{"Steven", 35, true, 1}
   var s2 = Student{"Sunny", 20, false, 0}
   var a *Student = &s1 //将s1的内存地址赋值给Student指针变量a
   var b *Student = &s2 //将s2的内存地址赋值给Student指针变量b
   fmt.Printf("s1类型为%T，值为%v \n", s1, s1)
   fmt.Printf("s2类型为%T，值为%v \n", s2, s2)
   fmt.Printf("a类型为%T，值为%v \n", a, a)
   fmt.Printf("b类型为%T，值为%v \n", b, b)
   fmt.Printf("*a类型为%T，值为%v \n", *a, *a)
   fmt.Printf("*b类型为%T，值为%v \n", *b, *b)
   fmt.Println(s1.name, s1.age, s1.married, s1.sex)
   fmt.Println(a.name, a.age, a.married, a.sex)
   fmt.Println(s2.name, s2.age, s2.married, s2.sex)
   fmt.Println(b.name, b.age, b.married, b.sex)
   fmt.Println((*a).name, (*a).age, (*a).married, (*a).sex)
   fmt.Println((*b).name, (*b).age, (*b).married, (*b).sex)
   fmt.Printf("&a类型为%T，值为%v\n", &a, &a)
   fmt.Printf("&b类型为%T，值为%v\n", &b, &b)
   fmt.Println(&a.name, &a.age, &a.married, &a.sex)
   fmt.Println(&b.name, &b.age, &b.married, &b.sex)
}

