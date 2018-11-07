package main

import "fmt"

/**
maps:
map 的 key: 使用哈希表,因此必须可以比较相等; 除了Slice map function的内建类型都可以作为key; struct类型不包含上述字段,也可作为key
 */
func main() {
	// 定义map
	fmt.Println("map define:")
	m := map[string]string{
		"name":    "Will T",
		"course":  "Golang",
		"site":    "wizarot.me",
		"quality": "notpad",
	}

	m2 := make(map[string]int) // make是怎麼回事兒來著? m2 == empty map

	var m3 map[string]int // m3 == nil
	fmt.Println(m, m2, m3)
	// 遍历map , 顺序和定义时不同.是无序的hash map,每次执行遍历出来的顺序也不一定是一样的
	fmt.Println("map traversing:")
	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
	// 访问值
	fmt.Println("map getting values:")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	// 尝试访问不存在的key,并不会报错,只返回一个空的
	couresName, ok := m["coures"] // ok变量可以看是否真的存在
	fmt.Println(couresName, ok)
	if siteName, ok := m["sittte"]; ok { // 检测一下
		fmt.Println(siteName)
	} else {
		fmt.Println("key does not exist!")
	}

	// 删除元素
	fmt.Println("map deleteing values:")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
