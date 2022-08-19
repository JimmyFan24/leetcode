package level4

import "strings"

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
