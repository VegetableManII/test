package main

import "sort"

/*
思路：
首先对原数组进行排序，求三数之和为0
固定中位数，分别从数组两边进行查找，向中位数聚集
从 index = 1 开始一次将每一个数字作为中位数去查找
*/
// 解法一 最优解，双指针 + 排序
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	// 固定 index 查找 nums[start] + nums[end] = -nums[index] 的结果
	for index = 1; index < length-1; index++ {
		// 双指针从两端开始移动
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			// 重复的数字直接跳过，左指针向index移动
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			// 重复的数字跳过，右指针向index移动
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
