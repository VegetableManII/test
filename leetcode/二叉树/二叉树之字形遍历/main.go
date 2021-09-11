package main

import (
	"fmt"

	. "Algorithm-Ex/leetcode/tree"
)

func main() {
	troot := CreateTree([]int{1, 2, 7, 8, -1, -1, 9, -1, -1, 4, 7, -1, -1, 2, -1, -1, 3, 5, -1, -1, 6, -1, -1}, 0)
	res := zigzagLevelOrder(troot)
	fmt.Println(res)
}

/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	// write code here
	// 层序遍历 从左到右和从右到左交替
	flag := false

	res := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	var prequeueAndStack []*TreeNode
	for len(queue) > 0 {
		arr := []int{}
		prequeueAndStack = queue
		queue = make([]*TreeNode, 0)
		for _, v := range prequeueAndStack {
			if v != nil {
				arr = append(arr, v.Val)
			}
		}
		// flag = false 该轮为从左到右,下一轮需要从右到左
		for len(prequeueAndStack) > 0 {
			v := prequeueAndStack[len(prequeueAndStack)-1]
			prequeueAndStack = prequeueAndStack[0 : len(prequeueAndStack)-1]
			if !flag {
				if v.Right != nil {
					queue = append(queue, v.Right)
				}
				if v.Left != nil {
					queue = append(queue, v.Left)
				}
			} else {
				// flag = true 该轮为从右到左，下一轮为从左到右
				// 当前队列应该按栈的形式先出尾部元素(即下一次从左到右的顺序的父节点)
				if v.Left != nil {
					queue = append(queue, v.Left)
				}
				if v.Right != nil {
					queue = append(queue, v.Right)
				}
			}
		}
		flag = !flag
		res = append(res, arr)
	}

	return res
}
