package try_test

import "testing"

func TestType(t *testing.T) {
	a := 1
	aPtr := &a // 取指针
	t.Log(a, aPtr)
	t.Logf("%T , %T ", a, aPtr)// %T 格式化输出变量类型
}

func TestString(t *testing.T)  {
	var s string
	t.Log("*"+s+"*") // 输出 **
	t.Log(len(s)) // 输出 0
}


const (
	Readable = 1 << iota
	Writable
	Executalbe
)


func TestConst(t *testing.T)  {
	a := 7 // 0111
	a = a &^ Readable // 清掉可读属性
	t.Log(a&Readable==Readable, a&Writable==Writable,a&Executalbe==Executalbe)// 位检测

}
