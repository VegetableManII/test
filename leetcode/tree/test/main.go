package main

import (
	. "Algorithm-Ex/leetcode/tree"
)

func main() {
	t := CreateTree([]int{2, 3, 4, 3, -1, 5, -1}, 1)
	PrintTree(t)
}
