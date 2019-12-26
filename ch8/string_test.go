package my_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 空字符串
	s = "hello"
	t.Log(len(s))          // 长度5
	s = "\xE4\xB8\xA5"     // 可以存储任何二进制数
	t.Log(s)               // 中文"严"
	t.Log(len(s))          // 长度3,并不一定是字符数
	s = "\xE4\xBA\xBB\xFF" // 乱码任意二进制
	t.Log(s)               // 输出乱码 亻�
	t.Log(len(s))          // 长度4,并不一定是字符数
	// s[1] = 3 // s是不可变的byte slice,因此这样是不可以的!

}

func TestUnicode(t *testing.T) {
	var s string
	s = "中"
	t.Log(s)      // 中
	t.Log(len(s)) // 3

	c := []rune(s)               // 转化为rune类型切片
	t.Log(len(c))                // 长度为1
	t.Logf("中 unicode %x", c[0]) // unicode 0x4E2D
	t.Logf("中 UTF8 %x", s)       // UTF-8 0xE4B8AD  string/[]byte [0xE4,0xB8,0xAD]
}

func TestTravelString(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c) // [1]代表都是第1号位置的变量, 会出中文字
	}
}

func TestStringFn(t *testing.T) { // string相关函数
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)                    // [A B C]
	t.Log(strings.Join(parts, "-")) // A-B-C
	s = strconv.Itoa(10)            // 整型转字符串
	t.Log("str:" + s)               // str:10
	if i, err := strconv.Atoi("13"); err == nil {
		t.Log(10 + i) // 做一个加法, 23
	}
}
