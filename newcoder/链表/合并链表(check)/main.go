package main

import (
	"study/newcoder/mylist"
)

func main() {
	l1 := mylist.CreateList([]int{5})
	l2 := mylist.CreateList([]int{1, 2, 4})
	mylist.PrintList(mergeTwoLists(l1, l2))
}

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param l1 ListNode类
 * @param l2 ListNode类
 * @return ListNode类
 */
func mergeTwoLists(l1 *mylist.ListNode, l2 *mylist.ListNode) *mylist.ListNode {
	// write code here
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var psmall, pbig, heade *mylist.ListNode
	if l1.Val >= l2.Val {
		psmall, pbig = l2, l1
		heade = psmall
	} else {
		psmall, pbig = l1, l2
		heade = psmall
	}
	// psmall 作为主链，pbig插入到psmall里面
	for psmall.Next != nil && pbig != nil {
		if psmall.Val <= pbig.Val && psmall.Next.Val >= pbig.Val {
			tmp := pbig.Next

			pbig.Next = psmall.Next
			psmall.Next = pbig

			pbig = tmp
		}
		psmall = psmall.Next
	}
	if psmall.Next == nil {
		psmall.Next = pbig
	}

	return heade
}
