package main
import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)
func main(){
	testHttpPost()
}
func testHttpPost() {
	//构建参数
	data := url.Values{
		"theCityName": {"重庆"},
	}
	//参数转化成body
	reader := strings.NewReader(data.Encode())
	//发起post请求 MIME格式
	//http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName?thecityname=%E5%A4%A9%E6%B4%A5
	response, err := http.Post("http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName",
		"application/x-www-form-urlencoded", reader)
	CheckErr(err)
	fmt.Printf("响应状态码: %v\n", response.StatusCode)
	if response.StatusCode == 200 {
		//操作响应数据
		defer response.Body.Close()
		fmt.Println("网络请求成功")
		CheckErr(err)

	} else {
		fmt.Println("请求失败", response.Status)
	}
}
func CheckErr(err error) {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("程序出现异常：", ins.Error())
		}
	}()
	if err != nil {
		panic(err)
	}
}
