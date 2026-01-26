package main

import (
	"fmt"
	"sync"
	"time"
)

var nameChan = make(chan string)

func goShopping(name string, wait *sync.WaitGroup) {
	defer wait.Done()
	fmt.Println(name, " start")
	time.Sleep(1 * time.Second)
	fmt.Println(name, " finish shopping")
	nameChan <- name
}

func main() {
	var wait sync.WaitGroup
	sl := []string{"malao", "mingji", "tangren"}
	wait.Add(len(sl))
	for _, si := range sl {
		go goShopping(si, &wait)
	}

	// 创建守护协程
	go func() {
		defer close(nameChan)
		wait.Wait()
	}()

	// for {
	// 	if name, ok := <- nameChan; ok {
	// 		fmt.Printf("收到%s的消息\n", name)
	// 	} else {break}
	// }
	for name := range nameChan {
		fmt.Printf("收到%s的消息\n", name)
	}

	fmt.Printf("执行了%d个工作协程，全部执行结束", len(sl))

}
