package main

import (
	"fmt"
)

func f1(data ...interface{}) {
	fmt.Println(data...)
}

func f2() func() int {
	a, b := 0, 1
	return func() int {
		tmp := a + b
		a, b = b, tmp
		return tmp
	}
}

func main() {
	// var array []interface{} = []interface{}{0, 1, 2, 3, 4}
	// f1(array)
	// f1(array...)

	// f := f2()
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d ", f())
	// }
	select {}
}
