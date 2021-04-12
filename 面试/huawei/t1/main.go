package main

import (
	"fmt"
	"math"
)

func main() {
	stuN, opM := 0, 0
	fmt.Scanf("%d %d", &stuN, &opM)
	grade := make([]int, 0, stuN)
	for i := 0; i < stuN; i++ {
		var tmp int
		fmt.Scanf("%d", &tmp)
		grade = append(grade, tmp)
	}

	var A, B int
	var op byte
	for i := 0; i < opM; i++ {
		fmt.Scanf("%c %d %d\n", &op, &A, &B)
		if op == 'Q' {
			fmt.Println(retMax(grade, A-1, B))
		}
		if op == 'U' {
			grade[A-1] = B
		}
	}
}
func retMax(arr []int, l, r int) int {
	res := math.MinInt32
	for i := l; i < r; i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
