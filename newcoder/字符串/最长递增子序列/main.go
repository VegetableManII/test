package main

import "fmt"

func main() {
	fmt.Println(LIS([]int{2, 1, 5, 3, 6, 4, 8, 9, 7}))
}

/**
 * retrun the longest increasing subsequence
 * @param arr int整型一维数组 the array
 * @return int整型一维数组
 */

// 贪心算法+二分查找+动态规划
func LIS(arr []int) []int {
	// write code here
	if len(arr) <= 1 {
		return arr
	}
	// 用来记录遍历过程中出现的最大值
	maxlength := 1
	// maxLen用来记录某一个位置的最长子序列的大小
	maxLen := make([]int, len(arr))
	// res用来记录最长子序列
	res := make([]int, 0, 32)
	for index, v := range arr {
		if len(res) == 0 {
			res = append(res, v)
			maxLen[0] = maxlength
			continue
		}
		if v >= res[len(res)-1] {
			res = append(res, v)
			maxlength = len(res)
			maxLen[index] = len(res)
		} else {
			// 如果 v 小于最后一个值则找到第一个大于 v 的值替换
			// 二分查找
			i, j := 0, len(res)-1
			for i < j {
				// 右移运算等于除以2的n次方，左移运算等于乘以2的n次方
				mid := i + (j-i)>>1
				if res[mid] > v {
					j = mid
				} else {
					i = mid + 1
				}
			}
			res[i] = v
			maxLen[index] = i + 1
		}
	}
	// 筛选最长子序列，maxLen中对应 i 则maxLen[i]的值从后往前依次递减取出即可得到
	// 最字典序小子序列
	/*
		实例数组：   [2 1 5 3 6 4 8 9 7 ]，需要 maxLen[] 记录以 arr[i] 为结尾的最长子序列的长度。
		本例 maxLen [1 1 2 2 3 3 4 5 4 ]
		又贪： 当 maxLen[i] == maxLen[i - 1] ，最小字典序要选 arr[i] ，跳过 i - 1。
		反证 ：本例 maxLen[4] maxLen[5]均为3 ，对应元素 6、4 。 如果换位成 4、6，那么 maxLen[5] 的将变为 4 ，
		矛盾。 （即 maxLen[i] == maxLen[i - 1] 推出 arr[i] <= arr[i - 1]）
	*/
	// 这里maxlength其实就是res的长度
	for i, j := len(maxLen)-1, maxlength; j > 0; i-- {
		if maxLen[i] == j {
			// maxLen里对应位置的长度和当前最大长度一致则替换res结果中的数据
			// 当前最大长度在依次替换之后会减1，倒序确定最长子序列
			j--
			res[j] = arr[i]
		}
	}
	return res

}
