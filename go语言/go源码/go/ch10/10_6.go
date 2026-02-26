package main

import (
	"net/http"
)

func main() {
	testFileServer()
}
func testFileServer() {
	//如果该文件里面有index.html文件，他会优先显示html文件，否则会看到文件目录
	http.ListenAndServe(":2003", http.FileServer(http.Dir("D:\\QQ数据\\go源码\\go\\ch10\\public")))
}
