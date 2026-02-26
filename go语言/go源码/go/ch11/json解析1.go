package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string
	Authors string
	Publisher string
	IsPublished  bool
	Price float64
}

func main() {
	gobook := Book {
		"Go语言编程",
		"Steven",
		"1000phone",
		true,
		89.99,
	}
	b , _ := json.Marshal(gobook)
	fmt.Println(string(b))


}
