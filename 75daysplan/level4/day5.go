package level4

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
Example 1:

Input: s = "leetcode", t = "coats"
Output: 7
Explanation:
- In 2 steps, we can append the letters in "as" onto s = "leetcode", forming s = "leetcodeas".
- In 5 steps, we can append the letters in "leede" onto t = "coats", forming t = "coatsleede".
"leetcodeas" and "coatsleede" are now anagrams of each other.
We used a total of 2 + 5 = 7 steps.
It can be shown that there is no way to make them anagrams of each other with less than 7 steps
*/
func MinSteps(s string, t string) int {

	m := make(map[rune]int, 26)

	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]++
	}
	count := 0
	sM := make(map[rune]int, 26)
	tM := make(map[rune]int, 26)
	for _, v := range s {
		sM[v]++
	}
	for _, v := range t {
		tM[v]++
	}
	for k, _ := range m {
		if sM[k] == tM[k] && sM[k] != 0 {
			continue
		} else if sM[k] != tM[k] {
			count += int(math.Abs(float64(sM[k]) - float64(tM[k])))
		}
	}
	return count
}

/*
A complex number can be represented as a string on the form "real+imaginaryi" where:

real is the real part and is an integer in the range [-100, 100].
imaginary is the imaginary part and is an integer in the range [-100, 100].
i2 == -1.
Given two complex numbers num1 and num2 as strings, return a string of the complex number that represents their multiplications.
Example 1:

Input: num1 = "1+1i", num2 = "1+1i"
Output: "0+2i"
Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.
*/

/*
"1+-1i"
"1+-1i"
*/
func ComplexNumberMultiply(num1 string, num2 string) string {

	num1Str := strings.Split(num1, "+")
	num2Str := strings.Split(num2, "+")

	num1Real, _ := strconv.Atoi(num1Str[0])
	num2Real, _ := strconv.Atoi(num2Str[0])
	real := num1Real * num2Real

	newNum1 := strings.Replace(num1Str[1], "i", "", -1)
	newNum2 := strings.Replace(num2Str[1], "i", "", -1)
	imageNum1, _ := strconv.Atoi(newNum1)
	imageNum2, _ := strconv.Atoi(newNum2)
	imaginary := num1Real*imageNum2 + imageNum1*num2Real
	real = real - imageNum2*imageNum1
	return strconv.Itoa(real) + "+" + strconv.Itoa(imaginary) + "i"

}

/*
ib.end < ia.start，此时我们应该继续判断B的下一个区间
ib.start > ia.end，此时我们应该继续判断A的下一个区间
ia.start > ib.start and ia.end < ib.end，此时ib包含ia，那么交集就是ia。
ia.start < ib.start and ia.end > ib.end，此时ia包含ib，那么交集就是ib。
剩下就是ia和ib相交的情况，那么我们取max(ia.start, ib.start), min(ia.end, ib.end)即可。
*/
func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {

	i, j := 0, 0
	res := [][]int{}
	for i < len(firstList) && j < len(secondList) {
		if firstList[i][len(firstList[i])-1] < secondList[j][0] {
			i++
			continue
		}
		if firstList[i][0] > secondList[j][len(secondList[j])-1] {
			j++
			continue
		}
		//ia.start > ib.start and ia.end < ib.end
		if firstList[i][0] > secondList[j][0] && firstList[i][len(firstList[i])-1] < secondList[j][len(secondList[j])-1] {

			res = append(res, firstList[i])
			i++
		}
		//ia.start < ib.start and ia.end > ib.end
		if firstList[i][0] < secondList[j][0] && firstList[i][len(firstList[i])-1] > secondList[j][len(secondList[j])-1] {

			res = append(res, secondList[j])
			j++
		}
		//17--20
		//6--8
		//剩下就是ia和ib相交的情况，那么我们取max(ia.start, ib.start), min(ia.end, ib.end)即可。
		res = append(res, []int{max(firstList[i][0], secondList[j][0]), min(firstList[i][len(firstList[i])-1], secondList[j][len(secondList[j])-1])})
		if firstList[i][len(firstList[i])-1] <= secondList[j][len(secondList[j])-1] {
			i++
		} else if firstList[i][len(firstList[i])-1] >= secondList[j][len(secondList[j])-1] {
			j++
		}
	}
	return res
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/*
You should rearrange the elements of nums such that the modified array follows the given conditions:

Every consecutive pair of integers have opposite signs.
For all integers with the same sign, the order in which they were present in nums is preserved.
The rearranged array begins with a positive integer.
Return the modified array after rearranging the elements to satisfy the aforementioned conditions
Example 1:

Input: nums = [3,1,-2,-5,2,-4]
Output: [3,-2,1,-5,2,-4]
Explanation:
The positive integers in nums are [3,1,2]. The negative integers are [-2,-5,-4].
The only possible way to rearrange them such that they satisfy all conditions is [3,-2,1,-5,2,-4].
Other ways such as [1,-2,2,-5,3,-4], [3,1,2,-2,-5,-4], [-2,3,-5,1,-4,2] are incorrect because they do not satisfy one or more conditions.
*/
func RearrangeArray(nums []int) []int {

	//1.正数数组
	positive := []int{}
	//2.负数数组
	negative := []int{}
	for _, v := range nums {
		if v > 0 {
			positive = append(positive, v)
		} else {
			negative = append(negative, v)
		}
	}
	res := []int{}
	i, j := 0, 0
	for i < len(positive) && j < len(negative) {

		res = append(res, positive[i])
		res = append(res, negative[j])
		i++
		j++

	}
	return res

}

/*
"011001"
"000000"
"010100"
"001000"
*/
func NumberOfBeams(bank []string) int {

	count := 0

	for i := 0; i < len(bank); i++ {
		j := i + 1

		for ; j < len(bank); j++ {
			row1 := bank[i]
			row2 := bank[j]
			countRow1, countRow2 := 0, 0
			for _, v := range row1 {
				if v == '1' {
					countRow1++
				}
			}
			for _, v := range row2 {
				if v == '1' {
					countRow2++
				}
			}
			if countRow2 != 0 && countRow1 != 0 {

				count += countRow2 * countRow1
				i = j - 1
				break
			}

		}

	}
	return count
}

/*
Given an integer array arr, partition the array into (contiguous) subarrays of length at most k.
After partitioning, each subarray has their values changed to become the maximum value of that subarray.

Return the largest sum of the given array after partitioning.
Test cases are generated so that the answer fits in a 32-bit integer.

Example 1:

Input: arr = [1,15,7,9,2,5,10], k = 3
Output: 84
Explanation: arr becomes [15,15,15,9,10,10,10]
*/
func MaxSumAfterPartitioning(arr []int, k int) int {

	dp := make([]int, len(arr)+1)

	for i := 0; i <= len(arr); i++ {
		j := i - 1
		curmax := 0
		//[1,15,7,9,2,5,10], k = 3
		/*
			dp[i] = dp[i-1]+1*arr[i] == 2 + 15 = 17
			dp[i] = dp[i-2]+2*max(arr[i],arr[i-1]) == 1 + 15*2 = 31
			dp[i] = dp[i-3]+3*max(arr[i],arr[i-1],arr[i-2]) ==
		*/
		//其中，i-j<=k,而MAX则为arr[j]到arr[i-1]之间的最大值。
		for ; i-j <= k && j >= 0; j-- {
			curmax = mmax(curmax, arr[j])
			dp[i] = mmax(dp[i], dp[j]+curmax*(i-j))
		}

	}
	return dp[len(arr)]

}
func mmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
There are n people that are split into some unknown number of groups.
Each person is labeled with a unique ID from 0 to n - 1.

You are given an integer array groupSizes, where groupSizes[i] is the size of the group that person i is in.
For example, if groupSizes[1] = 3, then person 1 must be in a group of size 3.

Return a list of groups such that each person i is in a group of size groupSizes[i].

Each person should appear in exactly one group,
and every person must be in a group. If there are multiple answers, return any of them.
It is guaranteed that there will be at least one valid solution for the given input.
*/
/*
Example 1:

Input: groupSizes = [3,3,3,3,3,1,3]
Output: [[5],[0,1,2],[3,4,6]]
Explanation:
The first group is [5]. The size is 1, and groupSizes[5] = 1.
The second group is [0,1,2]. The size is 3, and groupSizes[0] = groupSizes[1] = groupSizes[2] = 3.
The third group is [3,4,6]. The size is 3, and groupSizes[3] = groupSizes[4] = groupSizes[6] = 3.
Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].
*/
func GroupThePeople(groupSizes []int) [][]int {

	res := [][]int{}
	m := make(map[int][]int, len(groupSizes))
	for i, v := range groupSizes {

		if len(m[v]) < v {
			m[v] = append(m[v], i)
		}
		if len(m[v]) == v {
			res = append(res, m[v])
			m[v] = []int{}
		}
	}
	/*
		3-->6
		1-->1
	*/

	return res
}

/*
输入：queries = [3,1,2,1], m = 5
输出：[2,1,2,1]
解释：待查数组 queries 处理如下：
对于 i=0: queries[i]=3, P=[1,2,3,4,5], 3 在 P 中的位置是 2，接着我们把 3 移动到 P 的起始位置，得到 P=[3,1,2,4,5] 。
对于 i=1: queries[i]=1, P=[3,1,2,4,5], 1 在 P 中的位置是 1，接着我们把 1 移动到 P 的起始位置，得到 P=[1,3,2,4,5] 。
对于 i=2: queries[i]=2, P=[1,3,2,4,5], 2 在 P 中的位置是 2，接着我们把 2 移动到 P 的起始位置，得到 P=[2,1,3,4,5] 。
对于 i=3: queries[i]=1, P=[2,1,3,4,5], 1 在 P 中的位置是 1，接着我们把 1 移动到 P 的起始位置，得到 P=[1,2,3,4,5] 。
因此，返回的结果数组为 [2,1,2,1] 。
*/
type Node struct {
	Val  int
	Next *Node
}

func ProcessQueries(queries []int, m int) []int {

	head := &Node{}
	tmpHead := head
	for i := 1; i <= m; i++ {
		tmpHead.Next = &Node{
			Val:  i,
			Next: nil,
		}
		tmpHead = tmpHead.Next
	}

	res := []int{}
	pre := head
	cur := head.Next
	count := 0
	for _, v := range queries {
		for cur.Next != nil && cur.Val != v {
			pre = cur
			cur = cur.Next
			count++
		}
		res = append(res, count)
		if count == 0 {
			continue
		}
		//找到节点,移到最前
		pre.Next = cur.Next
		tmp := head.Next
		cur.Next = tmp
		head.Next = cur
		count = 0

	}

	return res
}

/*
Input: nums = [3,5,2,3]
Output: 7
Explanation: The elements can be paired up into pairs (3,3) and (5,2).
The maximum pair sum is max(3+3, 5+2) = max(6, 7) = 7.
*/
func MinPairSum(nums []int) int {

	sort.Ints(nums)

	max := math.MinInt32
	for i := 0; i < len(nums)/2; i++ {
		if nums[i]+nums[len(nums)-1-i] > max {
			max = nums[i] + nums[len(nums)-1-i]
		}
	}
	return max
}

/*
In a linked list of size n, where n is even,
the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node,
if 0 <= i <= (n / 2) - 1.

For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only nodes with twins for n = 4.
The twin sum is defined as the sum of a node and its twin.

Given the head of a linked list with even length, return the maximum twin sum of the linked list.

Input: head = [4,2,2,3]
Output: 7
Explanation:
The nodes with twins present in this linked list are:
- Node 0 is the twin of node 3 having a twin sum of 4 + 3 = 7.
- Node 1 is the twin of node 2 having a twin sum of 2 + 2 = 4.
Thus, the maximum twin sum of the linked list is max(7, 4) = 7.
*/
func PairSum(head *ListNode) int {

	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	max := math.MinInt32
	for i := 0; i < len(list)/2; i++ {
		if list[i]+list[len(list)-1-i] > max {
			max = list[i] + list[len(list)-1-i]
		}
	}
	return max
}
func WateringPlants(plants []int, capacity int) int {

	step := 0
	//2,2,3,5
	tmpcapacity := capacity
	for i, v := range plants {
		if tmpcapacity < v {
			step = step + 2*i + 1
			tmpcapacity = capacity - v
		} else {
			step++
			tmpcapacity -= v
		}
	}
	return step
}

/*
You are given the root node of a binary search tree (BST) and a value to insert into the tree.
Return the root node of the BST after the insertion.
It is guaranteed that the new value does not exist in the original BST.

Notice that there may exist multiple valid ways for the insertion,
as long as the tree remains a BST after insertion.
You can return any of them.
*/
func InsertIntoBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			return root

		}
		left := InsertIntoBST(root.Left, val)
		root.Left = left
		return root
	}
	if root.Right == nil {
		root.Right = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
		return root
	}
	right := InsertIntoBST(root.Right, val)
	root.Right = right
	return root
}

/*
Given the root of a binary tree,
return the same tree where every subtree (of the given tree) not containing a 1 has been removed.

A subtree of a node node is node plus every node that is a descendant of node.
*/
func PruneTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	left := PruneTree(root.Left)
	right := PruneTree(root.Right)
	if root.Val == 0 && left == nil && right == nil {
		return nil
	}
	root.Left = left
	root.Right = right
	return root

}

/*
给你一个二维整数数组 descriptions ，其中 descriptions[i] = [parenti, childi, isLefti] 表示 parenti 是 childi 在 二叉树 中的 父节点，二叉树中各节点的值 互不相同 。此外：

如果 isLefti == 1 ，那么 childi 就是 parenti 的左子节点。
如果 isLefti == 0 ，那么 childi 就是 parenti 的右子节点。
*/
//[[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]
func CreateBinaryTree(descriptions [][]int) *TreeNode {

	head := &TreeNode{}
	res := head
	queue := []*TreeNode{}

	for _, v := range descriptions {

		parent := &TreeNode{}
		if len(queue) == 0 {
			parent = &TreeNode{Val: v[0]}
			queue = append(queue, parent)
		} else {
			if queue[0].Val == v[0] {
				parent = queue[0]
			} else {
				parent = &TreeNode{Val: v[0]}
			}
		}

		if v[2] == 1 {
			parent.Left = &TreeNode{Val: v[1]}

		} else {
			parent.Right = &TreeNode{Val: v[1]}

		}
		if parent.Val != queue[0].Val {
			if parent.Left != nil {
				parent.Left = queue[0]
			} else if parent.Right != nil {
				parent.Right = queue[0]
			}
			queue = append(queue, parent)
			queue = queue[1:]
		}

	}
	return res
}
