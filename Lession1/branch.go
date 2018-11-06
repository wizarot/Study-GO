package main

import (
	"io/ioutil"
	"fmt"
)

func ifFunc() {
	const filename = "abc.txt"
	//contents , err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err) // 輸出錯誤信息
	//} else {
	//	fmt.Printf("%s\n",contents)// contents是一个byte[]
	//}

	// 可以写一起

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err) // 輸出錯誤信息
	} else {
		fmt.Printf("%s\n", contents) // contents是一个byte[]
	}
	//fmt.Printf("%s\n",contents)// contents作用域仅在if中,因此这里没法使用的.
}

func switchFunc(a, b int, op rune) int {
	var result int
	switch op {
	case '+':
		result = a + b // switch会自动break ,除非fallthrough
		//fallthrough 除非使用fallthrough, 节省很多代码啊!
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		result = a / b
	default:
		panic("不支持的操作符:" + string(op))
	}

	return result
}

func grade(score int) string {
	grade := ""
	switch { // switch后面可以没有条件语句
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score: %d", score))
	case score < 60:
		grade = "F"
	case score < 80:
		grade = "C"
	case score < 90:
		grade = "B"
	case score <= 100:
		grade = "A"
	default:
		panic(fmt.Sprintf("Wrong Score: %d", score))
	}

	return grade
}

// 分支语言结构
func main() {
	// if 结构
	ifFunc()
	// switch 分支
	fmt.Println(switchFunc(3, 5, '+'))
	fmt.Println(switchFunc(2, 2, '*')) // 单引号类型为rune ,双引号为string

	// 区分分支
	fmt.Println(grade(0),
		grade(20),
		grade(60),
		grade(80),
		grade(90),
		grade(100),
		//grade(101),// 异常数据
		)

}
