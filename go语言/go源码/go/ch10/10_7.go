package main

import (
	"fmt"
	"net/http"
)

func main() {
	//绑定路径，去触发方法
	http.HandleFunc("/index", indexHandler)
	//绑定端口
	//第一个参数为监听地址，第二个参数表示服务端处理程序，通常为nil，这意味着服务端调用http.DefalutServeMux进行处理。
	err := http.ListenAndServe("localhost:3013", nil)
	fmt.Println(err)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/index=========")
	w.Write([]byte("这是默认首页"))
}
