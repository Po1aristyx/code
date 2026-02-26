package main
import (
	"fmt"
)
type Person struct {
	name string
	age  int
	sex  string
}
type Student struct {
	Person
	schoolName string
}
func main() {
	//1、实例化并初始化Person
	p1 := Person{"Steven", 35, "男"}
	fmt.Println(p1)
	fmt.Println("-------------------")

	//2、实例化并初始化Student
	//	写法1：
	s1 := Student{p1, "北航软件学院"}
	printInfo(s1)

	//写法2：
	s2 := Student{Person{"Josh", 30, "男"}, "北外高翻学院"}
	printInfo(s2)

	//写法3：
	s3 := Student{Person: Person{
		name: "Penn",
		age:  19,
		sex:  "男",
	},
		schoolName: "北大元培学院",
	}
	printInfo(s3)

	//	写法4：
	s4 := Student{}
	s4.name = "Daniel"
	s4.sex = "男"
	s4.age = 12
	s4.schoolName = "北京十一龙樾"
	printInfo(s4)

}
func printInfo(s1 Student) {
	fmt.Println(s1)
	fmt.Printf("%+v \n", s1)
	fmt.Printf("姓名：%s， 年龄：%d ， 性别：%s ，学校：%s \n", s1.name, s1.age, s1.sex, s1.schoolName)
	fmt.Println("-------------------")
}