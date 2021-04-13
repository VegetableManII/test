package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	stuN, opM := 0, 0
	var A, B int
	var op byte
	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()
	for len(input) > 0 {
		in := strings.Split(string(input), " ")
		stuN, _ = strconv.Atoi(in[0])
		opM, _ = strconv.Atoi(in[1])
		grade := make([]int, 0, stuN)
		for i := 0; i < stuN; i++ {
			var tmp int
			fmt.Scanf("%d", &tmp)
			grade = append(grade, tmp)
		}
		fmt.Println()
		for i := 0; i < opM; i++ {
			fmt.Scanf("%c %d %d\n", &op, &A, &B)
			if op == 'Q' {
				if A <= B && A <= stuN {
					fmt.Println(retMax(grade, A-1, B))
				}
			}
			if op == 'U' {
				if A <= stuN {
					grade[A-1] = B
				}
			}
		}
		input, _, _ = reader.ReadLine()
		fmt.Println(len(input))
	}

}
func retMax(arr []int, l, r int) int {
	res := math.MinInt32
	for i := l; i < r; i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
