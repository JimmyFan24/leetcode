package level4

import (
	"math"
	"sort"
)

type SubrectangleQueries struct {
	Subrectangle [][]int
}

func Constructor1476(rectangle [][]int) SubrectangleQueries {

	return SubrectangleQueries{Subrectangle: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {

	for i := row1; i <= row2; i++ {

		for j := col1; j <= col2; j++ {
			this.Subrectangle[i][j] = newValue
		}

	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {

	return this.Subrectangle[row][col]
}

/*
  [3, 0, 8, 4],
  [2, 4, 5, 7],
  [9, 2, 6, 3],
  [0, 3, 1, 0],
*/
func MaxIncreaseKeepingSkyline(grid [][]int) int {
	rowMax := make(map[int]int, len(grid))
	rolMax := make(map[int]int, len(grid[0]))
	//1.找出每一行最大的
	for i, row := range grid {
		max := 0
		for j := 0; j < len(row); j++ {
			if max < row[j] {
				max = row[j]

			}
		}
		rowMax[i] = max
	}
	//2.找出每一列最大的
	for i := 0; i < len(grid[0]); i++ {

		max := 0

		for j := 0; j < len(grid); j++ {
			if max < grid[j][i] {
				max = grid[j][i]

			}

		}
		rolMax[i] = max
	}
	canAdd := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			//遍历二维数组,如果不是最值,那么进行自增
			if grid[i][j] != rowMax[i] && grid[i][j] != rolMax[j] {

				canAdd += getTheMin(rowMax[i], rolMax[j]) - grid[i][j]

			}
		}
	}
	return canAdd
}
func getTheMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinOperations(boxes string) []int {

	answer := []int{}
	//110
	for i := 0; i < len(boxes); i++ {

		step := 0
		for j := 0; j < len(boxes); j++ {
			if j == i {
				continue
			}
			if boxes[j] == '1' {
				step += int(math.Abs(float64(j) - float64(i)))
			}

		}
		answer = append(answer, step)
	}
	return answer
}

/*
Given n points on a 2D plane where points[i] = [xi, yi],
Return the widest vertical area between two points such that no points are inside the area.

A vertical area is an area of fixed-width extending infinitely along the y-axis (i.e., infinite height). The widest vertical area is the one with the maximum width.

Note that points on the edge of a vertical area are not considered included in the area.
*/
func MaxWidthOfVerticalArea(points [][]int) int {

	X := []int{}
	for _, v := range points {
		X = append(X, v[0])
	}
	sort.Ints(X)
	//points = [[3,1],[9,0],[1,0],[1,4],[5,3],[8,8]]

	max := math.MinInt32
	for i := 0; i < len(X)-1; i++ {

		if X[i]-X[i-1] > max {
			max = X[i] - X[i-1]
		}
	}
	return max
}

func PivotArray(nums []int, pivot int) []int {

	res := []int{}

	//1.先排小的
	for _, v := range nums {
		if v < pivot {
			res = append(res, v)
		}
	}
	for _, v := range nums {
		if v == pivot {
			res = append(res, v)
		}
	}
	for _, v := range nums {
		if v > pivot {
			res = append(res, v)
		}
	}
	return res
}

func ExecuteInstructions(n int, startPos []int, s string) []int {

	//下一条指令将会导致机器人移动到网格外。
	//没有指令可以执行
	answer := []int{}

	for i := 0; i < len(s); i++ {

		count := 0
		x, y := startPos[0], startPos[1]
		j := i
		for inBoard(x, y, n) && j < len(s) {
			switch s[j] {
			case 'L':
				y -= 1
			case 'R':
				y += 1
			case 'U':
				x -= 1
			default:
				x += 1
			}
			if inBoard(x, y, n) {
				count++
				j++
			} else {
				break
			}

		}
		answer = append(answer, count)
	}
	return answer

}
func inBoard(x, y, n int) bool {

	return x < n && y < n && x >= 0 && y >= 0
}
func DiagonalSort(mat [][]int) [][]int {

	res := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		res[i] = make([]int, len(mat[0]))
	}
	x := len(mat) - 1
	y := 0
	for x >= 0 || y < len(mat[0]) {

		sortSingalLine(mat, &res, x, y)

		if x == 0 {
			y++
		}
		if x > 0 {
			x--
		}
		if y >= len(mat[0]) {
			break
		}
	}

	return res

}
func sortSingalLine(mat [][]int, res *[][]int, x, y int) {

	tmp := []int{}
	i := x
	j := y
	for i < len(mat) && j < len(mat[0]) {

		tmp = append(tmp, mat[i][j])
		i++
		j++
	}
	sort.Ints(tmp)
	ii := x
	jj := y
	for ii < len(mat) && jj < len(mat[0]) {
		(*res)[ii][jj] = tmp[0]
		ii++
		jj++
		tmp = tmp[1:]
	}
}
func CheckArithmeticSubarrays(nums []int, l []int, r []int) []bool {

	res := []bool{}
	for i := 0; i < len(l); i++ {
		tmp := make([]int, r[i]-l[i]+1)
		copy(tmp, nums[l[i]:r[i]+1])
		if checkSubarrays(tmp) {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
func checkSubarrays(arr []int) bool {

	sort.Ints(arr)
	//2,3,4,5,6
	tmp := (arr[len(arr)-1] - arr[0]) / (len(arr) - 1)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]+tmp != arr[i+1] {
			return false
		}
	}
	return true
}

func FindAndReplacePattern(words []string, pattern string) []string {

	match := []string{}
	pM := make(map[rune]int, 26)
	for _, v := range pattern {
		pM[v]++
	}

	for _, v := range words {
		if isMatch(v, pattern) {
			match = append(match, v)
		}
	}
	return match
}
func isMatch(word string, pattern string) bool {

	wM := make(map[rune]rune, 26)
	pM := make(map[rune]rune, 26)
	for i, v := range word {
		if _, ok := wM[v]; !ok {
			wM[v] = rune(pattern[i])
		} else {
			if wM[v] != rune(pattern[i]) {
				return false
			}
		}

	}
	for i, v := range pattern {
		if _, ok := pM[v]; !ok {
			pM[v] = rune(word[i])
		} else {
			if pM[v] != rune(word[i]) {
				return false
			}
		}

	}
	return true

}
