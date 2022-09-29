package level7

import (
	"strconv"
	"strings"
)

/*
An integer m is a divisor of n if there exists an integer k such that n = k * m.
*/

func IsThree(n int) bool {

	if n <= 2 {
		return false
	}
	k := 2
	count := 0
	for ; k*k <= n; k++ {
		if n%k == 0 {
			if n/k == k {
				count++
			} else {
				count += 2
			}

		}
	}
	if count == 1 {
		return true
	}
	return false
}
func MakeGood(s string) string {

	stack := []string{}

	for i := 0; i < len(s); i++ {
		if len(stack) < 1 {
			stack = append(stack, string(s[i]))
		} else {
			//取栈顶元素和现在遍历的元素比较
			topLetter := stack[len(stack)-1]
			curLetter := string(s[i])
			if isLowerUpperLetter(topLetter, curLetter) {
				//如果是需要整理的字符,从栈里面去掉,否则加上现在遍历的
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, curLetter)
			}

		}
	}
	res := ""
	for _, v := range stack {
		res += v
	}
	return res
}
func isLowerUpperLetter(a, b string) bool {

	if a >= "a" && a <= "z" && b >= "A" && b <= "Z" {
		//a小,b大
		if strings.ToUpper(a) == b {
			return true
		}
	} else if a >= "A" && a <= "Z" && b >= "a" && b <= "z" {
		if strings.ToUpper(b) == a {
			return true
		}
	}
	return false

}

func FindPoisonedDuration(timeSeries []int, duration int) int {

	durationSum := 0
	for i := 0; i < len(timeSeries); i++ {
		endtime := timeSeries[i] + duration

		if i == len(timeSeries)-1 || endtime <= timeSeries[i+1] {
			durationSum += duration
		} else {
			durationSum += timeSeries[i+1] - timeSeries[i]
		}
	}
	return durationSum
}

func DivisorSubstrings(num int, k int) int {

	count := 0
	str := ""
	tmp := num
	for tmp > 0 {

		t := tmp % 10
		str = strconv.Itoa(t) + str
		tmp /= 10
	}
	for i := 0; i < len(str)-k+1; i++ {

		s := str[i : i+k]
		n, _ := strconv.Atoi(s)
		if n > 0 && num%n == 0 {
			count++
		}
	}
	return count
}

func CountHillValley(nums []int) int {

	ans := 0

	for i := 1; i < len(nums)-1; i++ {

		//排除第一个和最后一个

		//1.判断是不是谷
		left, right := i-1, i+1
		//找到左边不相同的值
		for left >= 0 && nums[left] == nums[i] {
			left--
		}
		for right < len(nums) && nums[right] == nums[i] {
			right++
		}
		if left >= 0 && right < len(nums) {
			if nums[left] > nums[i] && nums[i] < nums[right] {
				ans++
			}
			if nums[left] < nums[i] && nums[i] > nums[right] {
				ans++
			}
			if right != i+1 {
				i = right - 1
			}
		}

	}
	return ans
}

/*
Example 1:

Input: n = 9
Output: false
Explanation: In base 2: 9 = 1001 (base 2), which is palindromic.
In base 3: 9 = 100 (base 3), which is not palindromic.
Therefore, 9 is not strictly palindromic so we return false.
Note that in bases 4, 5, 6, and 7, n = 9 is also not palindromic.
*/
func IsStrictlyPalindromic(n int) bool {

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/*
Example 1:

Input: nums = [3,2,4,6]
Output: 7
Explanation:
Apply the operation with x = 4 and i = 3, num[3] = 6 AND (6 XOR 4) = 6 AND 2 = 2.
Now, nums = [3, 2, 4, 2] and the bitwise XOR of all the elements = 3 XOR 2 XOR 4 XOR 2 = 7.
It can be shown that 7 is the maximum possible bitwise XOR.
Note that other operations may be used to achieve a bitwise XOR of 7.
*/
func MaximumXOR(nums []int) int {
	ans := 0

	for _, n := range nums {
		ans |= n
	}

	return ans
}

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func AllPossibleFBT(n int) []*TreeNode {
	res := []*TreeNode{}
	//A full binary tree is a binary tree where each node has exactly 0 or 2 children.
	//1.平衡的话
	height := 0
	for n > 0 {
		n /= 2
		height++
	}

	return res
}

func NumTilePossibilities(tiles string) int {

	res := 0
	m := []string{}
	visited := make([]bool, len(tiles))

	for i := 1; i <= len(tiles); i++ {
		dfsNumTile(&m, &visited, tiles, 0, i, 0, "")
	}
	mm := make(map[string]int, len(m))
	for _, v := range m {
		mm[v]++
	}
	for _, v := range mm {
		if v > 0 {
			res++
		}
	}
	return res
}
func dfsNumTile(m *[]string, visited *[]bool, tiles string, start, end int, count int, tmp string) {

	if len(tmp) == end {
		*m = append(*m, tmp)
		return
	}

	for i := 0; i < len(tiles); i++ {
		//"AAB"
		if !(*visited)[i] {
			(*visited)[i] = true
			tmp += string(tiles[i])
			dfsNumTile(m, visited, tiles, i+1, end, count+1, tmp)
			tmp = tmp[:len(tmp)-len(string(tiles[i]))]
			(*visited)[i] = false
		}

	}
}

type FindElements struct {
	Elements map[int]int
}

func Constructor1261(root *TreeNode) FindElements {

	arr := []int{}
	root.Val = 0
	recoverTreeNode(root, &arr)
	m := make(map[int]int, len(arr))
	for _, v := range arr {
		m[v]++
	}
	return FindElements{Elements: m}
}

func recoverTreeNode(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	if root.Left != nil {
		root.Left.Val = 2*root.Val + 1

	}
	if root.Right != nil {
		root.Right.Val = 2*root.Val + 1
	}
	recoverTreeNode(root.Left, arr)
	recoverTreeNode(root.Right, arr)
}
func (this *FindElements) Find(target int) bool {

	return this.Elements[target] > 0
}
func ReverseOddLevels(root *TreeNode) *TreeNode {

	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {

		l := len(queue)
		tmp := make([]*TreeNode, l)
		copy(tmp, queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		if level%2 != 0 {
			//奇数
			ll := len(tmp)
			for i := 0; i < ll/2; i++ {
				tmp[i].Val, tmp[ll-i+1].Val = tmp[ll-i+1].Val, tmp[i].Val
			}
		}
		queue = queue[l:]
		level++
	}
	return root

}
