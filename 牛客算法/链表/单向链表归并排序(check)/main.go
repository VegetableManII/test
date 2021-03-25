package main

import (
	. "study/牛客算法/mylist"
)

func main() {
	tr := CreateList([]int{3, 5, 2, 6, 1})
	PrintList(tr)
	tr = sortInListByMerge(tr)
	PrintList(tr)
}
func merge(p1, p2 *ListNode) *ListNode {
	if p1.Next == nil && p2.Next == nil {
		if p1.Val > p2.Val {
			p2, p1 = p1, p2
		}
		p1.Next = p2
		return p1
	} else {
		head := &ListNode{Next: nil}
		p3 := head
		for p1 != nil && p2 != nil {
			if p1.Val < p2.Val {
				p3.Next = p1
				p1 = p1.Next
			} else {
				p3.Next = p2
				p2 = p2.Next
			}
			p3 = p3.Next
		}
		if p1 == nil {
			p3.Next = p2
		} else {
			p3.Next = p1
		}
		return head.Next
	}
}
func sortInListByMerge(head *ListNode) *ListNode {
	// 先找到中间节点
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	left := head
	right := slow.Next
	slow.Next = nil
	left = sortInListByMerge(left)   // 更新分组头
	right = sortInListByMerge(right) // 更新分组头
	return merge(left, right)

}

/**
 *
 * @param head ListNode类 the head node
 * @return ListNode类
 */
func sortInList(head *ListNode) *ListNode {
	// write code here
	// 冒泡
	if head == nil || head.Next == nil {
		return head
	}
	pre := &ListNode{Next: head}
	p1, p2, p3, p4 := pre, pre, head, head.Next // 位置1
	for p1 != nil {
		for p4 != nil {
			if p4.Val < p3.Val {
				p3.Next = p4.Next
				p4.Next = p3
				p2.Next = p4
				p3, p4 = p4, p3
			}
			p3, p4 = p3.Next, p4.Next
			p2 = p2.Next
		}
		PrintList(pre.Next)
		p2, p3, p4 = pre, pre.Next, pre.Next.Next // head头部已经改变不能按照位置1的方式赋值
		p1 = p1.Next
	}
	return pre.Next
}
