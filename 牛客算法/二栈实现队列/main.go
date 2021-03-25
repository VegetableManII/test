package main

import "fmt"

var stack1 []int
var stack2 []int

func main() {
	Push(1)
	Push(2)
	Push(3)
	fmt.Println(Pop())
	fmt.Println(Pop())
	Push(4)
	Push(5)
	fmt.Println(Pop())
	fmt.Println(Pop())
	fmt.Println(Pop())
}
func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		for i := len(stack1) - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = []int{}
	}
	a := stack2[len(stack2)-1]
	stack2 = stack2[0 : len(stack2)-1]
	return a
}
