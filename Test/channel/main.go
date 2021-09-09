package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	wg.Add(1)
	go func(c <-chan int, wg *sync.WaitGroup) {
		for {
			tmp, flag := <-c
			fmt.Println(tmp, flag)
		}
		wg.Done()
	}(ch, wg)
	for i := 0; i < 3; i++ {
		ch <- i
	}
	time.Sleep(5 * time.Second)
	close(ch)
	ch <- 4
	wg.Wait()
}
