package main

import "fmt"

type tree struct {
	val    int
	pLeft  *tree
	pRight *tree
}

var index int

func build(root *tree, numbers []int) *tree {
	if index >= len(numbers) {
		return nil
	}
	if numbers[index] == -1 {
		index++
		return nil
	}
	// 中序 构建树
	root.val = numbers[index]
	index++
	root.pLeft = build(&tree{}, numbers)
	root.pRight = build(&tree{}, numbers)
	return root
}
func main() {
	treeNodes := []int{1, 2, -1, -1, 3, 4, -1, -1, 5, 6, -1, -1, -1}
	t := &tree{}
	build(t, treeNodes)
	deepth := treeDeep(t)
	fmt.Println(deepth)
}
func treeDeep(root *tree) int {
	nodeQueue := make([]*tree, 0, 16)
	nodeQueue = append(nodeQueue, root)
	deepth := 0
	for len(nodeQueue) != 0 {
		tmpQueue := nodeQueue
		deepth++
		for len(tmpQueue) != 0 {
			node := tmpQueue[0]
			tmpQueue = tmpQueue[1:]
			nodeQueue = nodeQueue[1:]
			if node.pLeft != nil {
				nodeQueue = append(nodeQueue, node.pLeft)
			}
			if node.pRight != nil {
				nodeQueue = append(nodeQueue, node.pRight)
			}
		}
	}
	return deepth
}
