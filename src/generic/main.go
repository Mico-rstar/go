package main

import (
	"fmt"
	"notes/src/generic/queue"
	"sort"
)

type Partial[T sort.Interface] struct {
	vl T
}

func (p Partial[T]) sort() {
	sort.Sort(p.vl)
}

func (p Partial[T]) Vl() T{
	return p.vl
}


/*
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}*/

type Student struct {
	name 	string
	score	int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return  s[i].score < s[j].score
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := Students {
        {"张三", 85},
        {"李四", 92},
        {"王五", 78},
        {"赵六", 92},
	}
	pt := Partial[Students] {stus}
	pt.sort()
	fmt.Println(pt)

	// 创建队列
	ql := queue.New(1, 2, 3)
	fmt.Println(ql)
	for !ql.Empty() {
		if fptr, err := ql.Front(); err == nil {
			fmt.Println(*fptr)
		} else {break}
		ql.Pop()
	}
	ql.Push(1)
	ql.Push(3)
	fmt.Println(ql)
}