package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Score int
	Age int
}

type Students []Student

// 实现 sort.Interface接口
func (s Students) Less(i, j int) bool {
	return  s[i].Score > s[j].Score
}

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SortTest() {
	students := Students {
        {"张三", 85, 20},
        {"李四", 92, 19},
        {"王五", 78, 21},
        {"赵六", 92, 20},
	}
	fmt.Println("排序前")
	for _, s := range students {
		fmt.Printf("%s: %d分\n", s.Name, s.Score)
	}

	sort.Sort(students)

	fmt.Println("\n排序后")
	for _, s := range students {
		fmt.Printf("%s: %d分\n", s.Name, s.Score)
	}

}