package main

import (
	"fmt"
)

// AccountDetails 接口定义了获取账户详情的方法
type AccountDetails interface {
	GetDetails() string
}

// Account 基础账户结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 存款方法
func (account *Account) Deposit(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码不正确")
		return
	}

	if money <= 0 {
		fmt.Println("存款金额必须大于0")
		return
	}

	account.Balance += money
	fmt.Println("存款成功！")
}

// 取款方法
func (account *Account) Withdraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码不正确")
		return
	}

	if money <= 0 {
		fmt.Println("取款金额必须大于0")
		return
	}

	if money > account.Balance {
		fmt.Println("余额不足")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功！")
}

// 查询余额方法
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码不正确")
		return
	}

	fmt.Printf("账号: %s, 余额: %.2f元\n", account.AccountNo, account.Balance)
}

// PersonalAccount 个人账户
type PersonalAccount struct {
	Account
	Name     string
	IdNumber string
}

// GetDetails 实现AccountDetails接口
func (pa *PersonalAccount) GetDetails() string {
	return fmt.Sprintf("【个人账户】\n账号: %s\n户主姓名: %s\n身份证号: %s\n余额: %.2f元\n个人账户可享受基础服务费减免",
		pa.AccountNo, pa.Name, pa.IdNumber, pa.Balance)
}

// CorporateAccount 公司账户
type CorporateAccount struct {
	Account
	CompanyName string
	LicenseNo   string
}

// GetDetails 实现AccountDetails接口
func (ca *CorporateAccount) GetDetails() string {
	return fmt.Sprintf("【公司账户】\n账号: %s\n公司名称: %s\n营业执照号: %s\n余额: %.2f元\n公司账户可享受大额转账手续费优惠",
		ca.AccountNo, ca.CompanyName, ca.LicenseNo, ca.Balance)
}

// GovernmentAccount 机关账户
type GovernmentAccount struct {
	Account
	OrgName string
	OrgCode string
}

// GetDetails 实现AccountDetails接口
func (ga *GovernmentAccount) GetDetails() string {
	return fmt.Sprintf("【机关账户】\n账号: %s\n机构名称: %s\n机构代码: %s\n余额: %.2f元\n机关账户享受全免手续费特权",
		ga.AccountNo, ga.OrgName, ga.OrgCode, ga.Balance)
}

// 显示菜单
func showMenu() {
	fmt.Println("\n===== 银行账户管理系统 =====")
	fmt.Println("1. 存款")
	fmt.Println("2. 取款")
	fmt.Println("3. 查询余额")
	fmt.Println("4. 查看账户详情")
	fmt.Println("5. 退出")
	fmt.Print("请选择操作: ")
}

func main() {
	// 创建不同类型的账户
	accounts := []AccountDetails{
		&PersonalAccount{
			Account:  Account{AccountNo: "1001", Pwd: "123456", Balance: 5000.0},
			Name:     "tyx",
			IdNumber: "110101199001011234",
		},
		&CorporateAccount{
			Account:     Account{AccountNo: "2001", Pwd: "654321", Balance: 100000.0},
			CompanyName: "ABC科技有限公司",
			LicenseNo:   "91110105MA00001X1A",
		},
		&GovernmentAccount{
			Account: Account{AccountNo: "3001", Pwd: "888888", Balance: 500000.0},
			OrgName: "市政府财政部门",
			OrgCode: "12345678-9",
		},
	}

	fmt.Println("欢迎使用银行账户管理系统")

	// 登录验证
	var accountNo, pwd string
	var currentAccount AccountDetails
	var loggedIn bool

	for !loggedIn {
		fmt.Print("请输入账号: ")
		fmt.Scanln(&accountNo)
		fmt.Print("请输入密码: ")
		fmt.Scanln(&pwd)

		// 查找匹配的账户
		for _, acc := range accounts {
			switch a := acc.(type) {
			case *PersonalAccount:
				if a.AccountNo == accountNo && a.Pwd == pwd {
					currentAccount = a
					loggedIn = true
				}
			case *CorporateAccount:
				if a.AccountNo == accountNo && a.Pwd == pwd {
					currentAccount = a
					loggedIn = true
				}
			case *GovernmentAccount:
				if a.AccountNo == accountNo && a.Pwd == pwd {
					currentAccount = a
					loggedIn = true
				}
			}
		}

		if !loggedIn {
			fmt.Println("账号或密码错误，请重试")
		}
	}

	fmt.Println("登录成功！")

	// 根据账户类型获取对应的Account指针
	var acc *Account
	switch a := currentAccount.(type) {
	case *PersonalAccount:
		acc = &a.Account
	case *CorporateAccount:
		acc = &a.Account
	case *GovernmentAccount:
		acc = &a.Account
	}

	// 主循环
	for {
		showMenu()

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1: // 存款
			var amount float64
			fmt.Print("请输入存款金额: ")
			fmt.Scanln(&amount)
			acc.Deposit(amount, pwd)

		case 2: // 取款
			var amount float64
			fmt.Print("请输入取款金额: ")
			fmt.Scanln(&amount)
			acc.Withdraw(amount, pwd)

		case 3: // 查询余额
			acc.Query(pwd)

		case 4: // 查看账户详情
			fmt.Println("\n===== 账户详细信息 =====")
			fmt.Println(currentAccount.GetDetails())

		case 5: // 退出
			fmt.Println("感谢使用银行账户管理系统，再见！")
			return

		default:
			fmt.Println("无效的选择，请重试")
		}
	}
}
