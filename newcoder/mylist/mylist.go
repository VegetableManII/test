package mylist

import (
	"fmt"
)

// ListNode 用于做链表题
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateList 创建链表
func CreateList(array []int) *ListNode {
	list := &ListNode{array[0], nil}
	p := list
	array = array[1:]
	for i := range array {
		tmp := &ListNode{array[i], nil}
		p.Next = tmp
		p = tmp
	}
	return list
}

// PrintList 打印链表
func PrintList(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

// 找到链表中的环的入口
/*
快指针以2步的速度前进，慢指针以1步的速度前进
第一次相遇说明有环
快指针回到起点以1步的速度前进
第二次相遇找到入口
*/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// 删除链表中的重复节点
func deleteDuplication(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}
	var head *ListNode = &ListNode{0, pHead}
	pre := head
	cur := pHead.Next
	var step int = 0
	for cur != nil {
		if pre.Next.Val == cur.Val {
			step++
			cur = cur.Next
			if cur == nil {
				pre.Next = cur
			}
			continue
		}
		if step > 0 {
			pre.Next = cur
			cur = cur.Next
			step = 0
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}

	return head.Next
}

// 翻转链表
func revers(head *ListNode) *ListNode {
	cur, tmp := head, head
	var pre *ListNode = nil

	for cur != nil {
		/*
			1. 保存下一个节点
			2. 当前节点指向更改为指向前一个节点
			3. 双指针向后移动
		*/
		tmp = cur.Next
		cur.Next = pre
		// pre和cur同时向后移动
		pre = cur
		cur = tmp
	}
	return pre
}
