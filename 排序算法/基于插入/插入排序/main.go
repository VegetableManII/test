package main

import "fmt"

/*
时间复杂度最好 O(n) 最坏 O(n²)
*/
func insertionSort(arr []int) []int {
	len := len(arr)
	var preIndex, current int
	for i := 1; i < len; i++ {
		preIndex = i - 1
		current = arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
	return arr
}
func main() {
	arr := []int{3, 2, 4, 2, 6, 4, 4, 7, 8}
	fmt.Println(shell(arr))
}

// shell排序 不稳定  无法保证一样大的两个数的相对位置
func shell(arr []int) []int {
	for gap := len(arr) / 2; gap > 0; gap = gap / 2 {
		fmt.Printf("_________gap = %d__________\n", gap)
		for j := gap; j < len(arr); j++ {
			fmt.Printf("________number %d__________", arr[j])
			current := arr[j]
			i := j
			for i-gap >= 0 && current < arr[i-gap] {
				arr[i] = arr[i-gap]
				i = i - gap
			}
			arr[i] = current
			fmt.Printf("%v\n", arr)
		}
	}
	return arr
}
