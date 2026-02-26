package main

import "fmt"

func main() {
	funcA()
	funcB()
	funcC()
	fmt.Println("main over")
}

func funcA() {
	fmt.Println("这是funcA")
}

func funcB() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("恢复啦，获取recover的返回值:", msg)
		}

	}()
	fmt.Println("这是funcB")
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			//panic("funcB恐慌啦")
		}
	}
}

func funcC() {
	defer func() {
		fmt.Println("执行延迟函数")
		msg := recover()
		fmt.Println("获取recover的返回值：", msg)
	}()
	fmt.Println("这是funcC")
	panic("funcC恐慌了")
}
