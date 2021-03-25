package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarysearch(arr, 3))
	ar := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(ar, 0))
}

/*
二分查找中关于right的赋值
right = n-1 => while(left <= right) => right = middle-1;
right = n => while(left < right) => right = middle;
*/
/**
 *
 * @param A int整型一维数组
 * @param target int整型
 * @return int整型
 */
func binarysearch(A []int, target int) int {
	// write code here
	right := len(A)
	left := 0
	for left < right {
		mid := left + (right-left)>>1
		if A[mid] < target {
			left = mid
		} else if A[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return -1
}

/*
以上是对顺序的数组进行二分查找
对经过旋转之后的数组进行二分查找
如 1 2 3 10 11 12 13 14 15 经过一次旋转 10 11 12 13 14 15 1 2 3
   1 2 3 4 5 6 7 8 10 11 15 经过旋转 10 11 15 1 2 3 4 5 6 7 8
旋转后的数组首先要找到旋转的分界，判断的依据即是这一部分的有序数组其首元素大于有序数组的后面第一个元素
*/
func search(A []int, target int) int {
	right := len(A)
	left := 0
	for right > left {
		mid := left + (right-left)>>1
		if A[mid] == target {
			return mid
		}
		if A[left] <= A[mid] {
			// 说明 mid 左侧有序
			if A[mid] > target && A[left] <= target {
				// 如果待查找的数在左侧 [left,mid) 内
				right = mid
			} else {
				left = mid
			}
		} else {
			// 说明 mid 右侧有序
			if A[mid] < target && A[right-1] >= target {
				// 如果待查找的数在右侧侧 (mid,right] 内
				left = mid
			} else {
				right = mid
			}
		}
	}
	return -1
}
