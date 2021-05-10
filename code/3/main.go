package main

import "fmt"

type g struct {
	least int
	money int
}

func main() {
	goods := make(map[string]*g, 0)
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	var s string
	var w, c int
	for i := 0; i < n; i++ {
		fmt.Scanln(&s, &w, &c)
		goods[s] = &g{
			least: c,
			money: w,
		}
	}

	var t string
	var d int
	var yingli int
	for i := 0; i < m; i++ {
		fmt.Scanln(&t, &d)
		if goods[t].least >= d {
			yingli += d * goods[t].money
		} else {
			fmt.Println(-1 * (i + 1))
			return
		}
	}
	fmt.Println(yingli)
}
