package ch11

import (
	"fmt"
	"testing"
)

type Empployee struct {
	Id   string
	Name string
	Age  int
}

// 行为的定义
// 这种定义方式,在实例被调用时,实例的成员会进行复制
// func (e Empployee) String() string {
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

// 通常情况下,为了避免内存拷贝我们使用下面这种定义方式
func (e *Empployee) String() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStruct(t *testing.T) {
	e := Empployee{"0", "Bob", 30}
	e1 := Empployee{Name: "Mike", Age: 20}
	e2 := new(Empployee) // 返回指针
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 22
	t.Log(e)               // {0 Bob 30}
	t.Log(e1)              // { Mike 20}
	t.Log(e1.Id)           //
	t.Log(e2)              // &{2 Rose 22}
	t.Logf("e is %T", e)   // e is ch11.Empployee
	t.Logf("e2 is %T", e2) // e2 is *ch11.Empployee
}
