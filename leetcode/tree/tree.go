package tree

import "fmt"

// TreeNode 定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arrlen int

func buildByInOrder(n []int) *TreeNode {
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
		tmp.Left = buildByInOrder(n[len(n)-arrlen:])
		tmp.Right = buildByInOrder(n[len(n)-arrlen:])
	}
	return tmp
}
func buildByBFS(n []int) *TreeNode {
	queue := make([]*TreeNode, 0, 1)
	root := &TreeNode{
		Val:   n[0],
		Left:  nil,
		Right: nil,
	}
	arrlen--
	queue = append(queue, root)
	for arrlen > 0 {
		preQueue := queue
		for len(preQueue) > 0 {
			tmp := preQueue[0]
			if arrlen > 0 && n[len(n)-arrlen] != -1 {
				node := &TreeNode{}
				node.Val = n[len(n)-arrlen]
				arrlen--
				tmp.Left = node
				queue = append(queue, node)
			} else {
				arrlen--
			}
			if arrlen > 0 && n[len(n)-arrlen] != -1 {
				node := &TreeNode{}
				node.Val = n[len(n)-arrlen]
				arrlen--
				tmp.Right = node
				queue = append(queue, node)
			} else {
				arrlen--
			}
			preQueue = preQueue[1:]
			queue = queue[1:]
		}
	}
	return root
}

// CreateTree 0-先序创建 1-层序遍历
func CreateTree(n []int, flag int) *TreeNode {
	arrlen = len(n)
	switch flag {
	case 0:
		return buildByInOrder(n)
	case 1:
		return buildByBFS(n)
	}
	return nil
}

// PrintTree 先序遍历
func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val)
	PrintTree(root.Left)
	PrintTree(root.Right)
}
