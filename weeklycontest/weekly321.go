package weeklycontest

import "sort"

func PivotInteger(n int) int {

	// 1,2,3,4,5

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	frontSum := 0

	// 1,2,3,4,5,6
	for i := 1; i < n; i++ {

		if frontSum == sum {
			return i
		} else if frontSum > sum {
			break
		}
		frontSum += i
		sum = sum - i + 1
	}
	return -1
}

func AppendCharacters(s string, t string) int {

	i := 0
	count := len(t)
	for i < len(t) {

		j := 0
		for j < len(s) && s[j] != t[i] {
			j++
		}
		//如果在s中找到
		if j < len(s) {
			count--
			if j+1 < len(s) {
				s = s[j+1:]
			} else {
				break
			}

		}
		//如果在s中找不到，t当前i后面的都没办法了，因为顺序不能改，直接返回
		if j >= len(s) {
			return count
		}
		i++

	}
	return count
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNodes(head *ListNode) *ListNode {

	nodes := []*ListNode{}
	for head != nil {
		if len(nodes) != 0 {
			//如果栈里面的值比当前的值小，都要pop出来
			for len(nodes) > 0 && nodes[len(nodes)-1].Val < head.Val {
				nodes = nodes[:len(nodes)-1]
			}

		}
		nodes = append(nodes, head)
		head = head.Next
	}
	newHead := &ListNode{}
	res := newHead
	for _, node := range nodes {
		newHead.Next = &ListNode{Val: node.Val}
		newHead = newHead.Next
	}

	return res
}

func CountSubarrays(nums []int, k int) int {

	count := 0
	for i := 0; i < len(nums); i++ {
		j := i + 1
		for ; j <= len(nums); j++ {
			if matchMedian(k, nums[i:j]) {
				count++
			}
		}
	}
	return count
}
func matchMedian(k int, arr []int) bool {

	c := make([]int, len(arr))
	copy(c, arr)
	sort.Ints(c)
	if len(c)%2 == 0 {
		return c[len(c)/2-1] == k
	}
	return c[len(c)/2] == k
}
