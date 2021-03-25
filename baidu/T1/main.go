package main

import (
	"fmt"
)

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	sum -= arr[0]
	return sum
}
func main() {
	group := 0

	n, _ := fmt.Scan(&group)

	if n == 0 {
		fmt.Println("no data")
	}
	foods := make([][]int, 0, group)
	for i := 0; i < group; i++ {
		num := 0
		hungry := 0
		fmt.Scan(&num, &hungry)
		food := make([]int, 0, num)
		foods = append(foods, food)
		foods[i] = append(foods[i], hungry)
		for j := 0; j < num; j++ {
			f := 0
			fmt.Scan(&f)
			foods[i] = append(foods[i], f)
		}
	}

	for _, v := range foods {
		res, err := find(v[1:], v[0])
		switch err {
		case -1:
			fmt.Println(-1)
		case 0:
			for i := 0; i < len(v)-1; i++ {
				fmt.Printf("%d ", i+1)
			}
			fmt.Println()
		case 1:
			fmt.Println(len(res))
			for i := 0; i < len(res); i++ {
				fmt.Printf("%d ", res[i])
			}
			fmt.Println()
		}

	}

}

func find(arr []int, target int) ([]int, int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum < target {
		return nil, -1
	}
	if sum == target {
		return nil, 0
	}
	// 找出能够达到target的最小和
	res := make([]int, 0)
	foodMap := make(map[int]int, len(arr))
	total := 0
	for i, v := range arr {
		total += v
		foodMap[v] = i + 1
		if total > target {
			tmp := total
			for k := range foodMap {
				tmp -= k
				if tmp > target {
					delete(foodMap, k)
				} else {
					break
				}
			}
		}
	}
	for _, v := range foodMap {
		res = append(res, v)
	}
	return res, 1
}
