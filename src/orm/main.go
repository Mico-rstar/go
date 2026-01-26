package main

import (
	"fmt"
	"log"
)

func main() {
	stu := Student{
		name: "张三",
		id:   "202330331841",
		age: 18,
	}
	sql, err := Find(stu, "id = ? and name = ? and age = ?", "张三", "1342532525", 18)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sql)
}