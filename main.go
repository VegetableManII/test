package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	table := [][]string{
		{".", ".", "9", "7", "4", "8", ".", ".", "."},
		{"7", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", "2", ".", "1", ".", "9", ".", ".", "."},
		{".", ".", "7", ".", ".", ".", "2", "4", "."},
		{".", "6", "4", ".", "1", ".", "5", "9", "."},
		{".", "9", "8", ".", ".", ".", "3", ".", "."},
		{".", ".", ".", "8", ".", "3", ".", "2", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "6"},
		{".", ".", ".", "2", "7", "5", "9", ".", "."},
	}
	solveSudoku(table)
	formatOutput(table)

}

var hashmap1 [9]map[string]struct{} // 每一行
var hashmap2 [9]map[string]struct{} // 每一列
var hashmap3 [9]map[string]struct{} // 每一块
var x int
var y int

// debug
var bug bool

func pointToIndex(x, y int) int {
	if y == 0 {
		switch x {
		case 0:
			return 0
		case 1:
			return 1
		case 2:
			return 2
		}
	} else if y == 1 {
		switch x {
		case 0:
			return 3
		case 1:
			return 4
		case 2:
			return 5
		}
	} else if y == 2 {
		switch x {
		case 0:
			return 6
		case 1:
			return 7
		case 2:
			return 8
		}
	}
	return 0
}

func initalize(board [][]string) {
	for row, value := range board {
		for column, v := range value {
			if hashmap1[row] == nil {
				hashmap1[row] = make(map[string]struct{}, 9)
			}
			hashmap1[row][v] = struct{}{}
			if hashmap2[column] == nil {
				hashmap2[column] = make(map[string]struct{}, 9)
			}
			hashmap2[column][v] = struct{}{}
			kuai_row := row / 3
			kuai_column := column / 3
			if hashmap3[pointToIndex(kuai_column, kuai_row)] == nil {
				hashmap3[pointToIndex(kuai_column, kuai_row)] = make(map[string]struct{}, 9)
			}
			hashmap3[pointToIndex(kuai_column, kuai_row)][v] = struct{}{}
		}
	}
}

func solveSudoku(board [][]string) {
	// 初始化
	initalize(board)

	formatBlock()

	if ok := recursive(board); ok {
		log.Println("Yes")
		return
	}
	log.Println("No")
}

/*
 递归以每一行为起始，每一次递归都从可填的最小数字开始进行递归填充，并判断是够满足
 1、每一行不重复
 2、每一列不重复
 3、每3×3不重复
 hashmap 保存每一行/列/3×3是否存在
*/
func recursive(board [][]string) bool {
	// 从每一行开始
	// 当全部填充完成，goNext会越界
	if y == 9 {
		return true
	}
	if board[y][x] != "." {
		// 不填充
		curX, curY := x, y
		goNext()

		if ok := recursive(board); !ok {
			x, y = curX, curY
			return false
		}

		return true
	}
	curX, curY := x, y
	// 填充
	for i := 1; i < 10; i++ {
		// 判断行不存在
		if _, ok := hashmap1[curY][strconv.Itoa(i)]; !ok {
			// 判断列不存在
			if _, ok := hashmap2[curX][strconv.Itoa(i)]; !ok {
				// 判断块不存在
				kuai_column := curX / 3
				kuai_row := curY / 3

				if _, ok := hashmap3[pointToIndex(kuai_column, kuai_row)][strconv.Itoa(i)]; !ok {
					// 满足条件，尝试填充
					fill(strconv.Itoa(i), curX, curY, board)

					formatOutput(board)

					goNext()
					if ok := recursive(board); !ok {
						// 有冲突，则清除填充内容，继续循环用 i 尝试填充
						wipe(strconv.Itoa(i), curX, curY, board)
						x, y = curX, curY

						formatOutput(board)
						// debug
						// debug
						if curY == 0 && board[0][2] == "04" && board[0][3] == "06" && board[0][5] == "08" && board[0][6] == "09" {
							bug = true
						}
						if bug {
							fmt.Scanln()
						}
					} else {
						return true
					}
				} else {
					fmt.Printf("current: %v, num: %v, block: %v\n", i, pointToIndex(kuai_column, kuai_row), hashmap3[pointToIndex(kuai_column, kuai_row)])
				}
			} else {
				fmt.Printf("current: %v, num: %v, column: %v\n", i, curX, hashmap2[curX])
			}
		} else {
			fmt.Printf("current: %v, num: %v, row: %v\n", i, curY, board[curY])
		}
	}
	return false
}

/*
 ascii 码
 0 —— 48
 1 —— 49
 ...
 9 —— 57
*/
func fill(value string, x, y int, board [][]string) {
	board[y][x] = "0" + value

	hashmap1[y][value] = struct{}{}
	hashmap2[x][value] = struct{}{}
	kuai_row := y / 3
	kuai_column := x / 3
	hashmap3[pointToIndex(kuai_column, kuai_row)][value] = struct{}{}
}
func wipe(value string, x, y int, board [][]string) {
	board[y][x] = "."

	delete(hashmap1[y], value)
	delete(hashmap2[x], value)
	kuai_row := y / 3
	kuai_column := x / 3
	delete(hashmap3[pointToIndex(kuai_column, kuai_row)], value)
}

func goNext() {
	if x == 8 {
		y++
		x = 0
	} else {
		x++
	}
}

func formatOutput(board [][]string) {
	for _, v := range board {
		for _, value := range v {
			fmt.Printf("%2v ", value)
		}
		fmt.Println()
	}
	fmt.Println()
}

func formatBlock() {
	for i, m := range hashmap3 {
		fmt.Printf("num: %v, block: ", i)
		for k := range m {
			fmt.Printf("%v ", k)
		}
		fmt.Println()
	}
}
