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
	return res
}
func dfsRestoreIpAddress(s string, res *[]string, end, count int, tmp string, start int) {

	if count == end {
		*res = append(*res, tmp)
	}

	for i := start; i < len(s); i++ {
		if isVaild(tmp + string(s[i])) {
			dfsRestoreIpAddress(s, res, end, count+1, tmp+string(s[i]), start+1)
		}
		dfsRestoreIpAddress(s, res, end, count+1, tmp, start+1)
	}
	return
}
func isVaild(s string) bool {

	if s[0] == '0' {
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
