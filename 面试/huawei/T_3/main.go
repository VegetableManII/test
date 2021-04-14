package main

import (
	"fmt"
	"sort"
)

func main() {
	var number int
	zhongzhuan := []int{}
	fmt.Scanln(&number)
	var tmp int
	for i := 0; i < number; i++ {
		fmt.Scanf("%d", &tmp)
		zhongzhuan = append(zhongzhuan, tmp)
	}
	fmt.Println(minRouting(zhongzhuan, number-1))

}
func minRouting(a []int, length int) int {
	sort.Ints(a)
	rout := 0
	ans := 0
	for i := len(a) - 1; i >= 0; i-- {
		ans += a[i]
		rout++
		if ans > length {
			return rout
		}
	}
	return rout
}
