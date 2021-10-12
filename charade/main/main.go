package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("正在载入游戏~~~")
	// 初始化
	var guess_num int
	var count byte = 0
	var life byte = 3
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(int(101))

	fmt.Println("游戏开始！")
	fmt.Printf("你共有%d次机会!\n", life)
	// fmt.Println("随机生成的数字为：", num)

	for {
		count += 1
		fmt.Println("请输入一个0-100的数字：")
		fmt.Scanln(&guess_num)

		if guess_num == num {
			fmt.Printf("恭喜你，经过%d次猜对了！\n游戏结束", count)
			break
		} else {
			life -= 1
			if life == 0 {
				fmt.Println("所剩次数为0，游戏结束！")
				break
			} else {
				fmt.Printf("猜错了哦，当前剩余%d次机会,", life)
			}
		}
	}
}
