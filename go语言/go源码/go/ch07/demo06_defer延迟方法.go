package main

import "fmt"

type person struct {
	firstName, lastName string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func main() {
	p := person{"Steven" , "Wang"}
	defer p.fullName()
	fmt.Print("Welcome ")
}
