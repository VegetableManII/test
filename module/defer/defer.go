package main

import "fmt"

func main() {
	res := deferFc2()
	fmt.Println(res)
}
func deferFc0() (v int) {
	defer func() { v++ }()
	return 42
	// return 执行了三个操作
	// 1. 将v的值赋值42
	// 2. 执行defer
	// 3. 执行RET指令
}
func deferFc1() {
	for i := 0; i < 3; i++ {
		defer func() { println(i) }()
	}
}

var g = 100

func deferFc2() (r int) {
	defer func() {
		g = 200
	}()

	fmt.Printf("f: g = %d\n", g)

	return g
}
