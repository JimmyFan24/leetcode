package level7

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func SmallestEvenMultiple(n int) int {
	if n%2 == 5 {
		return n
	}
	return n * 2
}

/*
Example 1:

Input: names = ["Mary","John","Emma"], heights = [180,165,170]
Output: ["Mary","Emma","John"]
Explanation: Mary is the tallest, followed by Emma and John.
*/

func SortPeople(names []string, heights []int) []string {

	res := []string{}
	nameMap := make(map[int]string, len(names))
	for i := 0; i < len(names); i++ {
		nameMap[heights[i]] = names[i]
	}
	keys := []int{}
	for k, _ := range nameMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i := len(keys) - 1; i >= 0; i-- {
		res = append(res, nameMap[keys[i]])
	}
	return res

}

/*
Example 1:
Input: image = [[1,1,0],[1,0,1],[0,0,0]]
Output: [[1,0,0],[0,1,0],[1,1,1]]
Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
*/

func FlipAndInvertImage(image [][]int) [][]int {

	filp(&image)
	invert(&image)
	return image
}
func filp(image *[][]int) {
	m := len(*image)
	n := len((*image)[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			(*image)[i][j], (*image)[i][n-j-1] = (*image)[i][n-j-1], (*image)[i][j]

		}
	}
}
func invert(image *[][]int) {
	m := len(*image)
	n := len((*image)[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (*image)[i][j] == 0 {
				(*image)[i][j] = 1
			} else {
				(*image)[i][j] = 0
			}
		}
	}
}

/*
Input: s = "(()()) (()) (()(()))"
Output: "()()()()(())"
Explanation:
The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".
*/
func RemoveOuterParentheses(s string) string {

	stack := []string{}
	res := ""
	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			stack = append(stack, string(s[i]))
			if len(stack) > 1 {
				res += "("
			}
		} else if s[i] == ')' {

			if len(stack) > 1 {
				res += ")"
				stack = stack[:len(stack)-1]
			} else if len(stack) == 1 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return res
}

/*
Input: s = "a1c1e1"
Output: "abcdef"
Explanation: The digits are replaced as follows:
- s[1] -> shift('a',1) = 'b'
- s[3] -> shift('c',1) = 'd'
- s[5] -> shift('e',1) = 'f'
*/
func ReplaceDigits(s string) string {

	res := ""
	for i := 0; i < len(s); i++ {

		if s[i] >= '0' && s[i] <= '9' {
			index, _ := strconv.Atoi(string(s[i]))
			after := shift(rune(s[i-1]), index)
			res += string(after)
		} else {
			res += string(s[i])
		}
	}
	return res

}
func shift(before rune, index int) rune {

	return before + rune(index)
}

func FreqAlphabets(s string) string {

	res := ""
	for len(s) > 0 {
		index := strings.Index(s, "#")
		if index != -1 && index <= 2 {
			ss := s[:index]
			b, _ := strconv.Atoi(ss)
			res += string(rune('a' + b - 1))
			s = s[index+1:]
		} else {
			ss := s[:1]
			b, _ := strconv.Atoi(ss)
			res += string(rune('a' + b - 1))
			s = s[1:]
		}

	}
	return res
}

/*
Given the array of integers nums, you will choose two different indices i and j of that array.
Return the maximum value of (nums[i]-1)*(nums[j]-1).
Example 1:
Input: nums = [3,4,5,2]
Output: 12
Explanation: If you choose the indices i=1 and j=2 (indexed from 0), you will get the maximum value,
that is, (nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12.
*/
func MaxProduct(nums []int) int {

	sort.Ints(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}

func MinTimeToVisitAllPoints(points [][]int) int {

	move := 0
	for i := 0; i < len(points)-1; i++ {
		move += twoPoint(points[i], points[i+1])
	}
	return move
}
func twoPoint(pointA []int, pointB []int) int {

	move := 0
	yMove := int(math.Abs(float64(pointA[1]) - float64(pointB[1])))
	xMove := int(math.Abs(float64(pointA[0]) - float64(pointB[0])))
	diagonalMove := getmin(xMove, yMove)
	if diagonalMove == xMove {
		move += yMove
	} else {
		move += xMove
	}
	return move

}
func getmin(a, b int) int {

	if a < b {
		return a
	}
	return b
}

func CountGoodRectangles(rectangles [][]int) int {
	square := []int{}
	maxLan := 0
	for _, v := range rectangles {
		sq := getmin(v[0], v[1])
		if sq > maxLan {
			maxLan = sq
		}
		square = append(square, sq)
	}
	m := make(map[int]int, len(square))
	for _, v := range square {
		m[v]++
	}

	return m[maxLan]
}

func OddCells(m int, n int, indices [][]int) int {

	visit := make([][]bool, m)

	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}

	for _, v := range indices {

		row := v[0]
		rol := v[1]
		for j := 0; j < n; j++ {
			if visit[row][j] {
				visit[row][j] = false
			} else {
				visit[row][j] = true
			}

		}
		for i := 0; i < m; i++ {
			if visit[i][rol] {
				visit[i][rol] = false
			} else {
				visit[i][rol] = true
			}
		}
	}
	count := 0

	for _, v := range visit {
		for _, vv := range v {
			if vv {
				count++
			}
		}
	}
	return count
}
