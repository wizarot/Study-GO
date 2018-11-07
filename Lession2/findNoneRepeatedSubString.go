package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int) // map记录每个字母,最后出现的位置
	start := 0                         // start 指针,记录当前串开始位置
	maxLength := 0
	for i, ch := range []rune(s) { // 遍历字符串中的值,注意ch值类型是int32; 我们做一下转换将string转为 []byte

		if lastI ,ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength // 返回最大不重复字串长度
}

/**
找出最大长度不重复字串 leatcode著名问题你
理解rune类型
 */
func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdefg"))
	// 如果是中文呢?
	fmt.Println(lengthOfNonRepeatingSubStr("这里是中文"))// 直接中文就不对了,因为遍历时被转为byte是会出问题的..
	fmt.Println(lengthOfNonRepeatingSubStr("一二三二一"))// 将字符串类型转为rune来代替byte就正常了!
	fmt.Println(lengthOfNonRepeatingSubStr("Yes我爱GO语言!🏀⚽⚡👄👍🔥"))// 看一下复杂版本
}
