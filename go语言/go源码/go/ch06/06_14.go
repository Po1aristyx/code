package main
import (
	"math/rand"
	"time"
	"fmt"
)
func main() {
	playGame()
}
func playGame() {
	//获取随机数
	target := generateRandNum(10, 100)
	fmt.Println("请输入随机数：")
	fmt.Println("--------------------")

	//记录猜测的次数
	count := 0
	for {
		count++
		yourNum := 0
		fmt.Scanln(&yourNum)

		if yourNum < target {
			fmt.Println("小了❌")
		} else if yourNum > target {
			fmt.Println("大了❌")
		} else {
			fmt.Println("正确✅")
			fmt.Printf("您一共猜测了%d次！ \n", count)
			fmt.Println("---------------")
			playGame()
		}
		//错误提示
		alertInfo(count, target)
	}
}
//错误提示
func alertInfo(count, target int) {
	if count >= 6 {
		fmt.Printf("您一共猜了 %d 次都没有猜中，太笨了！😓 ", count)
		fmt.Println("正确数字：", target)
		fmt.Println("-------------")
		playGame()
	}
}
//生成随机数
func generateRandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
