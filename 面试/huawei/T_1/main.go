package main

import "fmt"

func main() {
	var str string
	fmt.Scanln(&str)
	ans := reverse([]byte(str))
	fmt.Printf("%s", ans)

}
func reverse(str []byte) []byte {
	stack := []int{}
	ans := []byte{}
	for i := range str {
		if str[i] == ')' {
			// 翻转
			r := i - 1
			for l := stack[len(stack)-1] + 1; l <= r; l++ {
				str[l], str[r] = str[r], str[l]
				r--
			}
			stack = stack[:len(stack)-1]
		}
		if str[i] == '(' {
			stack = append(stack, i)
		}
	}
	for i := range str {
		if str[i] == '(' || str[i] == ')' {
			continue
		}
		ans = append(ans, str[i])
	}
	return ans
}
