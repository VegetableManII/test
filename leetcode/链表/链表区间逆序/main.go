package main

import (
	. "Algorithm-Ex/leetcode/mylist"
)

func reverseBetween(head *ListNode, m, n int) *ListNode {
	if head == nil || m > n {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < m-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	// 获得第 m 个节点
	cur := pre.Next
	// 头插法
	for i := 0; i < n-m; i++ {
		// 获得"头"节点的下一个节点用于插入
		tmp := pre.Next
		// 将“头”结点后置节点设置为 cur 指针的下一个节点
		// 成为新 “头” 节点
		pre.Next = cur.Next
		// 将当前节点的后置设置为其后第二个节点
		cur.Next = cur.Next.Next
		// 将“头”节点的后置节点的后置指针设置为之前保存的 老 “头” 节点
		// 逆转之后 cur 向后移动，cur.Next 为下一个要插入到头部的节点
		pre.Next.Next = tmp
	}
	return newHead.Next
}

func main() {
	t := CreateList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	reverseBetween(t, 2, 5)
}
