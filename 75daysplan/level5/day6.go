package level5

import (
	"sort"
)

func MinFallingPathSum(matrix [][]int) int {

	//dp[i][j] += min{dp[i-1][j-1],dp[i-1][j],dp[i-1][j+1]}

	for i := 1; i < len(matrix); i++ {

		for j := 0; j < len(matrix[0]); j++ {

			if j == 0 && j == len(matrix[0])-1 {
				matrix[i][j] += matrix[i-1][j]
			} else if j == 0 {
				matrix[i][j] += getTheMin(matrix[i-1][j], matrix[i-1][j+1])
			} else if j == len(matrix[0])-1 {
				matrix[i][j] += getTheMin(matrix[i-1][j], matrix[i-1][j-1])
			} else {
				matrix[i][j] += getTheMin(matrix[i-1][j], getTheMin(matrix[i-1][j+1], matrix[i-1][j-1]))
			}
		}
	}

	sort.Ints(matrix[len(matrix)-1])
	return matrix[len(matrix)-1][0]
}
func getTheMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FindFarmland(land [][]int) [][]int {

	res := [][]int{}
	visited := make([][]bool, len(land))
	for i := 0; i < len(land); i++ {
		visited[i] = make([]bool, len(land[0]))
	}

	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[0]); j++ {
			dfsFindFarmLand(i, j, land, &res, &visited)
		}
	}

	return res
}

var state = [][]int{{0, 1}, {0, -1}, {1, 1}, {1, -1}}

func dfsFindFarmLand(x, y int, land [][]int, res *[][]int, visited *[][]bool) {

	if (*visited)[x][y] == false && land[x][y] == 1 {
		(*visited)[x][y] = true
		countFarm(land, visited, x, y, res)
	}
	(*visited)[x][y] = true
	for i := 0; i < len(state); i++ {
		x0 := x + state[i][0]
		y0 := y + state[i][1]

		if isInBoard(x0, y0, land) {
			if (*visited)[x0][y0] == true {
				continue
			}
			dfsFindFarmLand(x0, y0, land, res, visited)
		}
	}

	return
}
func isInBoard(x, y int, land [][]int) bool {

	return x >= 0 && x < len(land) && y >= 0 && y < len(land[0])
}
func countFarm(land [][]int, visited *[][]bool, x, y int, res *[][]int) {

	end := []int{x, y}

	for board := 1; board <= len(land)-x; board++ {
		if isFarm(land, board, x, y, visited) {
			end = []int{x + board - 1, y + board - 1}
		}
	}
	*res = append(*res, []int{x, y, end[0], end[1]})
	return
}
func isFarm(land [][]int, board int, x, y int, visited *[][]bool) bool {

	for i := x; i < x+board; i++ {
		for j := y; j < y+board; j++ {
			if isInBoard(i, j, land) {
				if land[i][j] == 0 {
					return false
				} else {
					(*visited)[i][j] = true
				}
			}
		}
	}
	return true
}

/*
You are given a 2D integer array logs where each logs[i] = [birthi, deathi]
indicates the birth and death years of the ith person.
The population of some year x is the number of people alive during that year.
The ith person is counted in year x's population if x
is in the inclusive range [birthi, deathi - 1].
Note that the person is not counted in the year that they die.

Return the earliest year with the maximum population.
Example 1:
//1950,1961],[1960,1971],[1970,1981
1950--1961
    1960--1971
        1970--1981
Input: logs = [[1993,1999],[2000,2010]]
Output: 1993
Explanation: The maximum population is 1, and 1993 is the earliest year with this population.
*/
