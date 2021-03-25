package main

func binSort(li []int, binnum int) {
	minnum, maxnum := li[0], li[0]
	for i := 0; i < len(li); i++ {
		if minnum > li[i] {
			minnum = li[i]
		}
		if maxnum < li[i] {
			maxnum = li[i]
		}
	}
	bin := make([][]int, binnum)
	for j := 0; j < len(li); j++ {
		n := (li[j] - minnum) / ((maxnum - minnum + 1) / binnum)
		bin[n] = append(bin[n], li[j])
		k := len(bin[n]) - 2
		for k >= 0 && li[j] < bin[n][k] {
			bin[n][k+1] = bin[n][k]
			k--
		}
		bin[n][k+1] = li[j]
	}
	o := 0
	for p, q := range bin {
		for t := 0; t < len(q); t++ {
			li[o] = bin[p][t]
			o++
		}
	}
}
