package main

import (
	"math"
)

// func main() {
// 	a := []int{1}
// 	b := []int{2, 3, 4}
// 	log.Println(binarSearch(a, b))
// 	log.Println(binarSearch(b, a))
// }

func binarSearch(nums1, nums2 []int) float64 {
	// 特殊处理
	total := len(nums1) + len(nums2)
	if total == 0 {
		return 0
	} else if len(nums1) == 0 {
		if total%2 == 0 {
			return (float64(nums2[total/2]) + float64(nums2[(total/2)-1])) / 2
		} else {
			return float64(nums2[total/2])
		}
	} else if len(nums2) == 0 {
		if total%2 == 0 {
			return (float64(nums1[total/2]) + float64(nums1[(total/2)-1])) / 2
		} else {
			return float64(nums1[total/2])
		}
	}
	mid := total / 2
	i, idxa, idxb := 0, 0, 0
	cur := math.MinInt32
	post := 0
	curArray := nums1
	flag := true
	if nums1[0] > nums2[0] {
		curArray = nums2
		flag = false
	}

	for i < total {
		if flag {
			// 从a中查找
			if cur < curArray[idxa] {
				cur = curArray[idxa]
			}
			idxa++
			// 判断下一个的来源，注意另一个数组是否已经便利结束
			if idxa == len(nums1) || (idxb != len(nums2) && curArray[idxa] > nums2[idxb]) {
				flag = false
				post = nums2[idxb]
				curArray = nums2
			} else {
				post = nums1[idxa]
			}

		} else {
			// 从b中查找
			if cur < curArray[idxb] {
				cur = curArray[idxb]
			}
			idxb++
			// 判断下一个的来源，注意另一个数组是否已经便利结束
			if idxb == len(nums2) || (idxa != len(nums1) && curArray[idxb] > nums1[idxa]) {
				flag = true
				post = nums1[idxa]
				curArray = nums1
			} else {
				post = nums2[idxb]
			}

		}
		i++
		if i == mid {
			// 判断是否找到中位数
			if total%2 == 1 {
				// 奇数个直接取中间的数 mid+1
				return float64(post)
			} else {
				// 偶数个返回中间两个数的均值
				return (float64(post) + float64(cur)) / 2
			}
		}
	}
	return 0
}
