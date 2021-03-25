package main

func countSort(li []int) {
	maxnum := li[0]
	for i := 1; i < len(li); i++ {
		if maxnum < li[i] {
			maxnum = li[i]
		}
	}
	arr := make([]int, maxnum+1)
	for j := 0; j < len(li); j++ {
		arr[li[j]]++
	}
	k := 0
	for m, n := range arr {
		for p := 0; p < n; p++ {
			li[k] = m
			k++
		}
	}
}
