package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, q int
	fmt.Scanf("%d %d", &n, &q)

	var tmp int
	s := make([]int, 0, n)
	qes := make([]int, 0, q)
	var totle int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		s = append(s, tmp)
		totle += tmp
	}
	for i := 0; i < q; i++ {
		fmt.Scanln(&tmp)
		qes = append(qes, tmp)
	}
	// index 为大于平均值的橘子
	answer := make([]string, q)
	find(s, qes, totle, &answer)
	for i := range answer {
		fmt.Printf("%s\n", answer[i])
	}
}

func find(s, qes []int, totle int, answer *[]string) {
	if len(qes) == 0 {
		return
	}

	average := totle / len(s)
	sort.Ints(s)
	index := 0
	for s[index] <= average {
		index++
	}
	lightOrange := func() int {
		t := 0
		for i := 0; i < index; i++ {
			t += s[i]
		}
		return t
	}()
	heavyOrange := func() int {
		t := 0
		for i := index; i < len(s); i++ {
			t += s[i]
		}
		return t
	}()

	for _, v := range qes {
		if v > totle {
			*answer = append(*answer, "NO")
			continue
		}
		if v == totle {
			*answer = append(*answer, "YES")
			continue
		}
		// 如果大于均值
		if v > lightOrange {
			if v == heavyOrange {
				*answer = append(*answer, "YES")
				continue
			}
			// 重新筛选==> 计算均值然后查找 恰好相等的橘子
			find(s[index:], qes, heavyOrange, answer)
		} else {
			if v == lightOrange {
				*answer = append(*answer, "YES")
				continue
			}
			// 重新筛选==> 计算均值然后查找 恰好相等的橘子
			find(s[:index], qes, heavyOrange, answer)
		}
		*answer = append(*answer, "NO")
	}
	return
}
