package main

import "fmt"

func main() {
	res := Inc()
	fmt.Println(res)
}
func Inc() (v int) {
	defer func() { v++ }()
	return 42
	// return 执行了三个操作
	// 1. 将v的值赋值42
	// 2. 执行defer
	// 3. 执行RET指令
}
func function() {
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
}
