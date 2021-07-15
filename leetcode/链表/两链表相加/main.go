package main

import (
	. "study/newcoder/mylist"
)

func main() {
	l1 := CreateList([]int{4, 5})
	l2 := CreateList([]int{3, 5, 1})
	l1 = addTwoNumbers(l1, l2)
	PrintList(l1)
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var jinwei int = 0
	head := l1
	// 注意：尾节点特殊处理
	for l1.Next != nil && l2.Next != nil {
		tmp := l1.Val + l2.Val + jinwei
		if tmp >= 10 {
			l1.Val = tmp - 10
			jinwei = 1
		} else {
			l1.Val = tmp
			jinwei = 0
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	// 尾节点相加
	var tmpNode *ListNode
	tmp := l1.Val + l2.Val + jinwei
	if tmp >= 10 {
		l1.Val = tmp - 10
		jinwei = 1
	} else {
		l1.Val = tmp
		jinwei = 0
	}
	// 第一个数的尾节点，有两种情况：1.第一个数比第二个数短，2.两个数一样长
	if l1.Next == nil {
		// 两个数一样长
		if l2.Next == nil {
			if jinwei == 1 {
				l1.Next = &ListNode{Val: 1, Next: nil}
			}
			return head
		}
		l2 = l2.Next
		l1.Next = l2
		tmpNode = l2
	}
	// 第二个数的尾节点：第二个数比第一个数短
	if l2.Next == nil {
		tmpNode = l1.Next
	}
	// 对其中一个数剩下的位进行进位运算，注意尾节点的特殊处理
	for tmpNode.Next != nil {
		tmp := tmpNode.Val + jinwei
		if tmp >= 10 {
			tmpNode.Val = tmp - 10
			jinwei = 1
		} else {
			tmpNode.Val = tmp
			jinwei = 0
		}
		tmpNode = tmpNode.Next
	}
	// 尾节点特殊处理
	tmp = tmpNode.Val + jinwei
	if tmp >= 10 {
		tmpNode.Val = tmp - 10
		jinwei = 1
	} else {
		tmpNode.Val = tmp
		jinwei = 0
	}
	if jinwei == 1 {
		tmpNode.Next = &ListNode{Val: 1, Next: nil}
	}
	return head
}
