package main
import (
	"fmt"
	"net/http"
)
func main(){
	testHttpGet()
}
func testHttpGet() {
	//获取服务器的数据
	response, err := http.Get("http://www.baidu.com")
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
