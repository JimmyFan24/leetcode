package level5

import (
	"strconv"
)

func RecoverTree(root *TreeNode) {

	var prev, target1, target2 *TreeNode

	_, target1, target2 = findMistakeNode(root, prev, target1, target2)
	if target2 != nil && target1 != nil {
		target1.Val, target2.Val = target2.Val, target1.Val
	}

}
func findMistakeNode(root, prev, target1, target2 *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {

	if root == nil {
		return prev, target1, target2
	}
	prev, target1, target2 = findMistakeNode(root.Left, prev, target1, target2)
	if prev != nil && prev.Val > root.Val {
		if target1 == nil {
			target1 = prev
		}
		target2 = root

	}
	prev = root
	prev, target1, target2 = findMistakeNode(root.Right, prev, target1, target2)
	return prev, target1, target2
}

/*
Example 1:

Input: s = "25525511135"
Output: ["255.255.11.135","255.255.111.35"]
*/
func RestoreIpAddresses(s string) []string {

	res := []string{}

	dfsRestoreIpAddress(s, &res, 4, 0, "", 0)

	for i := 0; i < len(res); i++ {
		res[i] = res[i][:len(res[i])-1]
	}
	return res
}
func dfsRestoreIpAddress(s string, res *[]string, end, count int, tmp string, start int) {

	if count == end {
		if len(tmp) != len(s)+4 {
			return
		}
		*res = append(*res, tmp)
	}

	//1.1.1.1 11.1.1.1

	j := start
	for i := 1; i <= 3; i++ {
		if j+i <= len(s) {
			t := s[j : j+i]
			tt := tmp
			if isVaild(t) {
				tmp = tmp + t + "."
				dfsRestoreIpAddress(s, res, end, count+1, tmp, j+i)
				tmp = tt
			}
		}
	}
	return
}
func isVaild(s string) bool {

	if s[0] == '0' && len(s) > 1 {
		return false
	}
	if len(s) > 3 || len(s) < 1 {
		return false
	}
	b, _ := strconv.Atoi(s)
	if b > 255 {
		return false
	}
	return true
}
