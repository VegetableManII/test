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
	fmt.Println(insertionSort(arr))
}
func iS(arr []int) []int {
	var current, preIndex int
	for i := 1; i < len(arr); i++ {
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
