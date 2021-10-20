package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println(order(a, b))
}

func order(a, b string) string {
	res := make([]byte, 0, len(a))
	res = append(res, []byte(a)...)
	// 记录相同字符在b中的顺序
	sameInB := make([]byte, 0, len(b))
	for _, v := range b {
		if strings.Contains(a, string(v)) {
			sameInB = append(sameInB, byte(v))
		}
	}
	// 记录相同字符在a中出现的位置
	stra := make(map[byte]int, 1)
	// 记录相同字符在a中出现的顺序
	sameInA := make([]byte, 0, len(a))
	for idx, v := range a {
		if strings.Contains(b, string(v)) {
			stra[byte(v)] = idx
			sameInA = append(sameInA, byte(v))
		}
	}
	i := 0
	for _, v := range sameInA {
		// 替换
		res[stra[v]] = sameInB[i]
		i++
	}
	return string(res)
}
