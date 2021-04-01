package main

import (
	"fmt"

	"study/newcoder/tree"
)

// 递归函数 切片要用取址
func main() {
	num := []int{1, 2, 3, 0, 0, 3, 0, 0, 2, 0, 0}
	root := tree.CreateTree(num)
	// tree.PrintTree(root)
	fmt.Println(isSymmetric(root))
}

/**
 *
 * @param root TreeNode类
 * @return bool布尔型
 */
func isSymmetric(root *tree.TreeNode) bool {
	// write code here
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}
func compare(l *tree.TreeNode, r *tree.TreeNode) bool {
	if l != nil && r != nil {
		if l.Val == r.Val {
			// 比较左右子树的外侧节点
			o := compare(l.Left, r.Right)
			i := compare(l.Right, r.Left)
			return o && i
		}
	}
	if l == nil && r == nil {
		return true
	}
	return false
}
