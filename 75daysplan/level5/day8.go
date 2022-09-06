package level5

/*
Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
*/
func Insert(intervals [][]int, newInterval []int) [][]int {

	//1.完全不相交
	//1.1 右边界 < 左边界 [1,2] [3,4]
	//intervals[i][1] < newInterval[0]
	//1.2 左边界 > 右边界 [3,4] [1,2]
	//intervals[i][0] > newInterval[1]
	//2.包含
	//[1,3] [2,4] ,[1,3] [3,4] ,[3,4] [3,10]
	//[1,4] ,[1,4] [3,10]
	//左边界选最小,右边界选最大
	//eg :[1,2],[3,5],[6,7],[8,10],[12,16] [4,8]
	/*
		1. [1,2] [3,8],[6,7],[8,10],[12,16]
	*/

	res := [][]int{}
	curIndex := 0
	for curIndex < len(intervals) && intervals[curIndex][1] < newInterval[0] {
		res = append(res, intervals[curIndex])
		curIndex++
	}
	for curIndex < len(intervals) && intervals[curIndex][0] <= newInterval[1] {
		start := getSmaller(newInterval[0], intervals[curIndex][0])
		end := getBigger(newInterval[1], intervals[curIndex][1])
		curIndex++
		newInterval = []int{start, end}
	}
	res = append(res, newInterval)
	//插入最右
	for curIndex < len(intervals) {
		res = append(res, intervals[curIndex])
		curIndex++
	}
	return res

}

func getBigger(a, b int) int {
	if a >= b {
		return a
	}
	return b

}
func getSmaller(a, b int) int {
	if a <= b {
		return a
	}
	return b

}

/*
Example 1:

Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.
*/
/*
Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

You may return the answer in any order.
*/
func Combine(n int, k int) [][]int {

	res := [][]int{}
	combineDFS(1, n, k, 0, &res, []int{})
	return res
}
func combineDFS(start, n, k int, count int, res *[][]int, tmp []int) {
	if count == k {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*res = append(*res, t)
		return
	}
	for i := start; i <= n; i++ {
		tmp = append(tmp, i)
		combineDFS(i+1, n, k, count+1, res, tmp)
		tmp = tmp[:len(tmp)-1]
	}
	return
}

/*
You are given the root of a binary tree containing digits from 0 to 9 only.
Each root-to-leaf path in the tree represents a number.
For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.
A leaf node is a node with no children.
*/
func SumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0

	sumDFS(&res, root)
	return res

}
func sumDFS(res *int, root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*res = *res*10 + root.Val
		return
	}
	tmpLeft, tmpRight := 0, 0
	if root.Left != nil {
		tmpLeft = *res*10 + root.Val
	}
	if root.Right != nil {
		tmpRight = *res*10 + root.Val
	}
	sumDFS(&tmpLeft, root.Left)
	sumDFS(&tmpRight, root.Right)
	*res = tmpRight + tmpLeft
	return

}

func PathSum(root *TreeNode, targetSum int) [][]int {

	res := [][]int{}
	if root == nil {
		return nil
	}
	dfsPathSum(&res, targetSum, root, []int{})
	return res
}
func dfsPathSum(res *[][]int, targetSum int, root *TreeNode, tmp []int) {

	if root.Left == nil && root.Right == nil {
		if root.Val-targetSum == 0 {
			tmp = append(tmp, root.Val)
			t := make([]int, len(tmp))
			copy(t, tmp)
			*res = append(*res, t)
			return
		} else {
			return
		}
	}

	tmp = append(tmp, root.Val)
	if root.Left != nil {
		dfsPathSum(res, targetSum-root.Val, root.Left, tmp)
	}

	if root.Right != nil {
		dfsPathSum(res, targetSum-root.Val, root.Right, tmp)
	}

	return
}

func Partition(head *ListNode, x int) *ListNode {

	biggerList := &ListNode{}
	smallerList := &ListNode{}
	tmpbigger := biggerList
	tmpsmaller := smallerList
	tmpHead := &ListNode{}
	res := tmpHead
	cur := head

	for cur != nil {
		if cur.Val >= x {
			biggerList.Next = &ListNode{
				cur.Val,
				nil,
			}
			biggerList = biggerList.Next
		}
		if cur.Val < x {
			smallerList.Next = &ListNode{
				cur.Val,
				nil,
			}
			smallerList = smallerList.Next
		}
		cur = cur.Next
	}

	tmpsmaller = tmpsmaller.Next
	for tmpsmaller != nil {
		tmpHead.Next = &ListNode{
			tmpsmaller.Val, nil,
		}
		tmpsmaller = tmpsmaller.Next
		tmpHead = tmpHead.Next
	}
	tmpbigger = tmpbigger.Next
	for tmpbigger != nil {
		tmpHead.Next = &ListNode{
			tmpbigger.Val, nil,
		}
		tmpbigger = tmpbigger.Next
		tmpHead = tmpHead.Next
	}
	return res.Next
}
