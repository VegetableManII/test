package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)

	pankick := make([]int, 0, n)
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		pankick = append(pankick, tmp)
	}

	sort.Ints(pankick)
	count := 1
	pankickHeap := 0
	idx := n - 1
	// 排序后 依次遍历数组
	// 只有遇到连续重复的数字重复次数超过当前堆数量时需要增加新的堆
	for idx >= 1 {
		if pankick[idx] == pankick[idx-1] {
			count++
			if count > pankickHeap {
				pankickHeap++
			}
		} else {
			count = 1
		}
		idx--
	}
	fmt.Println(pankickHeap)
}
