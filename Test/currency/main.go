package main

import (
	"fmt"
	"sync"
)

/*
通过range channel利用channel的机制实现多个goroutine之间的竞争
*/
func worker(ports chan int, number int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(number, " ", p)
		wg.Done()
	}
}
func main() {
	ports := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, i+1, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
