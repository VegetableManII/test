package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int, 0)
	order := make([]int, 0, 2)
	var a string
	for {
		fmt.Scanln(a)
		if _, ok := m[a]; ok {
			m[a]++
		} else {
			m[a] = 1
		}
	}
}
