package main

import "fmt"

var stack1 []int
var stack2 []int

type node struct {
	Val  int
	next *node
}

func main() {
	// fmt.Println(fibonacci(5))
	f := fibonacciBibao()
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	}

	// ok := find(4, [][]int{{1, 2, 3, 4}, {2, 3, 5, 6}, {3, 4, 5}})
	// fmt.Println(ok)
}

// 青蛙跳台阶
//     | = 1 ，x=1
// fn= | = 2 ，x=2
//     | = fn-1 + fn-2 ，x>2
func jumpFloor(n int) int {
	a, b := 1, 1
	for i := 1; i < n; i++ {
		a = a + b
		b = a - b
	}
	return a
}

// 青蛙一次可以跳1阶、2阶、3阶......n阶
// fn = 2*fn-1	f1 = 1
// 问题转化为求指数
// 矩形覆盖问题
// 		| = 1 ，x=1
// fn = | = 2 ，x=2
// 		| = fn-1 + fn-2 ，x>2

//  统计一个整数的二进制表示中1的个数
func numberOf1(n int) int {
	count := 0
	for int32(n) != 0 {
		count++
		n &= int((int32(n) - 1))
	}
	return count
}

// 斐波那契数列
func fibonacci(n int) int {
	// write code here
	if n == 0 || n == 1 {
		return n
	}
	a, b, result := 0, 1, 1
	for i := n - 2; i > 0; i-- {
		a = b
		b = result
		result = a + b
	}
	return result
}

// 闭包实现
func fibonacciBibao() func() int {
	var cur int = 0
	var post int = 1
	return func() int {
		tmp := cur
		cur = post
		post = tmp + cur
		tmp = post
		return tmp
	}
}

// 栈实现队列
func push(node int) {
	stack1 = append(stack1, node)
}
func pop() int {
	if len(stack2) == 0 {
		stack2 = append(stack2, stack1...)
		stack1 = []int{}
	}
	rs := stack2[0]
	stack2 = stack2[1:]
	return rs
}

// 判断二维数组中的的某个数
func find(target int, array [][]int) bool {
	// write code here
	l := len(array)
	if l == 0 {
		return false
	}
	for i := range array {
		if array[i][l-1] < target {
			continue
		}
		if array[i][0] > target {
			continue
		}
		for v := range array[i] {
			if array[i][v] == target {
				return true
			}
		}
	}
	return false
}
