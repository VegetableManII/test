package main

import (
	"fmt"

	"Algorithm-Ex/leetcode/mylist"
)

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类
 * @param k int整型
 * @return ListNode类
 */
func reverseKGroup(head *mylist.ListNode, k int) *mylist.ListNode {
	if !checkLen(head, k) {
		return head
	}
	var r *mylist.ListNode
	flag := false
	/*
		g 上一组的尾节点
		h 慢指针 正序的头结点，逆转序列的尾节点
		t 快指针 逆转序列的尾节点的下一个节点，即将逆转的节点
	*/
	if head.Next == nil {
		return head
	}
	g := &mylist.ListNode{
		Next: head,
	}
	h, t := head, head.Next

	for t != nil {
		for i := 0; i < k-1; i++ {
			tmp := g.Next
			g.Next = t
			h.Next = t.Next
			t.Next = tmp

			t = h.Next
		}
		if !flag {
			r = g.Next
			flag = true
		}
		if checkLen(t, k) {
			g = h
			h = h.Next
			t = t.Next
		} else {
			break
		}
	}
	return r
}

func checkLen(cur *mylist.ListNode, k int) bool {
	n := cur
	for i := 0; i < k; i++ {
		if n == nil {
			return false
		}
		n = n.Next
	}
	return true
}
func main() {

	list := mylist.CreateList([]int{1})

	fmt.Println("翻转前：")
	mylist.PrintList(list)

	list = reverseKGroupFast(list, 1)
	fmt.Println("翻转后：")
	mylist.PrintList(list)
}

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类
 * @param k int整型
 * @return ListNode类
 */

func reverseKGroupFast(head *mylist.ListNode, k int) *mylist.ListNode {
	if head == nil {
		return nil
	}

	root := &mylist.ListNode{
		Next: head,
	}

	e := root.Next
	h := root.Next
	beforee := root
	afterh := h.Next

	for h != nil {
		// 先让h移动到一组的最后一个节点
		// 如果不够k个元素直接返回
		for count := 0; count < k-1; count++ {
			h = h.Next
			if h == nil {
				return root.Next
			}
		}
		// 记录下一组的开始节点
		afterh = h.Next
		// 逆转函数
		reverseK(e, h)
		beforee.Next = h // 将设置的根节点root 连接到逆转之后的分组上
		e.Next = afterh  // 将上一组的末尾连接到下一组的起始

		beforee = e
		e = beforee.Next
		h = e
	}

	return root.Next
}

func reverseK(h *mylist.ListNode, e *mylist.ListNode) {
	root := &mylist.ListNode{
		Next: h,
	}

	pe := root
	p := pe.Next
	ph := p.Next

	for p != e {
		p.Next = pe
		pe = p
		p = ph
		ph = p.Next
	}

	p.Next = pe

	return
}
