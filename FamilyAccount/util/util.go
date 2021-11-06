package utils

import "fmt"

type FamilyAccount struct {
	key     string
	balance float64
	money   float64
	loop    bool
	details string
	note    string
}

func NewFamilyAccout() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		balance: 10000.0,
		money:   0.0,
		loop:    true,
		details: "收支\t收支额\t余额\t说明\n",
		note:    "",
	}
}

func (self *FamilyAccount) showdetails() {
	fmt.Println("----------------收支详情----------------")
	fmt.Println(self.details)
}

func (self *FamilyAccount) income() {
	fmt.Println("请输入收入额:")
	fmt.Scanln(&self.money)
	self.balance += self.money
	fmt.Println("请输入收入说明:")
	fmt.Scanln(&self.note)
	self.details += fmt.Sprintf("收入\t%v\t%v\t%v\n", self.money, self.balance, self.note)
}

func (self *FamilyAccount) pay() {
	fmt.Println("请输入支出额:")
	fmt.Scanln(&self.money)
	if self.money <= self.balance {
		self.balance -= self.money
		fmt.Println("请输入支出说明:")
		fmt.Scanln(&self.note)
		self.details += fmt.Sprintf("支出\t%v\t%v\t%v\n", self.money, self.balance, self.note)
	} else {
		fmt.Println("支出大于余额或支出额错误，不能支出！")
	}
}

func (self *FamilyAccount) exit() {
	justice := ""
	fmt.Println("真的要退出吗 y/n")
	for {
		fmt.Scanln(&justice)
		if justice == "y" || justice == "Y" {
			self.loop = false
			break
		} else if justice == "n" || justice == "N" {
			fmt.Println("感谢您的不退出,老泪纵横！")
			break
		} else {
			fmt.Println("输入有误，请重新输入！")
		}
	}
}

func (self *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n----------------家庭收支----------------")
		fmt.Println("                 菜单")
		fmt.Println("             1.详细查询")
		fmt.Println("             2.收入登记")
		fmt.Println("             3.支出登记")
		fmt.Println("             4.退    出")

		fmt.Println("请输入(1~4):")
		fmt.Scanln(&self.key)

		switch self.key {
		case "1":
			self.showdetails()
		case "2":
			self.income()
		case "3":
			self.pay()
		case "4":
			self.exit()
		default:
			fmt.Println("输入有误，请输入(1~4):")
		}
		if !self.loop {
			fmt.Println("退出家庭收支系统！")
			break
		}
	}
}
