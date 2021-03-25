package main

import "fmt"

func main() {
	fmt.Println(findKth([]int{1, 3, 5, 2, 2}, 5, 3))
}
func findKth(a []int, n int, K int) int {
	// write code here
	return quickSort(a, 0, n, K)
}

func quickSort(a []int, left, right, K int) int {
	i, j := left, right-1
	baseIndex := left + (right-left)/2
	a[i], a[baseIndex] = a[baseIndex], a[i]
	tmp := a[i]
	for i < j {
		for i < j && a[j] <= tmp {
			j--
		}
		a[i] = a[j]
		for i < j && a[i] >= tmp {
			i++
		}
		a[j] = a[i]
	}
	baseIndex = i
	a[baseIndex] = tmp
	if baseIndex+1 == K {
		return a[baseIndex]
	} else if baseIndex+1 < K {
		return quickSort(a, baseIndex+1, right, K)
	} else {
		return quickSort(a, left, baseIndex, K)
	}
}
