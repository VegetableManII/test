package main

import (
	"fmt"
	"math"

	. "Algorithm-Ex/leetcode/tree"
)

func main() {
	t := CreateTree([]int{90, 69, 49, -1, 52, -1, -1, 89, -1, -1, -1}, 0)
	res := minDiffInBST(t)
	fmt.Println(res)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minDiffInBST(root *TreeNode) int {
	// 中序遍历
	ans := math.MaxInt32
	pre := -1
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != -1 {
			if tmp := root.Val - pre; tmp < ans {
				ans = tmp
			}
		}
		pre = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return ans
}
