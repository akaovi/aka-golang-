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
	fmt.Printf("排序前的数组：%v\n", arr)
	return arr
}

// func find(num int, arr []int) bool {
// 	for _, v := range arr {
// 		if v == num {
// 			return true
// 		}
// 	}
// 	return false
// }

func quick_sort(arr []int) {
	if len(arr) > 1 {
		var left = 0
		var right = len(arr) - 1
		var temp = arr[left]
		i := left
		j := right
		flag := 2
		for {
			// fmt.Printf("%d", left)
			// fmt.Printf("%d", i)
			if left >= right {
				break
			} else {
				if j == left {
					left++
					i = left
					temp = arr[i]
					j = right
				} else if i == j {
					arr[i] = temp
					i = left
					j = right
					flag = 2
					temp = arr[left]
				} else if flag == 2 {
					if arr[j] >= temp {
						j--
					} else {
						arr[i] = arr[j]
						flag = 1
					}
				} else if flag == 1 {
					if arr[i] <= temp {
						i++
					} else {
						arr[j] = arr[i]
						flag = 2
					}
				}
			}
		}
	}
	fmt.Printf("排序后的数组：%v\n", arr)
}

func main() {
	fmt.Println("请输入需要的检验排序数组中数的个数：")
	var num int
	fmt.Scanln(&num)

	//arr2 := []int{53, 53, 53}
	// for k := 2; k < 21; k++ {
	// 	var num int = k
	// 	for i := 0; i < 100; i++ {
	// 		arr2 := make_arr(num)
	// 		temp := arr2
	// 		quick_sort(arr2)
	// 		for j := 0; i < len(arr2); i++ {
	// 			if find(temp[j], arr2) {
	// 				fmt.Printf("%v存在", temp[j])
	// 			} else {
	// 				fmt.Printf("排序在%d存在错误!", k)
	// 				break
	// 			}
	// 		}
	// 	}
	// }

	arr2 := make_arr(num)
	quick_sort(arr2)
}
