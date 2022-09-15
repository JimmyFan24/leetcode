package level6

import (
	"fmt"
	"math"
)

/*
Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.
You must implement a solution with a linear runtime complexity and use only constant extra space.
Example 1:

Input: nums = [2,2,3,2,4,4,4]
Output: 3
*/
func SingleNumber(nums []int) int {

	res := 0

	/*
			1111 1011
			1111 1011
			1111 1011
			0000 0001

			3333 3034
		    0000 0001

	*/
	bit := make([]int, 32)
	for i := 0; i < 32; i++ {

		sum := 0
		for j := 0; j < len(nums); j++ {
			sum += (nums[j] >> i) & 1
		}
		bit[i] = sum
		if i < 31 {
			res |= (sum % 3) << i
		} else {
			t := (sum % 3) << i
			res -= t
		}

	}
	fmt.Println(bit)
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReorderList(head *ListNode) {

	if head.Next == nil {
		return
	}
	nodeList := []*ListNode{}
	cur := head
	for cur != nil {
		nodeList = append(nodeList, cur)
		cur = cur.Next
	}
	before := nodeList[:len(nodeList)/2]
	after := nodeList[len(nodeList)/2:]

	b, a := 0, len(after)-1
	for b < len(before) && a >= 0 {
		after[a].Next = nil
		before[b].Next = after[a]
		b++

		if b >= len(before) {
			a--

			break
		}
		before[b].Next = nil
		after[a].Next = before[b]
		a--

	}

	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	for a >= 0 {
		after[a].Next = nil
		curr.Next = after[a]

		a--
	}
	return
}

func InsertionSortList(head *ListNode) *ListNode {
	r := &ListNode{Val: math.MinInt32}
	res := r
	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	tmp := head
	//1--3--2--5---4
	i := 1
	t := tmp
	for ; i <= length; i++ {
		tmpRes := res
		prev := tmpRes
		for tmpRes != nil && tmpRes.Val < t.Val {
			prev = tmpRes
			tmpRes = tmpRes.Next
		}

		//1,3,5,  2-4

		prevtmp := prev.Next
		prev.Next = t
		tt := t.Next
		t.Next = prevtmp
		t = tt

	}
	return res.Next
}
