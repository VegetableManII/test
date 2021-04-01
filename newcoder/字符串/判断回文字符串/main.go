package main

import "fmt"

func main() {
	fmt.Println(judge2("a"))
}

// 双指针最佳方案
func judge2(str string) bool {
	phead, pend := 0, len(str)-1
	for phead < pend {
		if str[phead] != str[pend] {
			return false
		}
		phead++
		pend--
	}
	return true
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串 待判断的字符串
 * @return bool布尔型
 */
func judge(str string) bool {
	// write code here
	stack := make([]byte, 0, len(str))
	var i int
	for ; i < len(str)/2; i++ {
		stack = append(stack, str[i])
	}
	if len(str)%2 == 0 {
		for ; i < len(str); i++ {
			if stack[len(stack)-1] == str[i] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	} else {
		i = i + 1
		for ; i < len(str); i++ {
			if stack[len(stack)-1] == str[i] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return true
}
