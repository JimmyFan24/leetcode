package level8

import "math"

/*
Input: target = [1,3], n = 3
Output: ["Push","Push","Pop","Push"]
Explanation: Initially the stack s is empty. The last element is the top of the stack.
Read 1 from the stream and push it to the stack. s = [1].
Read 2 from the stream and push it to the stack. s = [1,2].
Pop the integer on the top of the stack. s = [1].
Read 3 from the stream and push it to the stack. s = [1,3].
*/
func BuildArray(target []int, n int) []string {

	res := []string{}

	m := make(map[int]bool, len(target))
	for i := 0; i < len(target); i++ {
		m[target[i]] = true
	}

	targetIndex := 0
	for i := 1; i <= n; i++ {

		if m[i] {
			res = append(res, "Push")
			targetIndex++
		} else {
			res = append(res, "Push")
			res = append(res, "Pop")

		}
		if targetIndex == len(target) {
			break
		}

	}
	return res

}

/*
You are given the root of a binary tree with n nodes
where each node in the tree has node.val coins.
There are n coins in total throughout the whole tree.
In one move, we may choose two adjacent nodes and
move one coin from one node to another.
A move may be from parent to child, or from child to parent.
Return the minimum number of moves required to make every node have exactly one coin.
*/
func DistributeCoins(root *TreeNode) int {

	moveCount := 0
	distributeCoins(root, &moveCount)

	return moveCount
}
func distributeCoins(root *TreeNode, count *int) int {

	if root == nil {
		return 0
	}
	left, right := distributeCoins(root.Left, count), distributeCoins(root.Right, count)
	*count += int(math.Abs(float64(left)) + math.Abs(float64(right)))

	return left + right + root.Val - 1
}
