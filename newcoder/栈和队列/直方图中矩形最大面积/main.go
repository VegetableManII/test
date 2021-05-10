package main

import "fmt"

/*
单调栈：
单调栈依次保存直方图的高度的下标，保证单调递增，如果当前高度比栈顶元素小
则需要在当前 index 位置停留向前计算矩形面积
面积：取出当前栈顶保存的下标对应的高度值作为矩形高度，
	 矩形的宽度为，当前 index 与出栈元素的距离
向前依次计算，直到遇到 当前 index 的高度大于栈顶元素的高度，当前 index 入栈
*/

func largestRectangleArea(heights []int) int {
	maxArea, stack, height := 0, []int{}, 0
	for i := 0; i <= len(heights); i++ {
		if i == len(heights) {
			height = 0
		} else {
			height = heights[i]
		}
		// 高度单调递增，
		// 优先判断栈是否为空，如果栈为空直接入栈
		// 左边的判断语句优先于右边的判断语句这里不会出现越界
		if len(stack) == 0 || height >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			// 非单调时，特殊处理
			// 取出栈顶元素，其对应的高度即为矩形的高
			//
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 计算宽度
			length := 0
			//
			if len(stack) == 0 {
				length = i
			} else {
				length = i - tmp
				// length = i - 1 - stack[len(stack)-1]
			}
			maxArea = max(maxArea, heights[tmp]*length)
			// 停留在当前位置继续向前判断
			i--
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(largestRectangleArea([]int{5, 4, 4, 1, 3}))
}
