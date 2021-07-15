package main

import "fmt"

// TreeNode is ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := reConstructBinaryTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	zhongxu(t)

}

func zhongxu(t *TreeNode) {
	if t == nil {
		return
	}
	zhongxu(t.Left)
	fmt.Println(t)
	zhongxu(t.Right)
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pre int整型一维数组
 * @param vin int整型一维数组
 * @return TreeNode类
 */
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here
	return rebuild(pre, vin)

}
func rebuild(pre []int, vin []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := pre[0]
	var rootIndex int // rootIndex 也能表示左枝或右枝有多少个孩子
	// 找到每一分支的根节点
	// 前序的第一个数字
	// 中序的中间位置的数字  其左侧是左枝右侧是右枝
	for index, v := range vin {
		if v == root {
			rootIndex = index
			break
		}
	}
	tree := &TreeNode{root, nil, nil}
	// 前序 下一个数字即当前递归节点的左节点
	// 前序 下一个数字到 中序 的rootIndex+1的位置即当前递归节点左枝所拥有的节点数量
	tree.Left = rebuild(pre[1:rootIndex+1], vin[0:rootIndex])
	tree.Right = rebuild(pre[rootIndex+1:], vin[rootIndex+1:])
	return tree

}
