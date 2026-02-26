package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

const MAX_GUESS = 5

func main() {
	// 猜数游戏
	startGuessGame()

	// 对话系统
	userInfo := collectUserInfo()
	showNews(userInfo)

}

// 用户信息结构
type UserInfo struct {
	Name       string
	Birthday   time.Time
	Birthplace string
}

// 猜数游戏（完整实现）
func startGuessGame() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	remaining := MAX_GUESS

	fmt.Printf("\n欢迎参加猜数游戏（1-100），您有%d次机会\n", MAX_GUESS)

	for remaining > 0 {
		var input string
		fmt.Print("请输入数字：")
		fmt.Scanln(&input)

		// 输入验证
		num, err := strconv.Atoi(input)
		if err != nil || num < 1 || num > 100 {
			fmt.Println("请输入1-100的整数")
			continue
		}

		remaining--
		switch {
		case num > target:
			fmt.Printf("猜大了，剩余次数：%d\n", remaining)
		case num < target:
			fmt.Printf("猜小了，剩余次数：%d\n", remaining)
		default:
			fmt.Println("恭喜您猜对了！")
			return
		}
	}

	fmt.Printf("游戏结束，正确答案是：%d\n", target)
}

// 收集用户信息（使用闭包）
func collectUserInfo() UserInfo {
	var info UserInfo

	// 姓名收集
	inputLoop("请输入您的姓名：", func(input string) bool {
		if len(input) > 0 {
			info.Name = input
			return true
		}
		return false
	})

	// 生日收集（带正则验证）
	inputLoop("请输入您的生日（YYYY-MM-DD）：", func(input string) bool {
		if matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, input); matched {
			t, err := time.Parse("2006-01-02", input)
			if err == nil {
				info.Birthday = t
				return true
			}
		}
		return false
	})

	// 出生地收集
	inputLoop("请输入您的出生地：", func(input string) bool {
		if len(input) > 0 {
			info.Birthplace = input
			return true
		}
		return false
	})

	return info
}

// 通用输入循环（带验证函数）
func inputLoop(prompt string, validate func(string) bool) {
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)
		if validate(input) {
			break
		}
		fmt.Println("输入不合法，请重新输入")
	}
}

// 新闻展示（使用函数组合）
func showNews(info UserInfo) {
	// 新闻生成器闭包
	newsGenerator := func(choice int) string {
		now := time.Now()
		switch choice {
		case 1:
			return fmt.Sprintf("《历史上的%02d月%02d日》：在%d年，%s诞生了一位杰出人物：%s。他会飞行。",
				info.Birthday.Month(), info.Birthday.Day(),
				1900+rand.Intn(120), info.Birthplace, info.Name)
		case 2:
			lastYear := now.AddDate(-1, 0, 0)
			return fmt.Sprintf("去年今日（%s）：在%s发生了重大科技突破，%s的科研团队成功研发了量子计算机。",
				lastYear.Format("2006年01月02日"), info.Birthplace, info.Name)
		}
		return ""
	}

	// 新闻选择
	var choice int
	inputLoop("请选择新闻内容：\n1.历史上的今日\n2.去年今日\n请输选择（1-2）：", func(input string) bool {
		n, err := strconv.Atoi(input)
		if err == nil && (n == 1 || n == 2) {
			choice = n
			return true
		}
		return false
	})

	fmt.Println("\n" + newsGenerator(choice) + "\n")
}
