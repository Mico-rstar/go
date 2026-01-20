package main

import (
	"fmt"
	"sync"
)

// 线程不安全map
func ConcurrentUnsafeMap() {
	m := make(map[int]string)
	go func() {
		for {
			m[1] = "h"
		}
	}()
	go func() {
		for {
			fmt.Println(m[1])
		}
	}()
}

func ConcurrentSafeMap() {
	m := sync.Map{}
	go func() {
		for {
			m.Store(1, "a")
		}
	}()
	go func() {
		for {
			fmt.Println(m.Load(1))
		}
	}()
}

func main() {
	// ConcurrentUnsafeMap()
	ConcurrentSafeMap()
	select{}
}
