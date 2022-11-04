package level8

/*
Given two integer arrays, preorder and postorder
where preorder is the preorder traversal of a binary tree of distinct values
and postorder is the postorder traversal of the same tree,
reconstruct and return the binary tree.

If there exist multiple answers, you can return any of them.
preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
*/
func ConstructFromPrePost(preorder []int, postorder []int) *TreeNode {

	node := &TreeNode{Val: preorder[0]}
	if len(preorder) > 1 {
		i := 0
		for i < len(postorder) && postorder[i] != preorder[1] {
			i++
		}
		node.Left = ConstructFromPrePost(preorder[1:i+2], postorder[:i+1])
		if i+2 < len(preorder) {
			node.Right = ConstructFromPrePost(preorder[i+2:], postorder[i+1:len(postorder)-1])
		}
	}

	return node
}

/*
You are given an array arr of positive integers.
You are also given the array queries where queries[i] = [lefti, righti].
For each query i compute the XOR of elements from lefti to righti
(that is, arr[lefti] XOR arr[lefti + 1] XOR ... XOR arr[righti] ).

Return an array answer where answer[i] is the answer to the ith query.
Input: arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
Output: [2,7,14,8]
Explanation:
The binary representation of the elements in the array are:
1 = 0001
3 = 0011
4 = 0100
8 = 1000
The XOR values for queries are:
[0,1] = 1 xor 3 = 2
[1,2] = 3 xor 4 = 7
[0,3] = 1 xor 3 xor 4 xor 8 = 14
[3,3] = 8
*/
func XorQueries(arr []int, queries [][]int) []int {

	answer := []int{}
	for _, query := range queries {

		tmp := 0
		for j := query[0]; j <= query[1]; j++ {
			tmp ^= arr[j]
		}
		answer = append(answer, tmp)
	}
	return answer
}

/*

The bitwise AND of an array nums is the bitwise AND of all integers in nums.
For example, for nums = [1, 5, 3], the bitwise AND is equal to 1 & 5 & 3 = 1.
Also, for nums = [7], the bitwise AND is 7.
You are given an array of positive integers candidates.
Evaluate the bitwise AND of every combination of numbers of candidates.
Each number in candidates may only be used once in each combination.

Return the size of the largest combination of candidates with a bitwise AND greater than 0.
*/

func LargestCombination(candidates []int) int {

	ans := 1
	for i := 0; i <= 24; i++ {
		ct := 0
		for _, v := range candidates {
			if (v>>i)&1 == 1 {
				ct += 1
			}
		}
		ans = getmax(ans, ct)
	}
	return ans
}
func getmax(a, b int) int {

	if a > b {
		return a
	}
	return b
}
