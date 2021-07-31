package main

import (
	"fmt"
	"log"
	"sort"
)

var res [][]int

func main() {
	target := 8
	arr := []int{1, 10, 3, 4, 2, 5}
	res = make([][]int, 0, 1)
	sort.Ints(arr)
	// 有序数组 1,2,3,4,5,10
	// 回溯
	for _, v := range arr {
		tmp := []int{}
		trace(v, target, arr[1:], &tmp)
		res = append(res, tmp)
	}
	fmt.Println(res)
}

func trace(current int, target int, arr []int, res *[]int) {
	log.Println(current, target, arr, *res)
	if target < 0 {
		return
	}
	current = current + arr[0]
	find := target - current
	if exist(find, arr) {
		*res = append(*res, find)
	} else {
		trace(current, find, arr[1:], res)
		*res = append(*res, arr[0])
	}
}

func exist(a int, arr []int) bool {
	for _, v := range arr {
		if a == v {
			return true
		}
	}
	return false
}
