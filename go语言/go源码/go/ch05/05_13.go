package main

import "fmt"

func main() {
	countryMap := make(map[string]string)
	countryMap["China"] = "Beijing"
	countryMap["Japan"] = "Tokyo"
	countryMap["India"] = "New Delhi"
	countryMap["France"] = "Paris"
	countryMap["Italy"] = "Rome"
	countryMap["USA"] = "Hu"

	//查看元素是否在map中存在
	value, ok := countryMap["England"]
	fmt.Printf("%q \n", value)
	fmt.Printf("%T , %v \n", ok, ok)
	if ok {
		fmt.Println("首都：", value)
	} else {
		fmt.Println("首都信息未检索到！")
	}

	//或者
	if value, ok := countryMap["USA"]; ok {
		fmt.Println("首都：", value)
	} else {
		fmt.Println("首都信息未检索到！")
	}
}
