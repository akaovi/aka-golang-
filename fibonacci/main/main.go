package main

import "fmt"

func fibonacci(num byte) int64 {
	if num == 0 {
		return 0
	} else if num == 1 || num == 2 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func main() {
	var num byte
	fmt.Println("请输入要计算的斐波那契数列的项数：")
	fmt.Scanln(&num)
	result := fibonacci(num)
	fmt.Printf("第%d项斐波那契数为%d。", num, result)
}
