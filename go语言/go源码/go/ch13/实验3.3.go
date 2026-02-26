package main

import (
	"fmt"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		username := req.URL.Query().Get("username")
		password := req.URL.Query().Get("password")

		loginInfo := fmt.Sprintf("您的登录信息：\n用户名： %s, 字符串长度为 %d\n密码： %s，字符串长度为 %d",
			username, len(username), password, len(password))

		fmt.Fprintf(w, loginInfo)
	}
}

func main() {
	http.HandleFunc("/login", loginHandler)

	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
