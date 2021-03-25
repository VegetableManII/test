package main

import "fmt"

func main() {
	fmt.Println(solve(7, 3))
}

func solve(M int, N int) string {
	// write code here
	// M == 0
	if M == 0 {
		return fmt.Sprintf("%d", 0)
	}
	// M < 0
	var stack []byte
	if M < 0 {
		stack = append(stack, '-')
		M = -M
	}
	for M >= N {
		stack = append(stack, '0'+byte(M%N))
		M /= N
	}
	stack = append(stack, '0'+byte(M%N))
	for i := 0; i < len(stack)/2; i++ {
		stack[i], stack[len(stack)-i-1] = stack[len(stack)-i-1], stack[i]
	}
	s := string(stack[:])
	return s
}
