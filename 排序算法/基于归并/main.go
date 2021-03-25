package main

import "fmt"

func main() {
	arr := []int{56, 23, 2, 43, 56, 7}
	mergeSort(arr, 0, len(arr))
	fmt.Println(arr)
}

func merge(li []int, left, mid, right int) {
	i := left
	j := mid + 1
	tmp := []int{}
	for i <= mid && j <= right {
		if li[i] <= li[j] {
			tmp = append(tmp, li[i])
			i++
		} else {
			tmp = append(tmp, li[j])
			j++
		}
	}
	if i <= mid {
		tmp = append(tmp, li[i:mid+1]...)
	} else {
		tmp = append(tmp, li[j:right+1]...)
	}
	for k := 0; k < len(tmp); k++ {
		li[left+k] = tmp[k]
	}
}

func mergeSort(li []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(li, left, mid)
		mergeSort(li, mid+1, right)
		merge(li, left, mid, right)
	}
}
