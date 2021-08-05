package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	time.Sleep(time.Millisecond * 800)
	t2 := time.Now().Sub(t1).Seconds()
	fmt.Println(t2)
	t3 := time.Now()
	time.Sleep(800 * time.Millisecond)
	t4 := time.Since(t3).Seconds()
	fmt.Println(t4)
	var arr1 []string
	for _ = range arr1 {
		fmt.Println("__________")
	}
	arr2 := reNil()
	for range arr2 {
		fmt.Println("__________")
	}
}

func reNil() []string {
	return nil
}
