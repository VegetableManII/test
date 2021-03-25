package main

import (
	"fmt"
	"math"
)

/*
shell排序
1.第一层循环设置分组间隔，分组间隔逐渐变小直到变为1
2.第二层循环依次取得各个分组中的距离gap的数进行比较
3.第三层循环是从第二个分组开始距离gap的数与第一个分组中的current比较，大于则向后移动
时间复杂度最好 O(n) 最坏 O(n²)
*/
func shellSort(arr []int) []int {
	len := len(arr)
	for gap := int(math.Floor(float64(len) / 2)); gap > 0; gap = int(math.Floor(float64(gap) / 2)) {
		for i := gap; i < len; i++ {
			j := i
			current := arr[i] // 用于记录当前遍历的坐标
			for j-gap >= 0 && current < arr[j-gap] {
				arr[j] = arr[j-gap] //数据向后移动
				j = j - gap         // 用于跨分组，因为第二层循环如果是从第二分组的第一个元素开始依次向后遍历
				// 所以当遍历到第三个分组时为能够与第一个分组的对应位置的数进行比较需要让 j - gap
			}
			arr[j] = current //找到最小的，放在第一个分组上
		}
	}
	return arr
}
func sS(arr []int) {
	len := len(arr)
	for step := len / 2; step >= 1; step /= 2 {
		// 插入排序

		for i := step; i < len; i += step {
			for j := i - step; j >= 0; j -= step {
				// 满足插入则交换元素
				if arr[j+step] < arr[j] {
					arr[j+step], arr[j] = arr[j], arr[j+step]
					continue
				}
				break
			}

		}
	}
}
func main() {
	arr := []int{3, 2, 1, 4, 5, 6, 3, 6, 7}
	sS(arr)
	fmt.Printf("%v\n", arr)

}
