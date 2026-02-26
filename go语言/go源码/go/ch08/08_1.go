package main
import (
	"errors"
	"math"
	"fmt"
	"os"
)
func main() {
	//	异常情况1
	res := math.Sqrt(-100)
	fmt.Println(res)
	res , err := Sqrt(-100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	//异常情况2
	//res = 100 / 0
	//fmt.Println(res)
	res , err = Divide(100 , 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
	//异常情况3
	f, err := os.Open("/abc.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f.Name() , "该文件成功被打开！")
	}

}
//定义平方根运算函数
func Sqrt(f float64)(float64 , error) {
	if f<0 {
		return 0 , errors.New("负数不可以获取平方根")
	} else {
		return math.Sqrt(f) , nil
	}
}
//定义除法运算函数
func Divide(dividee float64 , divider float64)(float64 , error) {
	if divider == 0 {
		return 0 , errors.New("出错：除数不可以为0！")
	} else {
		return dividee / divider , nil
	}
}