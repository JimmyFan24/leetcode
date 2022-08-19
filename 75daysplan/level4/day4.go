package level4

import (
	"math"
)

/*
Example 1:

Input: n = 6
Output: 2
Explanation: arr = [1,3,5,7,9,11]
First operation choose x = 2 and y = 0, this leads arr to be [2, 3, 4]
In the second operation choose x = 2 and y = 0 again, thus arr = [3, 3, 3].
*/
func MinOperations(n int) int {

	sum := 0

	for i := 0; i < n/2; i++ {

		sum += n - (2*i + 1)
	}
	return sum

}
func FindingUsersActiveMinutes(logs [][]int, k int) []int {

	answer := make([]int, k)
	userMap := make(map[int]map[int]int, len(logs))

	for _, v := range logs {

		if _, ok := userMap[v[0]]; !ok {
			minMap := make(map[int]int, len(logs))
			minMap[v[1]]++
			userMap[v[0]] = minMap

		} else {
			userMap[v[0]][v[1]]++
		}
	}

	uam := make(map[int]int, len(userMap))
	for _, v := range userMap {
		uam[len(v)]++
	}
	for k, v := range uam {
		answer[k-1] = v
	}
	return answer

}
func BalanceBST(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	//中序遍历
	arr := inOrder(root)

	return buildBST(arr, 0, len(arr)-1)
}
func buildBST(arr []*TreeNode, l, r int) *TreeNode {

	if l > r {
		return nil
	}
	mid := l + ((r - l) >> 1)
	root := arr[mid]
	root.Left = buildBST(arr, l, mid-1)
	root.Right = buildBST(arr, mid+1, r)
	return root
}
func inOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	res := []*TreeNode{}
	left := inOrder(root.Left)
	for _, v := range left {
		v.Left = nil
		v.Right = nil
		res = append(res, v)
	}
	res = append(res, root)
	right := inOrder(root.Right)
	for _, v := range right {
		v.Left = nil
		v.Right = nil
		res = append(res, v)
	}
	return res
}
func insertRightNode(root, node *TreeNode) {

	for root.Right != nil && node.Val > root.Right.Val {
		root = root.Right
	}
	root.Right = node
	node.Left = nil
	node.Right = nil

}
func insertLeftNode(root, node *TreeNode) {

	for root.Left != nil && node.Val < root.Left.Val {
		root = root.Left
	}
	root.Left = node
	node.Left = nil
	node.Right = nil

}
func isBalanceBST(left, right *TreeNode) bool {
	return int(math.Abs(float64(getDepth(left))-float64(getDepth(right)))) <= 1

}
func getDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return max(getDepth(root.Left), getDepth(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CountCharacters(words []string, chars string) int {

	m := make(map[rune]int, 26)
	for _, v := range chars {
		m[v]++
	}
	res := 0
	for _, word := range words {
		tmp := make(map[rune]int, len(m))
		for k, v := range m {
			tmp[k] = v
		}
		if isGoodWord(word, tmp) && len(word) <= len(chars) {
			res += len(word)
		}
	}
	return res
}
func isGoodWord(word string, m map[rune]int) bool {

	for _, v := range word {
		if m[v] <= 0 {
			return false
		} else {
			m[v]--
		}
	}
	return true
}

func NearestValidPoint(x int, y int, points [][]int) int {

	//1.find the usefulpoint
	useful := [][]int{}
	for _, point := range points {
		if point[0] == x || point[1] == y {
			useful = append(useful, point)
		}
	}
	//2.get manhadun distance
	minDistance := math.MaxInt32
	minPoint := []int{}
	for _, use := range useful {
		tmp := int(math.Abs(float64(use[0])-float64(x))) + int(math.Abs(float64(use[1])-float64(y)))
		if tmp < minDistance {

			minDistance = tmp
			minPoint = use
		} else if tmp == minDistance {
			continue
		}
	}
	res := 0
	if len(minPoint) == 0 {
		return -1
	}
	for i, v := range points {
		if v[0] == minPoint[0] && v[1] == minPoint[1] {
			res = i
			break
		}
	}
	return res
}
