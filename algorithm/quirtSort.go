package main

import "fmt"

// 快速排序(基本)
/*
func quickSort(arr []int) []int {
	// 就一个量也不用继续拆了.直接返回
	if len(arr) <= 1 {
		return arr
	}
	var left, right []int
	// 基准值,直接选第一个了.
	base := arr[0]
	for i := 1; i < len(arr); i++ {
		// fmt.Println("debug:", base, v)
		if arr[i] <= base {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	// fmt.Println("temp:", left, right, append(left, right...))
	return append(append(quickSort(left), base), quickSort(right)...)
}
*/

// 优化一下,直接玩切片指针
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return // 啥都不用做了
	}
	mid := arr[0]
	head, tail := 0, len(arr)-1
	for i := 1; i <= tail; {
		if arr[i] > mid {
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		} else {
			arr[i], arr[head] = arr[head], arr[i]
			head++
			i++
		}
	}
	quickSort(arr[:head])
	quickSort(arr[head+1:])
}

// 快速排序
func main() {
	arr := []int{9, 10, 8, 7, 5, 6, 4, 3, 2, 1, 2, 3, 5, 6}
	quickSort(arr)
	fmt.Println(arr)
}
