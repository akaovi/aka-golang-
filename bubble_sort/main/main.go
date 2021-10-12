package main

import (
	"fmt"
	"math/rand"
	"time"
)

func make_arr(length int) []int {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(101))
	}
	return arr
}

func Bubble_sort(arr *[]int) {
	fmt.Printf("排序前的数组：%v\n", *arr)
	temp := 0
	for j := 0; j < len(*arr)-1; j++ {
		for i := 0; i < len(*arr)-j-1; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			}
			if i == len(*arr)-1 {
				break
			}
		}
	}
	fmt.Printf("排序后的数组：%v\n", *arr)
}

func main() {
	// 检验排序数组中数的个数
	var num int
	fmt.Println("请输入需要的检验排序数组中数的个数：")
	fmt.Scanln(&num)

	var arr1 [5]int = [5]int{77, 57, 67, 17, 7}
	arr11 := arr1[:]
	Bubble_sort(&arr11)

	arr2 := make_arr(num)
	Bubble_sort(&arr2)
}
