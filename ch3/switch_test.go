package try_test

import "testing"

func TestSwitch(t *testing.T)  {
	for i:=0;i<5 ;i++  {
		switch  {
		case i%2==0:
			t.Log("偶數")
		case i%2==1:
			t.Log("奇數")
		default:
			t.Log("不知道")

		}
	}
}
