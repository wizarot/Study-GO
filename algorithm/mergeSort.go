package main

import "fmt"

func mergeSort(data []int) []int {
	if len(data) < 2 {
		return data // 不用继续了
	}
	// 从中间截成两半
	mid := len(data) / 2

	return merge(mergeSort(data[:mid]), mergeSort(data[mid:]))
}

// 合并时候排个序
func merge(left, right []int) []int {
	var result []int
	// fmt.Println("debug:", left, right)
	// 按顺序排个序 (合并算法有点问题啊!问题解决...这个处理方式有点意思.)
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			x := left[0]
			left = left[1:]
			result = append(result, x)
		} else {
			x := right[0]
			right = right[1:]
			result = append(result, x)
		}
	}

	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}

	// fmt.Println("merge:", result)
	return result
}

// 归并排序
func main() {
	arr := []int{9, 10, 8, 7, 5, 6, 4, 3, 2, 1, 2, 3, 5, 6}
	fmt.Println(mergeSort(arr))
}
