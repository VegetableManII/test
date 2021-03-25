package main

import (
	"fmt"
)

type node struct {
	val   int
	pnext *node
}

func main() {
	list := deleN(createList([]int{1, 2, 3, 4, 5}), 2)
	printList(list)
}
func createList(arr []int) *node {
	h := &node{
		val:   arr[0],
		pnext: nil,
	}
	plist := h
	for i := 1; i < len(arr); i++ {
		tmp := &node{arr[i], nil}
		plist.pnext = tmp
		plist = plist.pnext
	}
	return h
}
func printList(l *node) {
	for l != nil {
		fmt.Println(l.val)
		l = l.pnext
	}
}
func deleN(l *node, n int) *node {
	head := l
	fast, slow := l, l
	// 判断n手否比链表大
	for i := 0; i < n+1; i++ {
		if fast == nil {
			return head
		}
		fast = fast.pnext
	}
	for fast != nil {
		fast = fast.pnext
		slow = slow.pnext
	}
	tmp := slow.pnext.pnext
	slow.pnext = tmp
	return head
}
