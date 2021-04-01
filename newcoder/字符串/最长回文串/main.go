package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}
func longestPalindrome(s string) string {
	// 中心扩展法
	if s == "" {
		return s
	}
	start, end := 0, 0
	// 以每一个字符为中心向外扩展
	// 回文串中心字符可能是单个字符也可能是相等的双字符
	// 两种情况都计算出结果，取其中的最大值
	for i := range s {
		l1, r1 := expand(s, i, i)
		l2, r2 := expand(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}
func expand(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
