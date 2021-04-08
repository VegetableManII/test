package main

type BinarySearchTree struct {
	Root *BinarySearchTreeNode
}
type BinarySearchTreeNode struct {
	Value  int64
	Times  int64
	Left   *BinarySearchTreeNode
	Right  *BinarySearchTreeNode
	Parent *BinarySearchTreeNode
}

func PrintBinarySearchTree(tree *BinarySearchTree) []int {
	res := make([]int, 0, 4)
	getBSTreeNode(tree.Root, &res)
	return res
}
func getBSTreeNode(node *BinarySearchTreeNode, res *[]int) {
	if node == nil {
		return
	}
	getBSTreeNode(node.Left, res)
	*res = append(*res, int(node.Value))
	getBSTreeNode(node.Right, res)
}
func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}
func (tree *BinarySearchTree) Add(value int64) {
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}
	tree.Root.Add(value)
}

// 添加元素
func (node *BinarySearchTreeNode) Add(value int64) {
	if value < node.Value {
		// 插入左子树
		// 如果左子树为空直接添加
		if node.Left == nil {
			node.Left = &BinarySearchTreeNode{Value: value}
			node.Left.Parent = node
		} else {
			// 否则递归添加
			node.Left.Add(value)
		}
	} else if value > node.Value {
		// 如果插入的值比节点的值大，则插入到右子树
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode{Value: value}
			node.Right.Parent = node
		} else {
			node.Right.Add(value)
		}
	} else {
		node.Times += 1
	}
}

// 删除元素
/*
删除元素有四种情况：

第一种情况，删除的是根节点，且根节点没有儿子，直接删除即可。
第二种情况，删除的节点有父亲节点，但没有子树，也就是删除的是叶子节点，直接删除即可。
第三种情况，删除的节点下有两个子树，因为右子树的值都比左子树大，那么用右子树中的最小元素来替换删除的节点，这时二叉查找树的性质又满足了。右子树的最小元素，只要一直往右子树的左边一直找一直找就可以找到。
第四种情况，删除的节点只有一个子树，那么该子树直接替换被删除的节点即可。
*/
func (tree *BinarySearchTree) Del(value int64) {
	if tree.Root == nil {
		return
	}
	tree.Root.Del(value)
}
func (node *BinarySearchTreeNode) Del(value int64) {
	if node == nil {
		return
	}
	if value < node.Value {
		node.Left.Del(value)
	} else if value > node.Value {
		node.Right.Del(value)
	} else {
		// 第一种情况 和 第二种情况 直接删除
		if node.Left == nil && node.Right == nil {
			// 删除父节点对它的引用
			if node.Parent.Left == node {
				node.Parent.Left = nil
			}
			if node.Parent.Right == node {
				node.Parent.Right = nil
			}
			node = nil
			return
		}
		// 第四种情况
		// 只有左子树，直接用子树替换
		if node.Right == nil {
			if node.Parent.Left == node {
				node.Parent.Left = node.Left
			}
			if node.Parent.Right == node {
				node.Parent.Right = node.Left
			}
			node = nil
		} else if node.Left == nil {
			if node.Parent.Left == node {
				node.Parent.Left = node.Right
			}
			if node.Parent.Right == node {
				node.Parent.Right = node.Right
			}
			node = nil
		} else {
			// 第三种情况左右子树都不为空
			// 当前节点的后继节点为右子树中最小的那个
			rightMin := node.Right.FindMinValue()
			node.Value = rightMin.Value
			node.Times = rightMin.Times
			// 判断自己是父节点的左枝还是右枝
			if rightMin.Parent.Left == rightMin {
				rightMin.Parent.Left = nil
			}
			if rightMin.Parent.Right == rightMin {
				rightMin.Parent.Right = nil
			}
			rightMin = nil
		}
	}
}

// 查找最大值或最小值
func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.FindMinValue()
}
func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	// 如果左子树为空证明当前节点为最左侧节点
	if node.Left == nil {
		return node
	}
	return node.Left.FindMinValue()
}

// 找出最大值的节点
func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	return tree.Root.FindMaxValue()
}
func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	// 右子树为空，表面已经是最右的节点了，该值就是最大值
	if node.Right == nil {
		return node
	}
	// 一直右子树递归
	return node.Right.FindMaxValue()
}

// 查找指定元素
func (tree *BinarySearchTree) Find(value int64) *BinarySearchTreeNode {
	if tree.Root == nil {
		// 如果是空树，返回空
		return nil
	}
	/*
			// 查找指定元素的父节点
			// 如果根节点等于该值，根节点其没有父节点，返回nil
		    if tree.Root.Value == value {
		        return nil
		    }
	*/
	return tree.Root.Find(value)
}
func (node *BinarySearchTreeNode) Find(value int64) *BinarySearchTreeNode {

	/*
		// 外层没有值相等的判定，因为在内层已经判定完毕后返回父亲节点。
	*/
	if value == node.Value {
		// 如果是当前节点则返回该节点
		return node
	} else if value < node.Value {

		/*
					// 左子树的根节点的值刚好等于该值，那么父亲就是现在的node，返回
			        if node.Left.Value == value {
			            return node
			        } else {
			            return node.Left.FindParent(value)
			        }
		*/

		// 如果小于当前节点的值则查找下一个
		if node.Left == nil {
			return nil
		}
		return node.Left.Find(value)
	} else {

		/*
					// 右子树的根节点的值刚好等于该值，那么父亲就是现在的node，返回
			        if node.Right.Value == value {
			            return node
			        } else {
			            return node.Right.FindParent(value)
			        }
		*/

		// 如果大于当前节点则查找右子树
		if node.Right == nil {
			return nil
		}
		return node.Right.Find(value)
	}
}
