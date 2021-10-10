package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)
	str := []byte{}
	count := 0
	for i := 0; i < N; i++ {
		var tmp byte
		fmt.Scanf("%c", &tmp)
		str = append(str, tmp)
		if len(str) < 2 {
			continue
		}
		if str[len(str)-2] == '0' && str[len(str)-1] == '1' {
			count++
			str = str[0 : len(str)-2]
		}
		if len(str) < 2 {
			continue
		}
		if str[len(str)-2] == '2' && str[len(str)-1] == '3' {
			count++
			str = str[0 : len(str)-2]
		}
	}
	fmt.Println(count * 2)
}
