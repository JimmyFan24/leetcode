package level8

import (
	"math"
	"strconv"
)

func CountDistinctIntegers(nums []int) int {

	m := make(map[int]int, len(nums)*2)
	for _, v := range nums {
		nums[v]++
	}

	for _, v := range nums {
		newNum := reverse(v)
		m[newNum]++
	}
	count := 0
	for _, v := range m {
		if v > 0 {
			count++
		}
	}
	return count
}
func reverse(n int) int {

	s := ""
	for n > 0 {
		tmp := n % 10
		s += strconv.Itoa(tmp)
		n /= 10
	}
	newNum, _ := strconv.Atoi(s)
	return newNum

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the root of a binary tree,
find the maximum value v for which there exist different nodes
a and b where v = |a.val - b.val| and a is an ancestor of b.
A node a is an ancestor of b if either:
any child of a is equal to b or any child of a is an ancestor of b.
*/
func MaxAncestorDiff(root *TreeNode) int {

	//1.v = |a.val - b.val| ,求v的最大值
	//2.a 是 b 的祖先

	res := 0
	dfsAncestorDiff(root, &res)
	return res

}
func dfsAncestorDiff(root *TreeNode, res *int) (int, int) {
	if root == nil {
		return -1, -1
	}
	leftMax, leftMin := dfsAncestorDiff(root.Left, res)
	if leftMax == -1 && leftMin == -1 {
		leftMax = root.Val
		leftMin = root.Val
	}
	rightMax, rightMin := dfsAncestorDiff(root.Right, res)
	if rightMax == -1 && rightMin == -1 {
		rightMax = root.Val
		rightMin = root.Val
	}
	//因为是绝对值,所以要算节点和节点后代里面最小值和最大值的差值的绝对值
	//root.Val-min(leftMin, rightMin) : 这个是节点本身与 左右最小值的差值的绝对值
	//root.Val-max(leftMax, rightMax) : 这个是节点本身和左右最大值的差值的绝对值
	*res = max(*res, max(int(math.Abs(float64(root.Val-min(leftMin, rightMin)))),
		int(math.Abs(float64(root.Val-max(leftMax, rightMax))))))
	return max(leftMax, max(rightMax, root.Val)), min(leftMin, min(rightMin, root.Val))

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FindBottomLeftValue(root *TreeNode) int {

	//层序遍历,最后一层的第一个就是
	queue := []*TreeNode{root}
	res := 0
	for len(queue) > 0 {
		l := len(queue)
		res = queue[0].Val
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return res

}
func leveltraversal(root *TreeNode) int {

	queue := []*TreeNode{root}
	res := 0
	for len(queue) > 0 {
		l := len(queue)
		res = queue[0].Val
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return res
}

/*
Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days
you have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.
Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
*/
func DailyTemperatures(temperatures []int) []int {

	//常规方法会超时,递减栈
	//递减栈里只会有递减的元素
	answer := make([]int, len(temperatures))
	statck := []int{}
	for i := 0; i < len(temperatures); i++ {

		//如果栈不空，且当前数字大于栈顶元素,需要取出栈顶元素,这个时候两者下标就是所求的
		for len(statck) > 0 && temperatures[statck[len(statck)-1]] < temperatures[i] {
			top := statck[len(statck)-1]
			answer[top] = i - top
			statck = statck[:len(statck)-1]
		}
		statck = append(statck, i)

	}
	return answer

}

/*
You are given a string s that contains some bracket pairs,
with each pair containing a non-empty key.
For example, in the string "(name)is(age)yearsold",
there are two bracket pairs that contain the keys "name" and "age".
You know the values of a wide range of keys.
This is represented by a 2D string array knowledge where each knowledge[i] = [keyi, valuei]
indicates that key keyi has a value of valuei.

You are tasked to evaluate all of the bracket pairs.
When you evaluate a bracket pair that contains some key keyi, you will:

Replace keyi and the bracket pair with the key's corresponding valuei.
If you do not know the value of the key, you will replace keyi and the bracket pair with a question mark "?" (without the quotation marks).
Each key will appear at most once in your knowledge. There will not be any nested brackets in s.

Return the resulting string after evaluating all of the bracket pairs.

Input: s = "(name)is(age)yearsold", knowledge = [["name","bob"],["age","two"]]
Output: "bobistwoyearsold"
Explanation:
The key "name" has a value of "bob", so replace "(name)" with "bob".
The key "age" has a value of "two", so replace "(age)" with "two".
*/
func Evaluate(s string, knowledge [][]string) string {

	keys := make(map[string]string, len(knowledge)*2)
	res := ""
	for _, v := range knowledge {
		key := v[0]
		value := v[1]
		keys[key] = value
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			j := i + 1
			tmp := ""
			for j < len(s) && s[j] != ')' {
				tmp += string(s[j])
				j++
			}
			i = j
			if keys[tmp] != "" {
				res += keys[tmp]
			} else {
				res += "?"
			}
		} else {
			res += string(s[i])
		}
	}

	return res
}

/*
You are given an array of strings products and a string searchWord.
Design a system that suggests at most three product
names from products after each character of searchWord is typed.
Suggested products should have common prefix with searchWord.
If there are more than three products with a common prefix return the three lexicographically minimums products.

Return a list of lists of the suggested products after each character of searchWord is typed.
Input: products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
Output: [["mobile","moneypot","monitor"],["mobile","moneypot","monitor"],["mouse","mousepad"],
["mouse","mousepad"],["mouse","mousepad"]]
Explanation: products sorted lexicographically = ["mobile","moneypot","monitor","mouse","mousepad"].
After typing m and mo all products match and we show user ["mobile","moneypot","monitor"].
After typing mou, mous and mouse the system suggests ["mouse","mousepad"]

*/
