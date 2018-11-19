package main

import "fmt"

func selectSort(arr []int) []int {
	for i, _ := range arr {
		current := i // 指向最大/最小值的指针
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[current] {
				current = j // 每次找到最小的
			}
		}
		arr[i], arr[current] = arr[current], arr[i]
	}
	return arr
}

// 选择排序
func main() {
	arr := []int{9, 10, 8, 7, 5, 6, 4, 3, 2, 1, 2, 3, 5, 6}
	fmt.Println(selectSort(arr))
}
