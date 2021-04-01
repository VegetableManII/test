package main

import (
	"study/newcoder/mylist"
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
 * @return bool布尔型
 */
func hasCycle(head *mylist.ListNode) bool {
	// write code here
	if head == nil {
		return false
	}
	// 快慢指针
	pfast, pslow := head, head
	for pfast != nil && pfast.Next != nil {
		pfast = pfast.Next.Next
		pslow = pslow.Next
		if pslow == pfast {
			return true
		}
	}
	return false
}
func main() {

}
