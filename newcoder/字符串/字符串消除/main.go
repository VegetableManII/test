package main

import "fmt"

func main() {
	// var n int
	// fmt.Scanln(&n)
	// number := make([]byte, n, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scanf("%c", &number[i])
	// }
	number := []byte{2, 1, 3, 7, 9, 2}
	length := len(number)
	for i := 0; i < len(number)-1; i++ {
		if number[i]+number[i+1] == 10 {
			length = checkTen(number, i, i+1, len(number))
			break
		}
	}
	fmt.Println(length)

}
func checkTen(number []byte, a, b, l int) int {
	if l == 0 || l == 1 {
		return l
	}
	if number[a]+number[b] == 10 {
		if a == 0 {
			return checkTen(number, b+1, b+2, l-2)
		} else if b == len(number)-1 {
			return checkTen(number, a-2, a-1, l-2)
		}
		return checkTen(number, a-1, b+1, l-2)
	}
	return a + len(number) - b + 1
}
