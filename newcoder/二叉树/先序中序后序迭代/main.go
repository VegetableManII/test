package main

import (
	"fmt"

	. "study/newcoder/tree"
)

func main() {
	troot := CreateTree([]int{1, 2, 7, 8, -1, -1, 9, -1, -1, 4, -1, -1, 3, 5, -1, -1, 6, -1, -1}, 0)
	inorder(troot)
}

// 先序遍历迭代
func preorder(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	cur := root
	/*

		算法一：根节点入栈
		采用栈的结构，把当前节点压入栈遍历该节点的左子树
		如果有左子树则压栈，如果左子树为空则从栈中取出当前节点
		算法二：右子树入栈
		当前节点直接使用然后其右子树入栈接着遍历左子树
		算法三：右子树入栈然后左子树入栈
		先让跟节点入栈然后栈非空取出根节点
		之后循环右子树入栈-->左子树入栈最后判断栈非空
		节点出栈
	*/
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			fmt.Println(cur.Val) // 取出数据
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// cur = stack[0]
			// stack = stack[1:]
			cur = cur.Right
		}
	}
}

// 中序遍历迭代
func inorder(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	cur := root
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println(cur.Val)
			cur = cur.Right
		}
	}
}

// 后序遍历
func postorder(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	cur := root
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			fmt.Println(cur.Val)
			cur = cur.Right
		} else {
			// stack.pop()
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = cur.Left
		}
	}
}
