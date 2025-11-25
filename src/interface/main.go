package main

import (
	"fmt"
	"notes/src/interface/animal"
) 

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says: 汪汪！"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + " says: 喵喵~"
}

type Bird struct {
	Name string
}

func (b Bird) Speak() string {
	return b.Name + " says: jiji~"
} 


func main() {
	var a Animal
	a = Dog{Name: "旺财"}

	if dog, ok := a.(Dog); ok {
		fmt.Println(dog.Speak())
	} else {
		fmt.Println("真不是狗")
	}
	// 方式2：直接断言（如果确定类型，失败会panic）
	dog2 := a.(Dog)
	fmt.Println(dog2.Name)

	// 方式3：错误示例 - 断言成不匹配的类型
	if cat, ok := a.(Cat); ok {
		fmt.Println("真的是一只猫！", cat.Name)
	} else {
		fmt.Println("不是一只猫，是其他动物")
	}

	a = Bird{
		Name: "kunkun",
	}
	if birdPtr, ok := a.(*Bird); ok {
		fmt.Println("一只鸟", birdPtr.Speak())
		fmt.Println(birdPtr)
	} // 断言失败

	SortTest()

	animal.AnimalTest()


}
