package level7

import (
	"math"
	"sort"
	"strings"
)

/*
[
[".",".",".",".",".",".",".","."],
[".",".",".","p",".",".",".","."],
[".",".","p","p",".",".",".","."],
[".","p","p","R",".","p",".","p"],
[".",".",".","p",".",".",".","."],
[".",".",".",".",".","p",".","."],
[".",".",".","p",".",".",".","."],
[".",".",".",".",".",".",".","."]]
*/
func NumRookCaptures(board [][]byte) int {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	startx, starty := 0, 0
	for i, v := range board {
		for j, vv := range v {
			if vv == 'R' {
				startx = i
				starty = j
			}
		}
	}
	count := 0
	//1.检查该行后半部分有没有
	for i := starty + 1; i < len(board[0]); i++ {
		if board[startx][i] == 'p' {
			count++
			break
		} else if board[startx][i] == 'B' {
			break
		}
	}
	//2.检查该行前半部分有没有
	for i := starty - 1; i >= 0; i-- {
		if board[startx][i] == 'p' {
			count++
			break
		} else if board[startx][i] == 'B' {
			break
		}
	}
	//3.检查该列上半部分有没有
	for i := startx - 1; i >= 0; i-- {
		if board[i][starty] == 'p' {
			count++
			break
		} else if board[i][starty] == 'B' {
			break
		}
	}
	//4.检查该列下半部分有没有
	for i := startx + 1; i < len(board); i++ {
		if board[i][starty] == 'p' {
			count++
			break
		} else if board[i][starty] == 'B' {
			break
		}
	}
	//dfsRookCaptures(board,&visited,&count,startx,starty)
	return count
}
func inBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

var position = [][]int{
	{1, 0}, {-1, 0}, {0, -1}, {0, 1},
}

func dfsRookCaptures(board [][]byte, visited *[][]bool, count *int, x, y int) {
	if board[x][y] == 'p' {
		*count += 1
		return
	}
	if board[x][y] == 'B' {
		return
	}

	for _, v := range position {
		x0 := v[0] + x
		y0 := v[1] + y
		if inBoard(board, x0, y0) && !(*visited)[x0][y0] {
			(*visited)[x0][y0] = true
			dfsRookCaptures(board, visited, count, x0, y0)

		}
	}
}
func IsToeplitzMatrix(matrix [][]int) bool {
	row := len(matrix)
	rol := len(matrix[0])
	for i := 0; i < row; i++ {
		if !judge(matrix, i, 0) {
			return false
		}
	}
	for i := 0; i < rol; i++ {
		if !judge(matrix, 0, i) {
			return false
		}
	}
	return true

}
func judge(matrix [][]int, x, y int) bool {

	for x < len(matrix)-1 && y < len(matrix[0])-1 {
		if matrix[x][y] != matrix[x+1][y+1] {
			return false
		}
		x++
		y++
	}
	return true
}

/*
A sequence of numbers is called an arithmetic
progression if the difference between any two consecutive elements is the same.

Given an array of numbers arr,
return true if the array can be rearranged to form an arithmetic progression. Otherwise, return false.
*/
func CanMakeArithmeticProgression(arr []int) bool {

	sort.Ints(arr)
	dis := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != dis {
			return false
		}
	}
	return true
}

/*
Given a string of English letters s, return the greatest English letter which occurs as both a
lowercase and uppercase letter in s.
The returned letter should be in uppercase.
If no such letter exists, return an empty string.

An English letter b is greater than another letter a if b appears after a in the English alphabet.
*/
func GreatestLetter(s string) string {
	greatest := ""
	lower := make(map[byte]int, 26)
	//1.先统计出现过的小写字母
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			//统计出现过的小写字母
			lower[s[i]]++
		}
	}
	//2.统计出现的大写字母数量
	upper := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			//如果是大写字母,而且出现的小写字母也是一个
			upper[s[i]]++
		}
	}

	for k, v := range upper {
		if v >= 1 {
			//查询是否小写也是只有一个
			l := k - 'A' + 'a'
			if lower[l] >= 1 {
				if greatest == "" {
					greatest = string(k)
				}
				if greatest > string(k) {
					greatest = string(k)
				}
			}
		}
	}
	return strings.ToUpper(greatest)

}

/*
You are given four integers row, cols, rCenter, and cCenter.
There is a rows x cols matrix and you are on the cell with the coordinates (rCenter, cCenter).

Return the coordinates of all cells in the matrix,
sorted by their distance from (rCenter, cCenter)
from the smallest distance to the largest distance.
You may return the answer in any order that satisfies this condition.

The distance between two cells (r1, c1) and (r2, c2) is |r1 - r2| + |c1 - c2|
Example 1:
Input: rows = 1, cols = 2, rCenter = 0, cCenter = 0
Output: [[0,0],[0,1]]
Explanation: The distances from (0, 0) to other cells are: [0,1]
*/
func AllCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {

	distance := make(map[int][][]int, rows*cols)

	res := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dis := countDistance(rCenter, cCenter, []int{i, j})
			distance[dis] = append(distance[dis], []int{i, j})
		}
	}

	keys := []int{}
	for k, _ := range distance {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		for _, dis := range distance[v] {
			res = append(res, dis)
		}
	}

	return res
}
func countDistance(rCenter int, cCenter int, cell []int) int {

	return int(math.Abs(float64(cell[0])-float64(rCenter))) + int(math.Abs(float64(cell[1])-float64(cCenter)))
}

/*
Example 1:

Input: mat =
[[1,1,0,0,0],

	[1,1,1,1,0],
	[1,0,0,0,0],
	[1,1,0,0,0],
	[1,1,1,1,1]],

k = 3
Output: [2,0,3]
Explanation:
The number of soldiers in each row is:
- Row 0: 2
- Row 1: 4
- Row 2: 1
- Row 3: 2
- Row 4: 5
The rows ordered from weakest to strongest are [2,0,3,1,4].
*/
func KWeakestRows(mat [][]int, k int) []int {

	res := []int{}

	m := make(map[int][]int, len(mat))
	for i, v := range mat {
		sum := countOne(v)
		m[sum] = append(m[sum], i)
	}
	keys := []int{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := 0; i < len(keys); i++ {
		for _, v := range m[keys[i]] {
			res = append(res, v)
		}
	}
	return res[:k]
}
func countOne(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

/*
Input: items1 = [[1,1],[4,5],[3,8]], items2 = [[3,1],[1,5]]
Output: [[1,6],[3,9],[4,5]]
Explanation:
The item with value = 1 occurs in items1 with weight = 1
and in items2 with weight = 5, total weight = 1 + 5 = 6.
The item with value = 3 occurs in items1 with weight = 8
and in items2 with weight = 1, total weight = 8 + 1 = 9.
The item with value = 4 occurs in items1 with weight = 5, total weight = 5.
Therefore, we return [[1,6],[3,9],[4,5]].
*/
func MergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {

	weight := make(map[int]int, len(items1)+len(items2))
	for _, v := range items1 {
		weight[v[0]] += v[1]
	}
	for _, v := range items2 {
		weight[v[0]] += v[1]
	}
	res := [][]int{}
	keys := []int{}
	for k, _ := range weight {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		res = append(res, []int{v, weight[v]})
	}
	return res
}

/*
You are given an n x n grid where we place some 1 x 1 x 1 cubes
that are axis-aligned with the x, y, and z axes.
Each value v = grid[i][j] represents a tower of v cubes placed on top of the cell (i, j).
We view the projection of these cubes onto the xy, yz, and zx planes.
A projection is like a shadow, that maps our 3-dimensional figure to a 2-dimensional plane.
We are viewing the "shadow" when looking at the cubes from the top, the front, and the side.
Return the total area of all three projections.
*/
/*
[1,2],[3,4]
*/
func ProjectionArea(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	top, left, front := 0, 0, 0
	//1.top 如果grid[i][j] != 0 ,那么累加
	for _, v := range grid {
		for _, vv := range v {
			if vv != 0 {
				top++
			}
		}
	}
	//2.正面,计算每一行最大值之和
	for i := 0; i < row; i++ {
		sum := 0
		for j := 0; j < col; j++ {
			sum = getBiggerOne(sum, grid[i][j])
		}
		front += sum
	}
	//3.计算侧面,每一列最大求和
	for i := 0; i < col; i++ {
		sum := 0
		for j := 0; j < row; j++ {
			sum = getBiggerOne(sum, grid[j][i])
		}
		left += sum
	}

	return top + left + front
}
func getBiggerOne(a, b int) int {

	if a > b {
		return a
	}
	return b

}
