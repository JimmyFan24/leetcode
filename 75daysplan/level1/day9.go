package level1

import (
	"strconv"
)

func LongestCommonPrefix(strs []string) string {

	m := len(strs[0])
	k := 0
	for i,s := range strs{
		if len(s)<m{
			m = len(s)
			k = i
		}
	}

	prefix := strs[k]

	for i:=0;i<len(strs);i++{
		if i == k{
			continue
		}
		for j:=0;j<len(prefix);j++{

			if strs[i][j] != prefix[j]{
				prefix = prefix[0:j]
				break
			}


		}
	}

	return prefix

}
func IsValid(s string) bool {

	match := []string{}
	mMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	for i:=0;i<len(s);i++{

		//左括号入栈
		if string(s[i]) == "(" || string(s[i]) == "{" ||string(s[i]) == "["{
			match = append(match,string(s[i]))
		}else{
			//遇到右括号,开始出栈对比
			if string(s[i]) == mMap[match[len(match)-1]]{
				match = match[:len(match)-1]
			}
		}


	}
	if len(match) == 0{
		return true
	}
	return false
}

func PlusOne(digits []int) []int {

	s := ""

	for _,v := range digits{

		s += strconv.Itoa(v)
	}

	n,_ := strconv.ParseInt(s,10,64)
	n = n+1
	ns := strconv.FormatInt(n,10)

	res := []int{}
	for _,b := range ns{
		nb ,_ := strconv.Atoi(string(b))
		res = append(res,nb)
	}
	return res

}




func IsPalindrome(head *ListNode) bool {

	h := head
	revList := ReverseListNode(head)
	for h.Next != nil && revList.Next != nil{
		if h.Val != revList.Val{
			return false
		}
		h = h.Next
		revList = revList.Next
	}
	return true

}

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}
func ReverseListNode(head *ListNode) *ListNode{


	var tail *ListNode = nil


	cur := head.Next
	pre := head
	//nil<--next--1 2(cur)-->2-->1
	head.Next = tail
	//1-->2-->2-->1
	//1--next-->2--next-->2--next-->1
	for cur != nil{
		//tmp==2
		tmp :=  cur.Next
		//改变2 的指向 nil<--1(pre)<--2(cur) -->2-->1
		cur.Next = pre
		pre = cur
		cur = tmp

	}
	return pre

}

func Subsets(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	for k := 0; k <= len(nums); k++ {
		GenerateSubsets(nums, k, 0, c, &res)
	}
	return res
}

func GenerateSubsets(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// i will at most be n - (k - c.size()) + 1
	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		c = append(c, nums[i])
		GenerateSubsets(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}

	return
}