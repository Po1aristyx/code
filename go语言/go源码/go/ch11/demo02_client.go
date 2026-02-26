package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	res , err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bytes , err := httputil.DumpResponse(res,false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n" , bytes)
}
