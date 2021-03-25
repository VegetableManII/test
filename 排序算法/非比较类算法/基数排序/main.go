package main

import (
	"math"
	"strconv"
)

func radixSort(li []int) {
	maxnum := li[0]
	for i := 0; i < len(li); i++ {
		if maxnum < li[i] {
			maxnum = li[i]
		}
	}
	for j := 0; j < len(strconv.Itoa(maxnum)); j++ {
		bin := make([][]int, 10)
		for k := 0; k < len(li); k++ {
			n := li[k] / int(math.Pow(10, float64(j))) % 10
			bin[n] = append(bin[n], li[k])
		}
		m := 0
		for p := 0; p < len(bin); p++ {
			for q := 0; q < len(bin[p]); q++ {
				li[m] = bin[p][q]
				m++
			}
		}
	}
}
