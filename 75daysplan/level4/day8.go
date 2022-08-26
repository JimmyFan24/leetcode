package level4

import "math"

type NewNode struct {
	Val      int
	Children []*NewNode
}

func LevelOrder(root *NewNode) [][]int {

	if root == nil {
		return nil
	}
	res := [][]int{}

	queue := []*NewNode{root}

	for len(queue) > 0 {

		l := len(queue)
		tmp := []int{}
		for i := 0; i < len(queue); i++ {

			for _, v := range queue[i].Children {
				queue = append(queue, v)
			}

			tmp = append(tmp, queue[i].Val)

		}
		queue = queue[l:]
		res = append(res, tmp)
	}
	return res

}
func FindSolution(customFunction func(int, int) int, z int) [][]int {

	res := [][]int{}

	for x := 1; x < 1001; x++ {
		y := 1000
		for y > 1 && customFunction(x, y) > z {
			y--
		}
		if customFunction(x, y) == z {
			res = append(res, []int{x, y})
		}
	}
	return res
}

func DelNodes(root *TreeNode, to_delete []int) []*TreeNode {

	res := []*TreeNode{}
	tDelete := make(map[int]bool, len(to_delete))
	for _, v := range to_delete {
		tDelete[v] = true
	}
	deleteNode(root, true, &res, tDelete)
	return res
}
func deleteNode(root *TreeNode, isRoot bool, res *[]*TreeNode, to_delete map[int]bool) bool {
	if root == nil {
		return false
	}

	if isRoot && !to_delete[root.Val] {
		(*res) = append((*res), root)
	}
	isRoot = false
	if to_delete[root.Val] {
		isRoot = true
	}
	if deleteNode(root.Left, isRoot, res, to_delete) {
		root.Left = nil
	}
	if deleteNode(root.Right, isRoot, res, to_delete) {
		root.Right = nil
	}

	return isRoot

}

func GetMaximumXor(nums []int, maximumBit int) []int {

	answer := []int{}
	max := 0
	for _, v := range nums {
		max ^= v
	}

	maxNum := int(math.Pow(2, float64(maximumBit))) - 1

	for len(nums) > 0 {
		answer = append(answer, maxNum^max)
		max ^= nums[len(nums)-1]
		nums = nums[:len(nums)-1]
	}
	return answer

}
func getTheMaxOne(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func MinAddToMakeValid(s string) int {
	left := []rune{}
	right := []rune{}

	//(()
	for _, v := range s {
		if v == '(' {
			left = append(left, v)
			continue
		}
		if v == ')' {

			if len(left) == 0 {
				right = append(right, v)
			} else {
				left = left[:len(left)-1]
			}

		}
	}
	return len(left) + len(right)
}
