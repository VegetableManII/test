package main

import "fmt"

/**
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func findKth(a []int, n int, K int) int {
	// write code here
	quick(a, 0, n)
	return a[K]

}
func main() {
	arr := []int{1, 3, 5, 2, 2}
	fmt.Println(findKth(arr, 5, 3))
}
func quick(a []int, l, r int) {
	if l >= r {
		return
	}
	// 基准点
	mid := (r-l)>>1 + l
	i, j := l, r-1
	tmp := a[mid]
	for i < j {
		for a[j] < a[mid] && i < j {
			j--
		}
		a[mid] = a[j]
		for a[i] > a[mid] && i < j {
			i++
		}
		a[j] = a[i]
	}
	a[i] = tmp
	mid = i
	quick(a[l:mid], l, mid)
	quick(a[mid+1:r], mid+1, r)
}
