package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"day09_http网络编程01/helper"
)

func main() {
	testHttpGet()

	//testClientGet()
}

func testHttpGet() {
	//获取服务器的数据
	response, err := http.Get("http://www.baidu.com")
	helper.CheckErr(err)

	//fmt.Printf("%+v\n", response)
	if response.StatusCode == 200 {
		//操作响应数据
		defer response.Body.Close()
		//读取body里面的数据
		content, err := ioutil.ReadAll(response.Body)
		helper.CheckErr(err)

		//打印响应体里面的数据
		fmt.Println(string(content))
		//io.Copy(os.Stdout , response.Body)

	} else {
		fmt.Println("请求失败", response.Status)
	}
}
