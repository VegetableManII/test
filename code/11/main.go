package main

import (
	"fmt"
)

func main() {
	var N int
	radius := make(map[float32]struct{}, N)
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		var tmp float32
		fmt.Scanf("%f", &tmp)
		if _, ok := radius[tmp]; ok {
			continue
		}
		radius[tmp] = struct{}{}
	}
	fmt.Println(len(radius))
}
