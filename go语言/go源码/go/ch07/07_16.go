package main
import "fmt"
type ISayHello interface {
	SayHello() string
}
type Duck struct {
	name string
}
type Bird struct {
	name string
}
type Person struct {
	name string
}
func (d Duck) SayHello() string {
	return d.name + "叫：ga ga ga!"
}
func (p Person) SayHello() string {
	return p.name + "说：Hello!"
}
func main() {
	duck := Duck{"Yaya"}
	person := Person{"Steven"}
	fmt.Println(duck.SayHello())
	fmt.Println(person.SayHello())
	fmt.Println("------------------")
	var i ISayHello
	i = duck
	fmt.Printf("%T , %v , %p \n" , i , i , &i)
	fmt.Println(i.SayHello())
	i = person
	fmt.Printf("%T , %v , %p \n" , i , i , &i)
	fmt.Println(i.SayHello())
}
