package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Student struct {
	name string
	id   string
	age		int 
}

// select name, id from student where ...
func Find(obj any, query ...any) (sql string, err error) {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		return "", errors.New("not a struct")
	}
	if len(query) == 0 {
		return "", errors.New("no query")
	}
	fields := make([]string, 0, 5)
	// 将结构体字段添加到fields
	for i := 0; i < t.NumField(); i++ {
		itemName := t.Field(i).Name
		fields = append(fields, itemName)
	}

	queryStr, ok := query[0].(string)
	if !ok {
		return "", errors.New("lack of query str")
	}
	queryNums := strings.Count(queryStr, "?")
	if queryNums != len(query)-1 {
		return "", errors.New("numbers of ? not match query")
	}
	for _, item := range query[1:] {
		// 针对不同类型需要以不同格式插入
		switch item := item.(type) {
		case string:
			queryStr = strings.Replace(queryStr, "?", fmt.Sprintf(`'%s'`, item), 1)
		case int:
			queryStr = strings.Replace(queryStr, "?", strconv.Itoa(item), 1)
		}
	}
	sql = fmt.Sprintf("select %s from %s where %s", strings.Join(fields, ","), strings.ToLower(t.Name()), queryStr)
	return sql, nil
}
