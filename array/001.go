package main

import "fmt"


//<--1.sum-->
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i,item := range nums{
		anotherNum := target - item
		if _,ok := m[anotherNum];ok{
			return []int{m[anotherNum],i}
		}
		m[nums[i]] = i
	}
	return nil
}
type ListNode struct {
	Val int
	Next *ListNode
}
/*
	3.最长子串
 */
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0{
		return 0
	}
	result,left,right := 0,0,-1
	var freq [127]int
	for left < len(s){
		//
		if right + 1 < len(s) && freq[s[right + 1]] == 0 {
			freq[s[right + 1]]++
			right++
		}else {
			b := freq[s[left]]
			freq[s[left]]--
			left++
			fmt.Println(b)
		}

		result = max(result,right-left+1)
	}
	fmt.Println(result)
	return result

}
func max( a int, b int)int{
	if a>b{
		return a
	}
	return b
}
/*
	2. 数字相加
 */
func addTwoNumbers(l1 *ListNode,l2 *ListNode) *ListNode{
	head := &ListNode{
		Val: 0,
	}
	n1,n2,carry,current := 0,0,0,head

	for l1 != nil || l2 != nil || carry !=0{
		if l1 == nil{
			n1 = 0
		}else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil{
			n2 = 0
		}else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
//<--2.-->
