/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode-cn.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (67.07%)
 * Likes:    964
 * Dislikes: 0
 * Total Accepted:    101.9K
 * Total Submissions: 151.9K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过填充空格来解决数独问题。
 *
 * 数独的解法需 遵循如下规则：
 *
 *
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
 *
 *
 * 数独部分空格内已填入了数字，空白格用 '.' 表示。
 *
 *
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：board =
 * [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
 *
 * 输出：[["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
 * 解释：输入的数独如上图所示，唯一有效的解决方案如下所示：
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * board.length == 9
 * board[i].length == 9
 * board[i][j] 是一位数字或者 '.'
 * 题目数据 保证 输入数独仅有一个解
 *
 *
 *
 *
 *
 */
package main

import (
	"fmt"
	"log"
)

func formatOutput(board [][]byte) {
	for _, v := range board {
		for _, value := range v {
			fmt.Printf("%2v ", value)
		}
		fmt.Println()
	}
	fmt.Println()
}

// func main() {
// 	table := [][]byte{
// 		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
// 		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
// 		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
// 		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
// 		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
// 		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
// 		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
// 		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
// 		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
// 	}
// 	solveSudoku(table)
// 	formatOutput(table)

// }

// @lc code=start
var hashmap1 [9]map[byte]struct{} // 每一行
var hashmap2 [9]map[byte]struct{} // 每一列
var hashmap3 [9]map[byte]struct{} // 每一块
var x int
var y int

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

func initalize(board [][]byte) {
	for row, value := range board {
		for column, v := range value {
			if hashmap1[row] == nil {
				hashmap1[row] = make(map[byte]struct{}, 9)
			}
			hashmap1[row][v] = struct{}{}
			if hashmap2[column] == nil {
				hashmap2[column] = make(map[byte]struct{}, 9)
			}
			hashmap2[column][v] = struct{}{}
			kuai_row := row / 3
			kuai_column := column / 3
			if hashmap3[pointToIndex(kuai_column, kuai_row)] == nil {
				hashmap3[pointToIndex(kuai_column, kuai_row)] = make(map[byte]struct{}, 9)
			}
			hashmap3[pointToIndex(kuai_column, kuai_row)][v] = struct{}{}
		}
	}
}

func solveSudoku(board [][]byte) {
	// 初始化
	initalize(board)
	if ok := recursive(board); ok {
		log.Println("Yes")
		return
	}
	log.Println("No")
	log.Println(hashmap1)
	log.Println(hashmap2)
	log.Println(hashmap3)
}

/*
	递归以每一行为起始，每一次递归都从可填的最小数字开始进行递归填充，并判断是够满足
	1、每一行不重复
	2、每一列不重复
	3、每3×3不重复
	hashmap 保存每一行/列/3×3是否存在
*/
func recursive(board [][]byte) bool {
	// 从每一行开始
	// 当全部填充完成，goNext会越界
	if y == 9 {
		return true
	}
	if board[y][x] != '.' {
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
		if _, ok := hashmap1[curY][byte(i+48)]; !ok {
			// 判断列不存在
			if _, ok := hashmap2[curX][byte(i+48)]; !ok {
				// 判断块不存在
				kuai_column := curX / 3
				kuai_row := curY / 3
				if _, ok := hashmap3[pointToIndex(kuai_column, kuai_row)][byte(i+48)]; !ok {
					// 满足条件，尝试填充
					fill(byte(48+i), curX, curY, board)
					goNext()
					if ok := recursive(board); !ok {
						// 有冲突，则清除填充内容
						wipe(byte(48+i), curX, curY, board)
						x, y = curX, curY
					} else {
						return true
					}

				}
			}
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
func fill(value byte, x, y int, board [][]byte) {
	board[y][x] = value

	hashmap1[y][value] = struct{}{}
	hashmap2[x][value] = struct{}{}
	kuai_row := y / 3
	kuai_column := x / 3
	hashmap3[pointToIndex(kuai_column, kuai_row)][value] = struct{}{}
}
func wipe(value byte, x, y int, board [][]byte) {
	board[y][x] = '.'

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

// @lc code=end
