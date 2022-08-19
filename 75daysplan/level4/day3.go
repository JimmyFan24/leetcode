package level4

import (
	"math"
	"strings"
)

func deepestLeavesSum(root *TreeNode) int {

	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	sum := 0
	for len(queue) > 0 {
		l := len(queue)
		tmp := []int{}
		for i := 0; i < l; i++ {

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}

		queue = queue[l:]
		if len(queue) == 0 {
			for _, v := range tmp {
				sum += v
			}
		}
	}
	return sum
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Input: head = [0,3,1,0,4,5,2,0]
Output: [4,11]
Explanation:
The above figure represents the given linked list. The modified list contains
- The sum of the nodes marked in green: 3 + 1 = 4.
- The sum of the nodes marked in red: 4 + 5 + 2 = 11.
*/
func MergeNodes(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	new := &ListNode{}
	res := new
	for head.Next != nil {

		tmp := head.Next
		sum := 0
		for tmp.Val != 0 {
			sum += tmp.Val
			tmp = tmp.Next
		}
		newnode := &ListNode{
			Val:  sum,
			Next: nil,
		}
		new.Next = newnode
		new = new.Next
		head = tmp
	}
	return res.Next
}
func averageOfSubtree(root *TreeNode) int {

	if root == nil {
		return 0
	}
	count := 0
	if matchNode(root) {
		count++
		count += averageOfSubtree(root.Left)
		count += averageOfSubtree(root.Right)
	}
	return count
}
func getCount(root *TreeNode) int {

	if root == nil {
		return 0
	}
	return getCount(root.Left) + getCount(root.Right) + 1

}
func getSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getSum(root.Left) + getSum(root.Right) + root.Val
}
func matchNode(root *TreeNode) bool {

	return root.Val == getSum(root)/getCount(root)
}

/*
Example 1:

Input: url = "https://leetcode.com/problems/design-tinyurl"
Output: "https://leetcode.com/problems/design-tinyurl"

Explanation:
Solution obj = new Solution();
string tiny = obj.encode(url); // returns the encoded tiny url.
string ans = obj.decode(tiny); // returns the original url after deconding it.
*/
type Codec struct {
	//字符集移后5位

}

func URLConstructor() Codec {
	return Codec{}
}

// Encodes a URL to a shortened URL.
func (this *Codec) Encode(longUrl string) string {

	res := ""
	urlList := strings.Split(longUrl, "/")
	res += urlList[0] + "//"
	for i := 2; i < len(urlList); i++ {
		tmp := ""
		for j := 0; j < len(urlList[i]); j++ {
			tmp += string(urlList[i][j] + 5)
		}
		if i == len(urlList)-1 {
			res += tmp
		} else {
			res += tmp + "/"
		}
	}
	return res
}

// Decodes a shortened URL to its original URL.
func (this *Codec) Decode(shortUrl string) string {
	res := ""
	urlList := strings.Split(shortUrl, "/")
	res += urlList[0] + "//"
	for i := 2; i < len(urlList); i++ {
		tmp := ""
		for j := 0; j < len(urlList[i]); j++ {
			tmp += string(urlList[i][j] - 5)
		}
		if i == len(urlList)-1 {
			res += tmp
		} else {
			res += tmp + "/"
		}

	}
	return res

}
func sumEvenGrandparent(root *TreeNode) int {

	if root == nil {
		return 0
	}
	sum := 0
	if root.Val%2 == 0 {
		if root.Left != nil {
			if root.Left.Left != nil {
				sum += root.Left.Left.Val
			}
			if root.Left.Right != nil {
				sum += root.Left.Right.Val
			}
		}
		if root.Right != nil {
			if root.Right.Left != nil {
				sum += root.Right.Left.Val
			}
			if root.Right.Right != nil {
				sum += root.Right.Right.Val
			}
		}

	}
	return sum + sumEvenGrandparent(root.Left) + sumEvenGrandparent(root.Right)
}
func getGrandSon(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return getSon(root.Left) + getSon(root.Right) + root.Val
}

func getSon(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return getSon(root.Right) + getSon(root.Left) + root.Val

}
func BstToGst(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}
	sum := 0
	ReverMidOrder(root, &sum)
	return root

}
func ReverMidOrder(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	if root.Right != nil {
		ReverMidOrder(root.Right, sum)
	}

	root.Val += *sum
	*sum = root.Val
	if root.Left != nil {
		ReverMidOrder(root.Left, sum)
	}
	return
}

/*
二叉树的根是数组中的最大元素。
左子树是通过数组中最大值左边部分构造出的最大二叉树。
右子树是通过数组中最大值右边部分构造出的最大二叉树。
*/
func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := nums[0]
	index := 0
	for i, v := range nums {
		if v > max {
			index = i
			max = v
		}
	}

	left := ConstructMaximumBinaryTree(nums[:index])
	right := ConstructMaximumBinaryTree(nums[index+1:])
	root := &TreeNode{
		Val: max,
	}
	root.Left = left
	root.Right = right
	return root

}
func BstFromPreorder(preorder []int) *TreeNode {

	lower := math.MinInt32
	upper := math.MaxInt32
	idx, n := 0, len(preorder)
	return constructTree(lower, upper, &idx, &n, preorder)
}
func constructTree(lower, upper int, idx *int, n *int, preorder []int) *TreeNode {

	if *idx == *n {
		return nil
	}
	val := preorder[*idx]
	if val < lower || val > upper {
		return nil
	}
	*idx++
	new := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	new.Left = constructTree(lower, val, idx, n, preorder)
	new.Right = constructTree(val, upper, idx, n, preorder)
	return new
}
