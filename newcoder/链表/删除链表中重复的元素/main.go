package main

import (
	. "study/newcoder/mylist"
)

func main() {
	l := CreateList([]int{1, 1, 1, 1, 1, 1, 1, 1})
	PrintList(deleteDuplication(l))
}
func deleteDuplication(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	if pHead.Next == nil {
		return pHead
	}
	pre := &ListNode{
		Next: pHead,
	} // pre 记录重复节点的前一个位置节点
	head, dupNode := pre, pHead.Next
	// dup记录连续重复的节点的后一个节点
	for dupNode != nil {
		if pre.Next.Val != dupNode.Val {
			pre = pre.Next
		} else {
			for dupNode != nil && pre.Next.Val == dupNode.Val {
				dupNode = dupNode.Next

			}
			pre.Next = dupNode
			if dupNode == nil {
				return head.Next
			}
		}
		dupNode = dupNode.Next
	}
	return head.Next
}
