package main

import (
	"fmt"
	"log"
	"reflect"
)

func getType(obj any) {
	t := reflect.TypeOf(obj)
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("Int")
	case reflect.Chan:
		fmt.Println("Chan")
	case reflect.String:
		fmt.Println("String")
	}
}

func getValue(obj any) {
	t := reflect.ValueOf(obj)
	switch t.Kind() {
	case reflect.String:
		fmt.Println("String", t.String())
	case reflect.Int:
		fmt.Println("Int", t.Int())
	case reflect.Struct:
		fmt.Println("Struct")
	}

}

// 修改值
func modifyValue(old any, new any) {
	ov := reflect.ValueOf(old)
	nv := reflect.ValueOf(new)
	// 检查old是否为指针
	if ov.Kind() != reflect.Ptr {
		log.Println("传入修改参数非指针")
		return
	}
	// 检查是否为空指针
	if !ov.Elem().CanSet() {
		log.Println("非法指针")
		return
	}
	if ov.Elem().Kind() != nv.Kind() {
		log.Println("传入类型不统一")
		return
	}
	switch ov.Elem().Kind() {
	case reflect.String:
		ov.Elem().SetString(new.(string))
	case reflect.Int:
		ov.Elem().SetInt(nv.Int())
	}
}

func main() {
	var ss string
	fmt.Println(ss == "")
	getType(ss)
	ss = "好啦"
	getValue(ss)

	name := "zz"
	age := 11
	fmt.Printf("Before modify %v, %v\n", name, age)
	modifyValue(&name, "ll")
	modifyValue(&age, 12)
	fmt.Printf("Before modify %v, %v\n", name, age)

	// 测试空指针
	var p1 *string
	var p2 *int
	modifyValue(p1, "shis")
	modifyValue(p2, 1)

	user := User{
		Name1: "mimi",
		Name2: "xiaohai",
	}
	Upper(&user)
	fmt.Println(user)

	province := Province {
		number: 1,
		City1: "guangzhou",
		City2: "dongguan",
	}
	Upper(&province)
	fmt.Println(province)

	student := Student {}
	ReflectCall(&student)
}
