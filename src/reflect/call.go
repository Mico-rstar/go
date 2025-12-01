package main

import (
	"fmt"
	"reflect"
)

type Student struct {

}

func (s Student) CallMe(name string) {
	fmt.Println("Hello, world")
	fmt.Println("我是你", name, "giegie")
}

func ReflectCall(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumMethod(); i++ {
		if t.Method(i).Name != "CallMe" {
			continue
		}
		v.Method(i).Call([]reflect.Value{reflect.ValueOf("kunkun")})
	}

}