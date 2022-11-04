package level8

/*
You are given a 0-indexed binary string target of length n.
You have another binary string s of length n that is initially set to all zeros.
You want to make s equal to target.
In one operation, you can pick an index i where 0 <= i < n and flip all bits in the inclusive range [i, n - 1].
Flip means changing '0' to '1' and '1' to '0'.

Return the minimum number of operations needed to make s equal to target.

Input: target = "10111"
Output: 3
Explanation: Initially, s = "00000".
Choose index i = 2: "00000" -> "00111"
Choose index i = 0: "00111" -> "11000"
Choose index i = 1: "11000" -> "10111"
We need at least 3 flip operations to form target.
*/
func MinFlips(target string) int {

	count := 0
	c := '0'
	for _, v := range target {
		if v != c {
			count++
			c = v
		}
	}

	return count
}

/*
You are given a 2D integer array descriptions
where descriptions[i] = [parenti, childi, isLefti]
indicates that parenti is the parent of childi in a binary tree of unique values. Furthermore,

If isLefti == 1, then childi is the left child of parenti.
If isLefti == 0, then childi is the right child of parenti.
Construct the binary tree described by descriptions and return its root.

The test cases will be generated such that the binary tree is valid.
[[85,82,1],[74,85,1],[39,70,0],[82,38,1],[74,39,0],[39,13,1]]
*/
func CreateBinaryTree(descriptions [][]int) *TreeNode {

	parentGroup := make(map[int]*TreeNode, len(descriptions))
	childGroup := make(map[int]*TreeNode, len(descriptions))
	root := &TreeNode{}
	//1.先得出所有的parent节点
	for _, v := range descriptions {
		parentGroup[v[0]] = &TreeNode{Val: v[0]}
	}
	for _, node := range descriptions {
		childi := &TreeNode{Val: node[1]}
		if node[2] == 1 {
			if childGroup[node[1]] != nil {
				parentGroup[node[0]].Left = childGroup[node[1]]
			} else if parentGroup[node[1]] != nil {
				parentGroup[node[0]].Left = parentGroup[node[1]]
				childGroup[node[1]] = parentGroup[node[1]]
			} else {
				parentGroup[node[0]].Left = childi
				childGroup[node[1]] = childi
			}

		} else {
			if childGroup[node[1]] != nil {
				parentGroup[node[0]].Right = childGroup[node[1]]
			} else if parentGroup[node[1]] != nil {
				parentGroup[node[0]].Right = parentGroup[node[1]]
				childGroup[node[1]] = parentGroup[node[1]]
			} else {
				parentGroup[node[0]].Right = childi
				childGroup[node[1]] = childi
			}

		}

	}

	for _, p := range parentGroup {
		if _, ok := childGroup[p.Val]; !ok {

			root = p
			break
		}
	}
	return root
}

/*
A happy string is a string that:
consists only of letters of the set ['a', 'b', 'c'].
s[i] != s[i + 1] for all values of i from 1 to s.length - 1 (string is 1-indexed).
For example, strings "abc", "ac", "b" and "abcbabcbcb" are all happy strings
and strings "aa", "baa" and "ababbc" are not happy strings.

Given two integers n and k, consider a list of all happy strings of length n sorted in lexicographical order.

Return the kth string of this list or return an empty string if there are less than k happy strings of length n.

Input: n = 3, k = 9
Output: "cab"
Explanation: There are 12 different happy string of length 3
["aba", "abc", "aca", "acb", "bab", "bac", "bca", "bcb", "cab", "cac", "cba", "cbc"].
You will find the 9th string = "cab"
*/
func GetHappyString(n int, k int) string {

	happyStringGroup := getHappyString(n)
	return happyStringGroup[k-1]
}

var happy = []byte{'a', 'b', 'c'}

func getHappyString(length int) []string {

	//1.a,b,c
	//2.s[i] != s[i + 1]
	res := []string{}
	bfsHappyString(&res, length, 0, "")
	return res

}
func bfsHappyString(res *[]string, length int, count int, tmp string) {

	if count == length {
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(happy); i++ {
		if len(tmp) != 0 {
			last := tmp[len(tmp)-1]
			if happy[i] != last {
				tmp += string(happy[i])
				bfsHappyString(res, length, count+1, tmp)
				tmp = tmp[:len(tmp)-1]
			}
		} else {
			tmp += string(happy[i])
			bfsHappyString(res, length, count+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}

	}
}
