package main

import "fmt"

func main() {
	fmt.Println(fibonacci(20))
}
func fibonacci(n int) []int {
	pre0, pre1 := 0, 1
	tmp := 1
	res := []int{}
	for i := 0; i < n; i++ {
		res = append(res, tmp)
		tmp = pre0 + pre1
		pre0 = pre1
		pre1 = tmp
	}
	return res
}
