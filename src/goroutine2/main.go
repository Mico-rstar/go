package main

import (
	"fmt"
	"sync"
	"time"
)

var chan1 = make(chan int)
var chan2 = make(chan int)
var doneChan = make(chan struct{})

func send(id int, num1 int, num2 int, wait *sync.WaitGroup) {
	defer wait.Done()
	fmt.Printf("Task%d start\n", id)
	time.Sleep(time.Second * 2)
	fmt.Printf("Task%d finish\n", id)
	chan1 <- num1
	chan2 <- num2
}

type Task struct {
	id    int
	data1 int
	data2 int
}

// 必须用defer调用
func Timing(now time.Time) {
	totol := time.Since(now)
	fmt.Println("total time = ", totol)
}

func main() {
	defer Timing(time.Now())
	var wait sync.WaitGroup

	tasks := []Task{
		{1, 3, 4},
		{2, 2, 3},
		{2, 3, 4},
	}
	wait.Add(len(tasks))
	for _, task := range tasks {
		go send(task.id, task.data1, task.data2, &wait)
	}

	var data1List []int
	var data2List []int

	go func() {
		defer close(chan1)
		defer close(chan2)
		defer close(doneChan)
		wait.Wait()
		doneChan<-struct{}{}
	}()

	func() {
		for {
			select {
			case data1 := <-chan1:
				data1List = append(data1List, data1)
			case data2 := <-chan2:
				data2List = append(data2List, data2)
			// 超时自动退出
			case <-time.After(3 * time.Second):
				fmt.Println("Time out")
				return
			case <-doneChan:
				return
			}
		}
	}()

	fmt.Println(data1List)
	fmt.Println(data2List)
}
