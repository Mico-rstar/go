package main

import (
	"fmt"
	"time"
)

var signal chan struct{} = make(chan struct{})

func task(id int, tick int) {
	fmt.Printf("Task%d started\n", id)
	time.Sleep(time.Second * time.Duration(tick))
	fmt.Printf("Task%d finished\n", id)
	signal<- struct{}{}
}

func main() {
	go task(1, 2)
	go task(2, 2)
	go task(3, 2)



}