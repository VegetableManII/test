package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}

func maxArea(height []int) int {
	// 双指针
	// p = min(l,r)
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		tmp := min(height[l], height[r])
		if area := (r - l) * tmp; max < area {
			max = area
		}
		if height[r] == tmp {
			r--
		} else {
			l++
		}
	}
	return max
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
