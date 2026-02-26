package main
import (
	"time"
	"fmt"
)
//1、定义结构体，表示自定义错误的类型
type MyError struct {
	When time.Time
	What string
}
//2、实现Error()方法
func (e MyError) Error() string {
	return fmt.Sprintf("%v : %v", e.When, e.What)
}
//3、定义函数，返回error对象。该函数求矩形面积
func getArea(width, length float64) (float64, error) {
	errorInfo := ""
	if width < 0 && length < 0 {
		errorInfo = fmt.Sprintf("长度：%v, 宽度：%v ， 均为负数", length, width)
	} else if length < 0 {
		errorInfo = fmt.Sprintf("长度：%v, 出现负数 ", length)
	} else if width < 0 {
		errorInfo = fmt.Sprintf("宽度：%v ， 出现负数", width)
	}
	if errorInfo != "" {
		return 0, MyError{time.Now(), errorInfo}
	} else {
		return width * length, nil
	}
}
func main() {
	res , err := getArea(-4, -5)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("面积为：" , res)
	}
}
