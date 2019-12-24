package try_test

import "testing"

const (
	Monday = iota +1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executalbe
)


func TestConstTry(t *testing.T)  {
	t.Log(Monday,Tuesday,Wednesday)
}

func TestConstTry2(t *testing.T)  {
	a := 7 // 0111
	t.Log(a&Readable==Readable, a&Writable==Writable,a&Executalbe==Executalbe)// 位检测

}
