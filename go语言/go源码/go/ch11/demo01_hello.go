package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/upload" , uploadHandler)
	http.HandleFunc("/hello" , helloHandler)
	err := http.ListenAndServe(":8080" , nil)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func helloHandler(w http.ResponseWriter , r *http.Request) {
	io.WriteString(w , "Welcome...")
	io.WriteString(w , "Steven  王向军...")
}

func uploadHandler(w http.ResponseWriter , r *http.Request) {
	io.WriteString(w , "Welcome...")
	io.WriteString(w , "Steven  王向军...")
}
