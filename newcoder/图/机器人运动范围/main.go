package main

import "fmt"

func main() {
	fmt.Println(movingCount(11, 8, 16))
}
func movingCount(m int, n int, k int) int {
	j := 0
	if m >= k+1 && n >= k+1 {
		for i := 1; i <= k+1; i++ {
			j += i
		}
	} else if m < k+1 && n >= k+1 {
		for i := k - m + 2; i <= k+1; i++ {
			j += i
		}
	} else if n < k+1 && m >= k+1 {
		for i := k - n + 2; i <= k+1; i++ {
			j += i
		}
	} else {
		if m+n <= k {
			j = m * n
			return j
		}
		if m > n {
			tmp := n
			n = m
			m = tmp
		}
		start := k - n + 2
		edg := k - m + 2
		j += edg * m
		for i := n - edg; i > 0; i-- {
			j += start
			start += 1
		}
	}
	return j
}
