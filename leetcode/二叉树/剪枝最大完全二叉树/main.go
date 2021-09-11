package main

import (
	"fmt"
	"sort"

	. "Algorithm-Ex/leetcode/tree"
)

func main() {
	// tr := CreateTree([]int{1, 2, 4, -1, -1, 5, -1, -1})
	// solve(tr)
	// PrintTree(tr)
	a := []int{10, 3, 4, 2, 1}
	sort.Ints(a)
	fmt.Println(a)
}
func solve(root *TreeNode) *TreeNode {
	// write code here
	if root == nil || root.Left == nil {
		return root
	}
	depth := 1
	// 层序遍历
	prelevel := []*TreeNode{}
	queue := make([]*TreeNode, 0, 16)
	queue = append(queue, root)
	prelevel = queue
	for len(queue) != 0 {
		midQueue := queue
		for len(midQueue) != 0 {
			tmp := midQueue[0]
			midQueue = midQueue[1:]
			queue = queue[1:]
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
		if len(queue) < (1 << depth) {
			// 剪枝
			for i := range prelevel {
				prelevel[i].Left = nil
				prelevel[i].Right = nil
			}
		} else {
			prelevel = queue
			depth++
		}
	}
	return root
}
