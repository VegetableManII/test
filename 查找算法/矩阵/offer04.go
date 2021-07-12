package main

func main() {

}

var row int
var column int

func findNumberIn2DArray(matrix [][]int, target int) bool {
	row = len(matrix)
	column = len(matrix[0])
	return exist(matrix, target)

}
func exist(matrix [][]int, target int) bool {
	for i := 0; i < row; i++ {
		if matrix[i][column-1] < target {
			row = i + 1
			exist(matrix, target)
		} else if matrix[i][column-1] > target {
			row = i
			for i := 0; i < column; i++ {
				if matrix[row][i] == target {
					return true
				}
			}
		} else {
			return true
		}
	}
	return false
}
