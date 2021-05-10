package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	s := make([]int, 0, n)
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		s = append(s, tmp)
	}

	resL, resR := 0, 0

	r, l := 0, 0
	bigger, littler := 0, 0
	// 遍历数组
	for i, v := range s {
		if v > k {
			// 遇到大于 k 的添加到结果集中
			// 窗口右边界移动
			bigger++
			r = i
		} else {
			// 如果遇到小于 k 的统计其个数
			// 如果小于 k 的个数没有超过 大于 k 的个数则添加
			littler++
			if littler > bigger {
				// 如果小于 k 的个数 比大于的还多
				// 记录当前窗口值
				if r-l > resR-resL {
					resL, resR = l, r
				}
				// 左侧窗口移动 要求移动后要成为完美数列
				if s[l] > k {
					bigger--
				} else {
					littler--
				}
				l++
			} else {
				// 添加
				r = i
			}
		}
	}
	if r-l > resR-resL {
		resL, resR = l, r
	}
	fmt.Println(resR - resL + 1)
}
