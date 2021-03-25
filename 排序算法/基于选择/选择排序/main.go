package main

import "fmt"

/*
选择排序
时间复杂度 O(n²) */
func main() {
	arr := []int{3, 2, 4, 2, 6, 4, 4, 7, 8}
	len := len(arr)
	var minIndex, temp int
	for i := 0; i < len-1; i++ {
		minIndex = i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] { // 寻找最小的数
				minIndex = j // 将最小数的索引保存
			}
		}
		temp = arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}
	fmt.Println(arr)
}
