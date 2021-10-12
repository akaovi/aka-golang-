package main

import "fmt"

func S_jinzita(tt int) {
	for i := 1; i <= tt; i++ {
		for j := 1; j <= tt*2-1; j++ {
			if j <= tt-i || j >= tt+i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func empty_jinzita(tt int) {
	for i := 1; i <= tt; i++ {
		for j := 1; j <= tt*2-1; j++ {
			if i == 1 || i == tt {
				if j <= tt-i || j >= tt+i {
					fmt.Print(" ")
				} else {
					fmt.Print("*")
				}
			} else {
				if j == tt-i+1 || j == tt+i-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

func nn_multiply() {
	// 左下角
	fmt.Println("左下角的九九乘法表：")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j <= i {
				fmt.Printf("%d*%d=%d\t", i, j, i*j)
			}
		}
		fmt.Println()
	}

	fmt.Println()

	// 右上角
	fmt.Println("右上角的九九乘法表：")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j < i {
				fmt.Print("     \t")
			} else {
				fmt.Printf("%d*%d=%d\t", i, j, i*j)
			}
		}
		fmt.Println()
	}
}

func main() {
	var tt int
	fmt.Println("请输入金字塔的层数：")
	fmt.Scanln(&tt)

	// 实心金字塔
	fmt.Println("实心金字塔如下：")
	S_jinzita(tt)
	fmt.Println()

	// 空心金字塔
	fmt.Println("空心金字塔如下：")
	empty_jinzita(tt)
	fmt.Println()
	// 九九乘法表
	nn_multiply()

	// 按任意键退出程序
	fmt.Print("按任意键退出程序！")
	fmt.Scanln()
}
