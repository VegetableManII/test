package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 4, 1, 4, 5, 1}
	quickSort(arr, 0, len(arr))
	fmt.Println(arr)
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right-1
	base := (right-left)>>1 + left
	tmp := arr[base]
	arr[i], arr[base] = arr[base], arr[i]

	for i < j {
		for i < j && arr[j] >= tmp {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] <= tmp {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = tmp
	quickSort(arr, left, i)
	quickSort(arr, i+1, right)
}
