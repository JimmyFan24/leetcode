package level6

import (
	"math"
	"strconv"
	"strings"
)

func RemoveDuplicates(nums []int) int {

	lenght := len(nums)
	//1,2,2,2,2,3,3,4,4,5
	for i := 0; i < lenght; {

		if i+1 < lenght && nums[i] == nums[i+1] {
			oriI := i
			for i+1 < lenght && nums[i] == nums[i+1] {
				i++
			}
			j := i
			remove := j - oriI - 1
			if remove > 0 {
				q := oriI + 1
				for j < lenght {

					nums[q] = nums[j]
					j++
					q++
				}
				lenght -= remove
			}
			i = oriI + 2

		} else {
			i++
		}

	}
	return lenght
}

func DeleteDuplicates(head *ListNode) *ListNode {

	newHead := &ListNode{}
	res := newHead
	cur := head
	for cur != nil {
		tmp := cur
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		//相等说明没有重复
		if cur == tmp {
			t := cur.Next
			cur.Next = nil
			newHead.Next = cur
			newHead = newHead.Next
			cur = t
		} else {
			cur = cur.Next
		}

	}
	if newHead == res {
		return nil
	}
	return res.Next

}

type BSTIterator struct {
	HasNextEl bool
	Arr       []int
}

func Constructor(root *TreeNode) BSTIterator {

	hasnext := true
	if root == nil {
		hasnext = false
	}
	return BSTIterator{
		hasnext,
		InOrderPrev(root),
	}
}
func InOrderPrev(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return nil
	}
	left := InOrderPrev(root.Left)
	for _, v := range left {
		res = append(res, v)
	}
	res = append(res, root.Val)
	right := InOrderPrev(root.Right)
	for _, v := range right {
		res = append(res, v)
	}
	return res
}
func (this *BSTIterator) Next() int {

	next := this.Arr[0]
	this.Arr = this.Arr[1:]
	if len(this.Arr) < 1 {
		this.HasNextEl = false
	}
	return next
}

func (this *BSTIterator) HasNext() bool {

	return this.HasNextEl
}

/*
Example 2:

Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
*/
func FindMin(nums []int) int {
	l, r := 0, len(nums)-1

	min := math.MaxInt32
	for l < r {
		mid := l + ((r - l) >> 1)
		//说明l-mid还是递增区间，最小值应该在[mid+1 - r],否则还是在l-mid区间
		//1.左边最小值如果大于右边最大值，那必然在mid-r
		if nums[mid] < min {
			min = nums[mid]
		}
		if nums[l] > nums[r] {
			if nums[mid] > nums[l] {
				l = mid
			} else {
				r = mid
			}

		} else if nums[l] < nums[r] {
			return nums[l]
		}

	}
	return nums[l+1]
}
func CountNodes(root *TreeNode) int {

	res := 0
	if root == nil {
		return 0
	}

	left := CountNodes(root.Left)
	right := CountNodes(root.Right)
	return res + left + right
}

/*
Example 1:

Input: s = "3+2*2"
Output: 7
*/
func Calculate(s string) int {

	numStack := []string{}
	//1.数字和运算符入栈，遇到乘法或者除法运算符要 先处理完再把处理结果压栈
	for i := 0; i < len(s); {
		//遇到数字，压栈
		if s[i] >= '0' && s[i] <= '9' {
			number := ""
			j := i
			//获取下一个数字
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				number += string(s[j])
				j++
			}
			numStack = append(numStack, number)
			i = j
		} else if s[i] == '*' || s[i] == '/' {
			//乘法和除法要处理
			//
			before := numStack[len(numStack)-1]
			beforeNumber, _ := strconv.Atoi(before)
			//获取下一个数字
			next := ""
			j := i + 1
			//获取下一个数字

			for j < len(s) && !(s[j] >= '0' && s[j] <= '9') {
				j++
			}
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				next += string(s[j])
				j++
			}

			nextnumber, _ := strconv.Atoi(next)
			curnumber := 0
			if s[i] == '/' {
				curnumber = beforeNumber / nextnumber
			} else {
				curnumber = nextnumber * beforeNumber
			}
			//把处理结果重新压栈
			numStack = numStack[:len(numStack)-1]
			numStack = append(numStack, strconv.Itoa(curnumber))

			i = j
		} else if s[i] == '-' || s[i] == '+' {
			numStack = append(numStack, string(s[i]))
			i++
		} else {
			i++
		}

	}
	newvalue := 0

	//剩下的栈里面只有+和-，从左到右执行就行

	for i := 0; i < len(numStack); {

		//如果当前是数字
		if strings.Contains(numStack[i], "-") {
			i++
			after := numStack[i]
			a, _ := strconv.Atoi(after)
			newvalue -= a
			i++
		} else if strings.Contains(numStack[i], "+") {
			i++
			after := numStack[i]
			a, _ := strconv.Atoi(after)
			newvalue += a
			i++
		} else {
			before := numStack[i]
			b, _ := strconv.Atoi(before)
			newvalue += b
			i++
		}
		//如果当前不是数字

	}
	return newvalue
}
func calcul(v1, v2, sign string) int {
	res := 0
	vv1, _ := strconv.Atoi(v1)
	vv2, _ := strconv.Atoi(v2)
	switch sign {
	case "+":
		res = vv1 + vv2
	case "-":
		res = vv2 - vv1
	case "*":
		res = vv2 * vv1
	case "/":
		res = vv2 / vv1

	default:
		return res
	}
	return res
}

func CombinationSum3(k int, n int) [][]int {

	res := [][]int{}

	tmp := []int{}
	dfsCombibnationSum3(1, k, n, &res, 0, tmp)
	return res

}
func dfsCombibnationSum3(start, k, n int, res *[][]int, count int, tmp []int) {

	if count == k {
		if n == 0 {
			t := make([]int, k)
			copy(t, tmp)
			*res = append(*res, t)

		}
		return
	}

	for i := start; i <= 9; i++ {

		if n-i >= 0 {
			tmp = append(tmp, i)
			dfsCombibnationSum3(i+1, k, n-i, res, count+1, tmp)
			tmp = tmp[:len(tmp)-1]
		} else {
			break
		}
	}
	return
}
