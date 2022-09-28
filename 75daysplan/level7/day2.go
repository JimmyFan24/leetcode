package level7

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
Input: arr = [0,1,2,3,4,5,6,7,8]
Output: [0,1,2,4,8,3,5,6,7]
Explantion: [0] is the only integer with 0 bits.
[1,2,4,8] all have 1 bit.
[3,5,6] have 2 bits.
[7] has 3 bits.
The sorted array by bits is [0,1,2,4,8,3,5,6,7]
*/
func SortByBits(arr []int) []int {

	bitMap := make(map[int][]int, len(arr))
	for _, v := range arr {
		bitMap[countBits(v)] = append(bitMap[countBits(v)], v)
	}
	res := []int{}
	keys := []int{}
	for k, _ := range bitMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		sort.Ints(bitMap[v])
		for _, vv := range bitMap[v] {
			res = append(res, vv)
		}
	}
	return res

}
func countBits(a int) int {

	count := 0

	for a > 0 {
		a = a & (a - 1)
		count++
	}
	return count
}

func MinTimeToType(word string) int {

	//1.如果下标之差小于13,顺时针,否则逆时针
	costtime := 0

	if word[0] == 'a' {
		costtime = 1
	} else {
		if int(math.Abs(float64(word[0])-float64('a'))) > 13 {
			costtime += 26 - int(math.Abs(float64(word[0])-float64('a'))) + 1
		} else {
			costtime += int(math.Abs(float64(word[0])-float64('a'))) + 1
		}
	}
	for i := 1; i < len(word); i++ {
		offset := int(math.Abs(float64(word[i]) - float64(word[i-1])))
		if offset > 13 {
			//逆时针
			costtime += 26 - offset + 1
		} else {
			//顺时针
			costtime += offset + 1
		}
	}
	return costtime
}

/*
xample 1:
Input: s = "loveleetcode", c = "e"
Output: [3,2,1,0,1,0,0,1,2,2,1,0]
Explanation: The character 'e' appears at indices 3, 5, 6, and 11 (0-indexed).
The closest occurrence of 'e' for index 0 is at index 3, so the distance is abs(0 - 3) = 3.
The closest occurrence of 'e' for index 1 is at index 3, so the distance is abs(1 - 3) = 2.
For index 4, there is a tie between the 'e' at index 3 and the 'e' at index 5, but the distance is still the same: abs(4 - 3) == abs(4 - 5) = 1.
The closest occurrence of 'e' for index 8 is at index 6, so the distance is abs(8 - 6) = 2.
*/
func ShortestToChar(s string, c byte) []int {

	res := []int{}
	index := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			index[i] = true
		}
	}
	for i := 0; i < len(s); i++ {
		if !index[i] {

			right := i
			left := i
			for right < len(s) && !index[right] {
				right++
			}
			for left >= 0 && !index[left] {
				left--
			}
			if left == -1 && right != len(s) {
				res = append(res, int(math.Abs(float64(right)-float64(i))))
			} else if left != -1 && right == len(s) {
				res = append(res, int(math.Abs(float64(left)-float64(i))))
			} else {
				offsetRight := int(math.Abs(float64(right) - float64(i)))
				offsetLeft := int(math.Abs(float64(left) - float64(i)))
				if offsetRight < offsetLeft {
					res = append(res, offsetRight)
				} else {
					res = append(res, offsetLeft)
				}
			}

		} else {
			res = append(res, 0)
		}
	}
	return res
}

func CheckDistances(s string, distance []int) bool {

	for i, v := range distance {

		index := strings.Index(s, string(rune('a'+i)))
		if index != -1 {
			if index+v+1 >= len(s) {
				return false
			}
			if s[index+v+1] != s[index] {
				return false
			}
		}

	}
	return true
}

func MinimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minAbsDiff := math.MaxInt32
	//1.找出最小的绝对值差
	for i := 0; i < len(arr)-1; i++ {

		abs := int(math.Abs(float64(arr[i+1]) - float64(arr[i])))
		if abs < minAbsDiff {
			minAbsDiff = abs
		}
	}
	res := [][]int{}
	//2.开始找对
	for i := 0; i < len(arr)-1; i++ {
		abs := int(math.Abs(float64(arr[i+1]) - float64(arr[i])))
		if abs == minAbsDiff {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}
func Intersection(nums [][]int) []int {

	l := 0
	for _, v := range nums {
		if len(v) > l {
			l = len(v)
		}

	}
	res := make(map[int]int, l*len(nums))
	for _, num := range nums {

		for _, n := range num {

			res[n]++
		}
	}
	values := []int{}
	for k, v := range res {
		if v == len(nums) {
			values = append(values, k)
		}

	}
	sort.Ints(values)
	return values
}

func IslandPerimeter(grid [][]int) int {

	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	count := 0
	bfsIslandPerimeter(grid, &count, &visited, 0, 0)
	return count
}

var offset = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0}}

func bfsIslandPerimeter(grid [][]int, count *int, visited *[][]bool, x, y int) {

	if !(*visited)[x][y] {
		(*visited)[x][y] = true
	}
	if grid[x][y] == 1 {
		//遍历周围四个点
		for _, v := range offset {
			x1 := v[0] + x
			y1 := v[1] + y
			if inboard(x1, y1, grid) && grid[x1][y1] == 0 {
				*count++
			}
			if !inboard(x1, y1, grid) {
				*count++
			}
		}
	}
	for _, v := range offset {

		x0 := v[0] + x
		y0 := v[1] + y
		if inboard(x0, y0, grid) && !(*visited)[x0][y0] {
			(*visited)[x0][y0] = true
			bfsIslandPerimeter(grid, count, visited, x0, y0)
		}
	}
}
func inboard(x, y int, grid [][]int) bool {

	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func FindWords(words []string) []string {
	res := []string{}
	keyBoardMap := make([]map[string]int, 3)
	var keyboard = []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for i, v := range keyboard {
		keyBoardMap[i] = make(map[string]int, len(v))
	}
	for i, v := range keyboard {
		for _, vv := range v {
			keyBoardMap[i][string(vv)]++
		}
	}
	for _, v := range words {
		if isFind(v, keyBoardMap) {
			res = append(res, v)
		}
	}
	return res
}

func isFind(s string, keyBoardMap []map[string]int) bool {

	for _, keyMap := range keyBoardMap {

		i := 0
		for _, v := range s {
			c := strings.ToLower(string(v))
			if keyMap[c] < 1 {
				break
			}
			i++
		}
		if i == len(s) {
			return true
		}
	}
	return false

}

func MinStartValue(nums []int) int {

	//-3,2,-3,4,2
	//dp[i] 表示 能让0-i 的累加都大于1 的最小正整数
	//dp[i] = if nums[i] >0 dp[i] = dp[i-1]
	//       if nums[i] < 0 dp[i] = dp[i-1] + 1
	res := 0
	for n := 1; ; n++ {
		sum := n
		index := 0
		for _, v := range nums {
			sum += v

			if sum < 1 {
				break
			}
			index++
		}
		if index == len(nums) {
			res = n
			break
		}
	}
	return res
}

func CountPrimeSetBits(left int, right int) int {

	res := 0
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19}
	m := make(map[int]bool, 8)
	for _, v := range prime {
		m[v] = true
	}
	for i := left; i <= right; i++ {
		if m[countOneBinary(i)] {
			res++
		}

	}
	return res
}
func countOneBinary(n int) int {
	count := 0
	for n > 0 {

		n = n & (n - 1)
		count++
	}
	return count
}

/*
如果一个正方形矩阵满足下述全部条件，则称之为一个 X 矩阵 ：
矩阵对角线上的所有元素都 不是 0
矩阵中所有其他元素都是 0
*/
func CheckXMatrix(grid [][]int) bool {
	//1.检查对角线是否不是0

	/*
		2,0,0,1],
		[0,3,1,0],
		[0,5,2,0],
		[4,0,0,2]
	*/
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == j {
				if grid[i][j] == 0 {
					return false
				}
				continue
			}
			if i+j == len(grid)-1 {
				if grid[i][j] == 0 {
					return false
				}
				continue
			}
			if grid[i][j] != 0 {
				return false
			}

		}
	}
	return true
}
func CountStudents(students []int, sandwiches []int) int {

	m1 := make(map[int]int, len(students))
	m2 := make(map[int]int, len(sandwiches))

	for _, v := range students {
		m1[v]++
	}
	for _, v := range sandwiches {
		m2[v]++
	}
	for _, v := range sandwiches {

		if m1[v] < 1 {
			break
		} else {
			m1[v]--
		}
	}
	count := 0
	for _, v := range m1 {
		if v > 0 {
			count += v
		}
	}
	return count
}
func ShiftGrid(grid [][]int, k int) [][]int {

	/*
		1.Element at grid[i][j] moves to grid[i][j + 1].
		2.Element at grid[i][n - 1] moves to grid[i + 1][0].
		3.Element at grid[m - 1][n - 1] moves to grid[0][0].
	*/

	for i := 0; i < k; i++ {
		m := len(grid)
		n := len(grid[0])
		res := make([][]int, m)
		for i := 0; i < m; i++ {
			res[i] = make([]int, n)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				//1.Element at grid[i][j] moves to grid[i][j + 1].
				if j+1 < n {
					grid[i][j] = res[i][j+1]
				}
				//2.Element at grid[i][n - 1] moves to grid[i + 1][0].
				if i+1 < m {
					grid[i][n-1] = res[i+1][0]
				}
				//3.Element at grid[m - 1][n - 1] moves to grid[0][0].
				grid[m-1][n-1] = res[0][0]
			}
		}
		grid = res
	}

	return grid

}

/*
A square triple (a,b,c) is a triple where a, b, and c are integers and a2 + b2 = c2.

Given an integer n, return the number of square triples such that 1 <= a, b, c <= n.

Example 1:

Input: n = 5
Output: 2
Explanation: The square triples are (3,4,5) and (4,3,5).
*/
func CountTriples(n int) int {

	//1.a2 + b2 = c2.
	//2.1 <= a, b, c <= n.
	count := 0
	for i := 1; i <= n-2; i++ {
		for j := n; j > i+1; j-- {
			m := j - 1
			for ; m > i; m-- {
				if i*i+m*m == j*j {
					count++
				}
			}
		}
	}
	return count * 2
}
func FindRelativeRanks(score []int) []string {

	res := []string{}
	c := make([]int, len(score))
	copy(c, score)
	sort.Ints(c)
	scoreM := make(map[int]int, len(score))
	rank := 1
	for i := len(c) - 1; i >= 0; i-- {
		scoreM[c[i]] = rank
		rank++
	}
	for _, v := range score {
		if scoreM[v] == 1 {
			res = append(res, "Gold Medal")
		} else if scoreM[v] == 2 {
			res = append(res, "Silver Medal")
		} else if scoreM[v] == 3 {
			res = append(res, "Bronze Medal")
		}
		res = append(res, strconv.Itoa(scoreM[v]))
	}
	return res
}
