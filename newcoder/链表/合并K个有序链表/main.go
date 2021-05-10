package main

import . "study/newcoder/mylist"

/*
分治，链表两两合并
*/

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists1(left, right)
}

/*
递归到两个链表的尾部，对于某节点：l1 < l2，则归并 l1.Next 和 l2，同理，l2
——如果 l1==nil，则需要把 l2 剩下的节点拼接到 l1 上，返回 l2
——同理 l2 返回 l1
回溯合并两个链表，从 l1 或 l2 的尾部开始，当前节点及其后面的已归并链表作为返回值进行链表合并
——
*/
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists1(l1, l2.Next)
	return l2
}
