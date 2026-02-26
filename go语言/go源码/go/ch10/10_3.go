package main
import (
	"fmt"
	"net/http"
	"net/url"
)
func main(){
	testClientPost()
}
func testClientPost() {
	//创建客户端
	client := http.Client{}
	//构建参数
	data := url.Values{
		"theCityName": {"北京"},
	}
	response, err := client.PostForm("http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName",
		data)
	CheckErr(err)
	fmt.Printf("响应状态码: %v\n", response.StatusCode)
	if response.StatusCode == 200 {
		defer response.Body.Close()
		fmt.Println("网络请求成功")
		CheckErr(err)
	}
}
//检查错误
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