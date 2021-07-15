package main

import (
	. "study/leetcode/mylist"
)

func main() {
	l := CreateList([]int{1, 2, 3, 4, 5})
	PrintList(ReverseList(l))
}

/*
	解法一：位置逐个交换(头插法)
	解法二：将后一项指针指向前一项最后返回最后的尾节点
*/
// ReverseList 节点逐个交换到头部
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	head, tail, pnode := pHead, pHead, pHead
	if pnode.Next == nil {
		return pHead
	}
	pnode = pnode.Next

	for pnode != nil {
		tmp := pnode.Next
		tail.Next = tmp
		pnode.Next = head

		head = pnode
		pnode = tmp
	}
	return head
}
