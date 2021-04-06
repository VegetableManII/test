package main

import "fmt"

func main() {
	res := generateParenthesis(2)
	fmt.Println(res)
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}
	res := make([][]string, 2, n)
	res[0] = []string{""}
	res[1] = []string{"()"}
	// 动态规划
	for i := 2; i < n+1; i++ {
		l := make([]string, 0, 0)
		for j := 0; j < i; j++ {
			p := res[j]
			q := res[i-j-1]
			for k1 := range p {
				for k2 := range q {
					l = append(l, string("("+p[k1]+")"+q[k2]))
				}
			}
		}
		res = append(res, l)
	}
	return res[n]
}
