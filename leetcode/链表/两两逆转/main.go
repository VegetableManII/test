package main

import (
	. "study/leetcode/mylist"
)

func main() {
	l := CreateList([]int{})
	l = swapPairs(l)
	PrintList(l)
}

// 递归实现
func swapPairsTraceBack(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairsTraceBack(head.Next)
	newHead.Next = head
	return newHead

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ppre, p, pnext := &ListNode{}, head, head.Next
	for pnext != nil {
		p.Next = pnext.Next

		tmp := pnext
		pnext = pnext.Next

		if ppre.Next == nil {
			head = tmp
		}
		ppre.Next = tmp
		tmp.Next = p

		ppre = tmp.Next
		p = p.Next
		if pnext != nil {
			pnext = pnext.Next
		} else {
			break
		}
	}
	return head
}
