package main

import "fmt"

// 自定义一个容器类型，例如列表
type IntList struct {
    items []int
}

// All 是一个迭代器函数，符合 range 的要求
func (l *IntList) All() func(func(int) bool) {
    return func(yield func(int) bool) {
        for _, v := range l.items {
            // yield 返回 false 表示中断迭代
            if !yield(v) {
                return
            }
        }
    }
}

func main() {
    list := &IntList{items: []int{10, 20, 30}}

    // 使用 for/range 迭代自定义类型
    for v := range list.All() {
        fmt.Println(v)
    }
}
