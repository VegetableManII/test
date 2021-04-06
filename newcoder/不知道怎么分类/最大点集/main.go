package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(str)

	points := make([][2]int, n)
	for i := 0; i < n; i++ {
		str, _ := reader.ReadString('\n')
		strs := strings.Split(str, " ")
		points[i][0], _ = strconv.Atoi(strs[0])
		points[i][1], _ = strconv.Atoi(strs[1])
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

}
