package main

import "fmt"

func bubbleSort(arr []int) []int {

	// 这个写法似乎和给人家给的动图解释有点不太一样,那么来个比较相近的..
	// for i, _ := range arr {
	// 	for j := i; j < len(arr); j++ {
	// 		if arr[i] < arr[j] {
	// 			arr[i], arr[j] = arr[j], arr[i]
	// 		}
	// 	}
	// }
	// 這個就比较像了.
	for i, _ := range arr {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	// 定义个切片
	arr := []int{9, 10, 8, 7, 5, 6, 4, 3, 2, 1, 2, 3, 5, 6}
	// 冒泡排个序
	fmt.Println(bubbleSort(arr))

}
