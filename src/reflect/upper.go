package main

import (
	"log"
	"reflect"
	"strings"

)

type User struct {
	Name1 string `upper:"-"`
	Name2 string
}

type Province struct {
	number	int		`upper:"-"`
	City1	string	`upper:"-"`
	City2	string
}


// 修改结构体中的string字段为大写
func Upper(obj any) {
	v := reflect.ValueOf(obj)
	// 检查非指针
	if v.Kind() != reflect.Ptr {
		log.Default().Println("传入obj非指针")
		return
	}
	// 检查是否为合法指针
	if !v.Elem().CanSet() {
		log.Default().Println("传入非法指针")
	}
	v = reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		upper := t.Field(i).Tag.Get("upper")
		if field.Kind() != reflect.String {
			log.Default().Println("在非string字段上标注upper无效")
			continue
		}
		if upper == "" {
			continue
		}
		field.SetString(strings.ToUpper(field.String()))
	}
}


