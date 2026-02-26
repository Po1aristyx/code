package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	name := strings.TrimPrefix(req.URL.Path, "/hello/")
	fmt.Fprintf(w, "Hello %s", name)
}

func main() {
	http.HandleFunc("/hello/", helloHandler)

	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
