package main

import "fmt"

func main() {
	arr := []int{3, 2, 4, 2, 6, 4, 4, 7, 8}
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if arr[j] > arr[j+1] { // 相邻元素两两对比
				temp := arr[j+1] // 元素交换
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr)
}
