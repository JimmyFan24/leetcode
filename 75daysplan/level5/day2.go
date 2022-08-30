package level5

import (
	"strings"
)

func GarbageCollection(garbage []string, travel []int) int {

	travelCost := make(map[rune]int, len(travel))
	hourseCost := make(map[rune]int, len(travel))

	hourse := 0

	getTravelSum := make([]int, len(travel))
	getTravelSum[0] = travel[0]
	for j := 1; j < len(travel); j++ {
		getTravelSum[j] = getTravelSum[j-1] + travel[j]
	}
	for i, v := range garbage {

		if i > 0 {

			if strings.Contains(v, string('P')) {
				travelCost['P'] = getTravelSum[hourse]
			}
			if strings.Contains(v, string('M')) {

				travelCost['M'] = getTravelSum[hourse]
			}
			if strings.Contains(v, string('G')) {
				travelCost['G'] = getTravelSum[hourse]
			}
			hourse++
		}
		for _, vv := range v {
			switch vv {
			case 'G':
				hourseCost['G']++
			case 'P':
				hourseCost['P']++
			case 'M':
				hourseCost['M']++
			default:

			}
		}

	}
	sum := 0
	for _, v := range hourseCost {
		sum += v
	}
	for _, v := range travelCost {
		sum += v
	}
	return sum
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	res := list1
	count := 0
	start := list1
	for a > count {
		start = list1
		list1 = list1.Next
		count++
	}
	for b > count {
		list1 = list1.Next
		count++
	}
	//0,1,2,3,4,5
	bList := list1
	start.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = bList.Next
	return res
}

func CountSquares(matrix [][]int) int {

	m := len(matrix)
	n := len(matrix[0])
	sum := 0
	side := m
	if m > n {
		side = n
	}
	for i := 1; i <= side; i++ {
		sum += countOne(i, matrix)
	}
	return sum
}

/*
1,1,1
1,1,1
1,1,1
*/
func countOne(side int, matrix [][]int) int {

	//i+side <len(row)
	count := 0
	for i := 0; i <= len(matrix)-side; i++ {
		for j := 0; j <= len(matrix[0])-side; j++ {
			if isSquare(matrix, i, j, side) {
				count++
			}
		}
	}
	return count
}
func isSquare(matrix [][]int, x, y int, side int) bool {

	for i := x; i < side+x; i++ {
		for j := y; j < side+y; j++ {
			if matrix[i][j] != 1 {
				return false
			}
		}
	}
	return true
}

/*
Given two strings s and part, perform the following operation on s until all occurrences of the substring part are removed:

Find the leftmost occurrence of the substring part and remove it from s.
Return s after removing all occurrences of part.

A substring is a contiguous sequence of characters in a string.
*/
/*
Example 1:

Input: s = "daabcbaabcbc", part = "abc"
Output: "dab"
Explanation: The following operations are done:
- s = "daabcbaabcbc", remove "abc" starting at index 2, so s = "dabaabcbc".
- s = "dabaabcbc", remove "abc" starting at index 4, so s = "dababc".
- s = "dababc", remove "abc" starting at index 3, so s = "dab".
Now s has no occurrences of "abc".
*/
func RemoveOccurrences(s string, part string) string {

	for i := 0; i < len(s); {
		index := strings.Index(s, part)

		if index >= 0 && index < len(s) {
			s = s[:index] + s[index+len(part):]
			i = index
		} else {
			break
		}

	}
	return s
}
func LetterCasePermutation(s string) []string {

	count := 0
	for i, _ := range s {
		if isLetter(s, i) {

			count++
		}
	}
	s = strings.ToLower(s)
	res := []string{}
	res = append(res, s)
	for i := 1; i <= count; i++ {
		dfsLowerLetter(s, 0, i, 0, &res)

	}
	return res
}
func dfsLowerLetter(s string, start int, end int, count int, res *[]string) {
	if end == count {
		*res = append(*res, s)
		return
	}

	ss := s
	for j := start; j < len(ss); j++ {
		if isLetter(ss, j) {
			ss = ss[:j] + strings.ToUpper(string(ss[j])) + ss[j+1:]
		} else {
			continue
		}
		dfsLowerLetter(ss, j+1, end, count+1, res)
		ss = s
	}
	return
}
func isLetter(s string, i int) bool {

	return (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z')
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RemoveLeafNodes(root *TreeNode, target int) *TreeNode {

	if root == nil {
		return nil
	}

	if removeTarget(root.Left, target) != nil {
		root.Left = nil
	}
	if removeTarget(root.Right, target) != nil {
		root.Right = nil
	}
	if root.Val == target && root.Left == nil && root.Right == nil {
		root = nil
	}
	return root

}
func removeTarget(node *TreeNode, target int) *TreeNode {

	if node == nil {
		return nil
	}
	if node.Left != nil {
		if removeTarget(node.Left, target) != nil {
			node.Left = nil
		}
	}
	if node.Right != nil {
		if removeTarget(node.Right, target) != nil {
			node.Right = nil
		}
	}
	if node.Val == target && node.Left == nil && node.Right == nil {
		return node
	}

	return nil
}
