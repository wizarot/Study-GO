package try_test

import "testing"

func TestSliceInt(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0)) // 0,0
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0)) // 1,1

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1)) // 4,4

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2)) // 3,5
}
