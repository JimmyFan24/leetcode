package level5

import (
	"math"
)

func SubtreeWithAllDeepest(root *TreeNode) *TreeNode {

	if root.Left == nil && root.Right == nil {
		return root
	}

	count := 0
	r := getDepth(root, count)
	if r == 2 && root.Right != nil && root.Left != nil {
		return root
	}

	left := getDepth(root.Left, count)
	right := getDepth(root.Right, count)
	if left > right {
		return SubtreeWithAllDeepest(root.Left)
	} else if right > left {
		return SubtreeWithAllDeepest(root.Right)
	}
	return root

}
func getDepth(node *TreeNode, count int) int {

	if node == nil {
		return 0
	}

	left := getDepth(node.Left, count)
	right := getDepth(node.Right, count)
	count = getMaxDepth(left, right)
	return count + 1
}
func getMaxDepth(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MctFromLeafValues(arr []int) int {

	sum := 0
	for len(arr) > 1 {
		min := math.MaxInt32
		index := -1
		for i := 0; i < len(arr); i++ {
			if arr[i] < min {
				min = arr[i]
				index = i
			}
		}
		smaller := -1
		if index+1 < len(arr) && index-1 >= 0 {
			if arr[index-1] < arr[index+1] {
				smaller = arr[index-1]
			} else {
				smaller = arr[index+1]
			}
		} else if index+1 >= len(arr) && index-1 >= 0 {
			smaller = arr[index-1]
		} else if index+1 < len(arr) && index-1 < 0 {
			smaller = arr[index+1]
		}
		nonLeafNode := smaller * min
		sum += nonLeafNode
		tmp := []int{}
		for i, v := range arr {
			if i != index {
				tmp = append(tmp, v)
			}
		}
		arr = arr[:len(arr)-1]

		copy(arr, tmp)
	}
	return sum
}
