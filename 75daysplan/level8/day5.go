package level8

/*
Given an m x n binary matrix filled with 0's and 1's,
find the largest square containing only 1's and return its area.
*/
func MaximalSquare(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}
	//dp[i][j] 含义为正方形以 dp[i][j] 作为右下角时的最大边长值

	//dp[i][j] 与 dp[i-1][j]、dp[i][j-1]、dp[i-1][j-1]
	//如果该位置的值是 0，则 dp(i, j) = 0，因为当前位置不可能在由1组成的正方形中；
	//如果该位置的值是 1，则 dp(i, j)值由其上方、左方和左上方的三个相邻位置的 dp 值决定。具体而言，
	//当前位置的元素值等于三个相邻位置的元素中的最小值加 1，状态转移方程如下：
	//dp[i][j] = min( dp[i-1][j], dp[i-1][j-1], dp[i][j-1] ）+ 1
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				dp[i][j] = int(matrix[i][j]) - int('0')
			} else {
				dp[i][j] = getmin(dp[i-1][j], getmin(dp[i-1][j-1], dp[i][j-1])) + 1

			}
			res = getmax(res, dp[i][j])
		}
	}
	return res * res
}
func getmin(a, b int) int {
	if a < b {
		return a
	}
	return b

}

func isSquare(x, y int, square int, matrix [][]string) bool {

	if x+square-1 >= len(matrix) || y+square-1 >= len(matrix[0]) {
		return false
	}
	for i := x; i < x+square; i++ {
		for j := y; j < y+square; j++ {
			if matrix[i][j] == "0" {
				return false
			}
		}
	}
	return true
}

/*

Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix.
This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
*/

func SearchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	r := len(matrix)
	c := len(matrix[0])
	//从左下角开始
	x, y := r-1, 0
	for x >= 0 && y < c {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			y++
		} else {
			x--
		}
	}

	return false

}
func binarySearchMatrix(matrix [][]int, tRow, BRow, lRol, RRol int, target int) bool {

	mRow := tRow + (BRow-tRow)/2
	mRol := lRol + (RRol-lRol)/2

	if matrix[mRow][mRol] == target {
		return true
	}

	if matrix[mRow][mRol] > target {
		if mRow+1 < len(matrix) && mRol+1 < len(matrix[0]) {
			//如果大于当前的对角线的值,而且小于下一个对角线的值,那么就在中间找
			next := matrix[mRow+1][mRol+1]
			if target > next {

			}
		}
		return binarySearchMatrix(matrix, tRow, mRow, lRol, mRol, target)
	}

	return binarySearchMatrix(matrix, mRow, BRow, mRol, RRol, target)

}
