package my_map

import "testing"

func TestAccessNotExistKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) // 0
	m1[2] = 0
	t.Log(m1[2]) // 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not exist")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

func TestMapWithFunction(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	if mySet[1] { // 判断元素是否存在
		t.Log("1 is exist") // 存在
	} else {
		t.Log("1 is not exist")
	}
	mySet[3] = true
	t.Log(len(mySet)) // 长度为2
	delete(mySet, 1)  // 删掉元素1
	if mySet[1] {     // 判断元素是否存在
		t.Log("1 is exist")
	} else {
		t.Log("1 is not exist") // 不存在
	}
	t.Log(len(mySet)) // 长度为1,删除了一个元素
}
