package tree

import "fmt"

// TreeNode 定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arrlen int

// CreateTree 先序遍历创建二叉树
func build(n []int) *TreeNode {
	if n[0] == -1 {
		arrlen--
		return nil
	}
	tmp := &TreeNode{
		Val:   n[0],
		Left:  nil,
		Right: nil,
	}
	arrlen--
	if arrlen > 0 {
		tmp.Left = build(n[len(n)-arrlen:])
		tmp.Right = build(n[len(n)-arrlen:])
	}
	return tmp
}
func CreateTree(n []int) *TreeNode {
	arrlen = len(n)
	return build(n)
}

// PrintTree 先序遍历
func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PrintTree(root.Left)
	PrintTree(root.Right)
}
