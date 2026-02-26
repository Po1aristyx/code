package main
import "fmt"
type A interface {
}
type Cat struct {
	name string
	age  int
}
type Person struct {
	name string
	sex  string
}
func main() {
	var a1 A = Cat{"Mimi", 1}
	var a2 A = Person{"Steven", "男"}
	var a3 A = "Learn golang with me!"
	var a4 A = 100
	var a5 A = 3.14
	showInfo(a1)
	showInfo(a2)
	showInfo(a3)
	showInfo(a4)
	showInfo(a5)
	fmt.Println("------------------")
	//1、fmt.println参数就是空接口
	fmt.Println("println的参数就是空接口，可以是任何数据类型", 100, 3.14, Cat{"旺旺", 2})
	//2、定义map。value是任何数据类型
	map1 := make(map[string]interface{})
	map1["name"] = "Daniel"
	map1["age"] = 13
	map1["height"] = 1.71
	fmt.Println(map1)
	fmt.Println("------------------")
	//	3、定义一个切片，其中存储任意数据类型
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, a5)
	fmt.Println(slice1)
}

func showInfo(a A) {
	fmt.Printf("%T , %v \n", a, a)
}
