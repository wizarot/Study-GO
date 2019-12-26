package ch10

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDefer(t *testing.T)  {
	defer func() {
		t.Log("defer run at the end!")
	}()

	t.Log("Started")
	panic("Fatal error!")
}



func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5)) // 15
	t.Log(Sum(8, 9, 10, 11))  // 38
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 类似装饰器的方式,输出函数执行时长
func timeSpend(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)

	tsSf := timeSpend(slowFunc)
	t.Log(tsSf(5))
}
