package main
import "fmt"
type Address struct {
	province, city string
}
type Person struct {
	name    string
	age     int
	address *Address
}
func main() {
	//模拟结构体对象之间的聚合关系
	p := Person{}
	p.name = "Steven"
	p.age = 35
	//赋值方式1：
	addr := Address{}
	addr.province = "北京市"
	addr.city = "海淀区"
	p.address = &addr
	fmt.Println(p)
	fmt.Println("姓名：", p.name, " 年龄：", p.age, " 省：", p.address.province," 市：", p.address.city)
	fmt.Println("-----------------------")
	//修改Person对象的数据，是否会影响Address对象的数据？
	p.address.city = "昌平区"
	fmt.Println("姓名：", p.name," 年龄：", p.age," 省：", p.address.province," 市：", p.address.city," addr市：", addr.city)
	fmt.Println("-----------------------")
	//修改Address对象的数据，是否会影响Person对象的数据？
	addr.city = "大兴区"
	fmt.Println("姓名：", p.name," 年龄：", p.age," 省：", p.address.province," 市：", p.address.city," addr市：", addr.city)
	fmt.Println("-----------------------")
	//赋值方式2
	p.address = &Address{
		province: "陕西省",
		city:     "西安市",
	}
	fmt.Println(p)
	fmt.Println("姓名：", p.name, " 年龄：", p.age, " 省：", p.address.province," 市：", p.address.city)
}
