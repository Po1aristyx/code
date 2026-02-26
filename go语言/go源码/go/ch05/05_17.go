package main
import (
	"container/list"
	"fmt"
)
func main() {
	copyList()
}

//list是值类型，不过采用New()方法声明的是一个指针。所以在拷贝操作和参数传递时具有引用类型的特征。
func copyList() {
	//声明list
	list1 := list.New()
	printListInfo2("刚声明的list1", list1)

	//list1赋值
	list1.PushBack("one")
	list1.PushBack(2)
	list1.PushBack("三")
	printListInfo2("赋值后的list1", list1)
	iterateList2(list1)

	//将list1拷贝给list2。其实拷贝的仅仅是list1的地址
	list2 := list1
	printListInfo2("刚拷贝的list2", list2)
	iterateList2(list2)

	//修改list2
	list2.PushBack(250)
	list2.PushBack(350)
	list2.PushBack(450)
	printListInfo2("修改后的list2", list2)
	iterateList2(list2)

	//	list2的修改是否会影响list1
	printListInfo2("修改list2后的list1", list1)
	iterateList2(list1)

	//list2调用Init之后
	list2.Init()
	printListInfo2("list2调用Init之后", list1)
	iterateList2(list1)
}

func printListInfo2(info string, l *list.List) {
	fmt.Println(info + "-------------")
	fmt.Printf("%T , %v \t ，长度：%d \n", l, l, l.Len())
	fmt.Println("--------------")
}

//顺序遍历
func iterateList2(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v \t", e.Value)
	}
	fmt.Println("\n----------------")
}