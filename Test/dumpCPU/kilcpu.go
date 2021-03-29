package main

import (
	"sync"
)

func f1() {
	var i = 0
	for {
		i++
	}
}

var wg sync.WaitGroup

func main() {
	//runtime.GOMAXPROCS(4)

	wg.Add(1)
	for i := 0; i < 10; i++ {
		go f1()
	}
	wg.Wait()
}
