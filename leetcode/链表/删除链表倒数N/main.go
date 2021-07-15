package main

import (
	. "study/newcoder/mylist"
)

func main() {
	l := CreateList([]int{1, 2, 3, 4, 5})
	FindKthToTail(l, 5)
	PrintList(l)
}

// FindKthToTail is...
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	f, s := pHead, pHead
	for i := 0; i < k; i++ {
		f = f.Next
	}
	// 如果f，移动到末尾说明要删除的元素是第一个节点
	if f == nil {
		return pHead.Next
	}
	for f.Next != nil {
		f = f.Next
		s = s.Next
	}
	tmp := s.Next
	s.Next = tmp.Next
	return pHead
}
