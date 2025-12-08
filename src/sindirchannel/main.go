package main

import "fmt"


func producer(out chan<- int) {
	for i:=0; i<10; i++ {
		out<- i*i
	}
	close(out)
}

func main() {
	// 直接定义单向channel
	var chw chan<- string // 只能写入string
	var chr <-chan string // 只能读取string

	chw = make(chan<- string)
	chr = make(<-chan string)

	fmt.Println(chw)
	fmt.Println(chr)

	// 先定义，后赋值为单向chan
	c := make(chan int, 3)
	var send chan<- int = c
	var read <-chan int = c
	producer(c)
	fmt.Println(send)
	fmt.Println(read)
}