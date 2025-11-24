package main

import (
	"fmt"
)

type Person struct {
	name	string
	age		int
}


// 验证值拷贝和引用语义
func testRefCopy() {
	// 赋值-拷贝语义
	num1 := 1
	num2 := num1
	num2 = 2
	fmt.Printf("num1 = %d\nnum2 = %d\n", num1, num2)

	person1 := Person{"manba", 1}
	person2 := person1
	person2.age = 2
	fmt.Printf("person1: %#v\nperson2: %#v\n", person1, person2)
	
	// 创建引用类型-指针拷贝，共享底层数据
	mp1 := make(map[string]int)
	mp1["hajimi"] = 1
	mp2 := mp1
	mp2["haha"] = 2
	mp2["sdd"] = 2
	mp2["sffg"] = 2
	mp2["sffg"] = 2
	mp2["sfee"] = 2
	mp2["scv"] = 2
	mp2["hd"] = 2
	mp2["hahda"] = 2

	fmt.Printf("mp1: %#v\nmp2: %#v\n", mp1, mp2)

	// 用new创建切片
	ss := new([]int)
	fmt.Printf("%T\n", ss)
	fmt.Printf("%v\n", len(*ss))



}