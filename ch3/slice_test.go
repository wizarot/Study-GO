package try_test

import "testing"

// func TestSliceInt(t *testing.T) {
// 	var s0 []int
// 	t.Log(len(s0), cap(s0)) // 0,0
// 	s0 = append(s0, 1)
// 	t.Log(len(s0), cap(s0)) // 1,1
//
// 	s1 := []int{1, 2, 3, 4}
// 	t.Log(len(s1), cap(s1)) // 4,4
//
// 	s2 := make([]int, 3, 5)
// 	t.Log(len(s2), cap(s2)) // 3,5
// }

func TestSliceShareMemory(t *testing.T) {
	year := []string{"1","2","3","4","5","6","7","8","9","10","11","12"}
	Q2 := year[3:6]
	t.Log(Q2,len(Q2),cap(Q2)) // [4 5 6] 3 9
	summer := year[5:8]
	t.Log(summer,len(summer),cap(summer)) // [6 7 8] 3 7
	summer[0] = "Unknown"
	t.Log(Q2) // [4 5 Unknown]
	t.Log(year) // [1 2 3 4 5 Unknown 7 8 9 10 11 12]
}
