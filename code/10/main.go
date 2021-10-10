package main

import (
	"log"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param grid int整型二维数组
 * @return int整型
 */
func ncov_defect(grid [][]int) int {
	// write code here
	count := 0
	for row, v := range grid {
		for column, ncov := range v {
			if ncov == 1 {
				// 扩展
				if column+1 < len(v) && (v[column+1] == 1) {
					continue
				}
				if row+1 < len(grid) && grid[row+1][column] == 1 {
					continue
				}
				// 安插守卫
				if column-1 > 0 && v[column-1] == 0 {
					count++
				}
				if column+1 < len(v) && v[column+1] == 0 {
					count++
				}
				if row+1 < len(grid) && grid[row+1][column] == 0 {
					count++
				}
				if row-1 > 0 && grid[row-1][column] == 0 {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	// grid := [][]int{
	// 	{},
	// }
	// fmt.Println(ncov_defect(grid))
	a := []string{"a", "b", "c"}
	e := arr(a)
	a = append(a, "e")
	a = append(a, "f")
	log.Println("函数内 ", e)
	log.Println("函数外 ", a)

}

func arr(s []string) []string {
	return append(s, "d")
}

func test(p interface{}) {
	if p == nil {
		log.Println("p is nil")
		return
	}
	log.Println("p is not nil")
}
