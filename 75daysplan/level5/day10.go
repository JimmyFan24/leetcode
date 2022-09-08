package level5

import (
	"sort"
	"strings"
)

/*
Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   -1
  2 3
 1 -1 -3
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
*/
func MinimumTotal(triangle [][]int) int {

	for i := 1; i < len(triangle); i++ {

		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][0]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += getMinTriangle(triangle[i-1][j], triangle[i-1][j-1])
			}

		}
	}
	sort.Ints(triangle[len(triangle)-1])
	return triangle[len(triangle)-1][0]
}
func getMinTriangle(a, b int) int {

	if a < b {
		return a
	}
	return b
}
func SortedListToBST(head *ListNode) *TreeNode {

	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return buildBST(arr)
}
func buildBST(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	mid := len(arr) / 2

	head := &TreeNode{}
	head.Val = arr[mid]
	left := buildBST(arr[:mid])
	right := buildBST(arr[mid+1:])
	head.Left = left
	head.Right = right
	return head
}

/*
Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]
*/
func ReverseBetween(head *ListNode, left int, right int) *ListNode {

	cur := head
	newList := &ListNode{}
	curNewList := newList
	count := 1
	var leftPrev *ListNode
	var rightPrev *ListNode
	for cur != nil {
		if count+1 == left {
			leftPrev = cur
		}
		if count-1 == right {
			rightPrev = cur
		}
		if count >= left && count <= right {
			curNewList.Next = &ListNode{
				Val:  cur.Val,
				Next: nil,
			}
			curNewList = curNewList.Next
		}
		count++
		cur = cur.Next
	}
	l := reverseList(newList.Next)
	if leftPrev != nil {
		leftPrev.Next = l
	} else {
		head = l
	}
	for l.Next != nil {
		l = l.Next
	}
	//1--2--3--4--5--6
	l.Next = rightPrev
	return head
}
func reverseList(node *ListNode) *ListNode {

	var newHead *ListNode
	prev := newHead
	cur := node
	for cur != nil {
		//newhead<--1--2--3--4
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

/*
Example 1:

Input: s = "the sky is blue"
Output: "blue is sky the"
*/
func ReverseWords(s string) string {
	list := strings.Split(s, " ")
	realList := []string{}
	for i := 0; i < len(list); i++ {
		if len(list[i]) != 0 {
			realList = append(realList, list[i])
		}
	}
	for i := 0; i < len(realList)/2; i++ {
		realList[i], realList[len(realList)-i-1] = realList[len(realList)-i-1], realList[i]
	}
	res := ""
	for _, v := range realList {
		res += v + " "
	}
	return res[:len(res)-1]
}

/*
Given a string s and a dictionary of strings wordDict,
return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.


Example 1:

Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
*/
func WordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(wordDict); j++ {
			if i >= len(wordDict[j]) && dp[i-len(wordDict[j])] && wordDict[j] == s[i-len(wordDict[j]):i] {
				dp[i] = true
			}
		}
	}
	return dp[len(dp)-1]
}
