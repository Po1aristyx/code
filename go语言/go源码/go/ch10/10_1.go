package main

import (
	"fmt"
	"net/http"
)

func main() {
	testHttpNewRequest()
}
func testHttpNewRequest() {
	//1.创建一个客户端
	client := http.Client{}
	//2.创建一个请求。请求方式既可以是Get，也可以是Post。
	request, err := http.NewRequest("GET", "https://www.baidu.com/search/suggest/initial_page/", nil)
	CheckErr(err)
	//3.客户端发送请求
	cookName := &http.Cookie{Name: "username", Value: "Steven"}
	//添加cookie
	request.AddCookie(cookName)
	response, err := client.Do(request)
	CheckErr(err)
	//设置请求头
	request.Header.Set("Accept-Lanauage", "zh-cn")
	defer response.Body.Close()
	//查看请求头的数据
	fmt.Printf("Header:%+v\n", request.Header)
	fmt.Printf("响应状态码: %v\n", response.StatusCode)
	//4操作数据
	if response.StatusCode == 200 {
		//data, err := ioutil.ReadAll(response.Body)
		fmt.Println("网络请求成功")
		CheckErr(err)
		//fmt.Println(string(data))
	} else {
		fmt.Println("网络请求失败", response.Status)
	}
}

// 检查错误
func CheckErr(err error) {
	//fmt.Println("09---------------")
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("程序出现异常：", ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}
