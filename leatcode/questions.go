package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int,len(nums)) // 生成一个不定长的map,无需声明长度,go自动分配.但如果能确定长度,或许可以提高执行效率
	for i, j := range nums { //遍历slice或map等的方法
		if pos, ok := m[target-j]; ok == true { //从map中通过key取值,如果不存在,那么可以通过第二个参数ok来判断
			return []int{pos, i}
		} else {
			m[j] = i
		}
	}
	return nil
}

// 并不准备执行,进用作笔记和格式化工具
func main() {

}
