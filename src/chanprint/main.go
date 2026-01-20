package main

import (
	"fmt"
	"sync"
)

func printNums(start <-chan int, done chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 1
	for {
		startSignal := <-start
		if startSignal == 0 {
			break
		}
		fmt.Printf("%d%d", i, i+1)
		i += 2
		done <- struct{}{}
	}
}

// start
// 1 represent continue
// 0 represent end
func printChars(chars []string, start chan<- int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(chars); i += 2 {
		start <- 1
		<-done
		fmt.Print(chars[i])
		if i+1 < len(chars) {
			fmt.Print(chars[i+1])
		}
	}
	start <- 1
	<-done
	start <- 0
}

func main() {
	var chars []string = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var wg sync.WaitGroup
	start := make(chan int, 1)
	done := make(chan struct{})
	wg.Add(2)
	go printNums(start, done, &wg)
	go printChars(chars, start, done, &wg)
	wg.Wait()

}
