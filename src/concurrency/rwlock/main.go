package main

import (
	"fmt"
	"sync"
)

func main() {
	var rwlock sync.RWMutex
	var num int
	var wg sync.WaitGroup

	wg.Add(2)
	// write go
	go func ()  {
		defer wg.Done()
		rwlock.Lock()
		defer rwlock.Unlock()
		num = 12
	}()

	// read go
	go func ()  {
		defer wg.Done()
		rwlock.RLock()	// 加读锁，读锁不会阻止读
		defer rwlock.RUnlock()
		fmt.Println("在读goroutine中读取num:", num)
	}()

	wg.Wait()

}