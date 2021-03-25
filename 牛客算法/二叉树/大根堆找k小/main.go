package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := GetLeastNumbers_Solution([]int{4, 5, 1, 6, 2, 7, 3, 8}, 3)
	sort.Ints(arr)
	fmt.Println(arr)
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */
func GetLeastNumbers_Solution(input []int, k int) []int {
	// write code here
	if input == nil || len(input) == 0 || k > len(input) {
		return nil
	}
	// len(input)/2 - 1 是为了找到数组不越界的情况下最后面的节点的父节点
	for i := len(input)/2 - 1; i >= 0; i-- {
		// 建立大根堆
		adjustHeap(input, i, k)
	}
	for i := k; i < len(input); i++ {
		//将后面的数依次和K个数的最大值比较
		if input[0] > input[i] {
			input[i], input[0] = input[0], input[i]
			adjustHeap(input, 0, k)
		}
	}
	maxHeap := make([]int, 0, k)
	maxHeap = append(maxHeap, input[0:k]...)
	return maxHeap
}

/*
查找数组中某个数的父结点和左右孩子结点，比如已知索引为i的数，那么

1.父结点索引：(i-1)/2（这里计算机中的除以2，省略掉小数）
2.左孩子索引：2*i+1
3.右孩子索引：2*i+2
*/
func adjustHeap(arr []int, index, n int) {
	left := 2*index + 1
	right := 2*index + 2
	bigest := index
	if left < n && arr[left] > arr[bigest] {
		bigest = left // 更新最大值索引
	}
	if right < n && arr[right] > arr[bigest] {
		bigest = right // 更新最大值的索引
	}
	// 最大值不是父节点则进行交换
	if bigest != index {
		arr[index], arr[bigest] = arr[bigest], arr[index]
		// 当前节点和他的叶子节点大根堆已排好，接着排叶子节点
		adjustHeap(arr, bigest, n)
	}
}
