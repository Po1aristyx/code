package main
import (
	"fmt"
	"net/http"
)
func main(){
	testClientGet()
}
func testClientGet() {
	//创建客户端
	client := http.Client{}
	//通过client去请求
	response, err := client.Get("https://www.toutiao.com/search/suggest/initial_page")
	CheckErr(err)
	fmt.Printf("响应状态码: %v\n", response.StatusCode)

	if response.StatusCode == 200 {
		fmt.Println("网络请求成功")
		defer response.Body.Close()
		//处理
	}
}
//检查错误
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



