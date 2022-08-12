package level3

import (
	"leetcode/75daysplan/level2"
	"strconv"
	"strings"
)

/*
Example 1:

Input: pattern = "abba", s = "dog cat cat dog"
Output: true
Example 2:

Input: pattern = "abba", s = "dog cat cat fish"
Output: false
 */
func WordPattern(pattern string, s string) bool {


	sList := strings.Split(s," ")
	if len(pattern) != len(sList){
		return false
	}
	sMap := make(map[rune]string,len(pattern))
	lMap := make(map[string]rune,len(sList))

	for i,v := range pattern{
		if _,ok := lMap[sList[i]];!ok{
			lMap[sList[i]] = v
		}else{
			if lMap[sList[i]] !=v{
				return false
			}
		}
		if _,ok := sMap[v];!ok{
			sMap[v] = sList[i]
		}else{
			if sMap[v] != sList[i]{
				return false
			}
		}

	}

	return true
}
func IsPowerOfFour(n int) bool {

	if n ==1{
		return true
	}
	for n>0{

		if n%4==0{
			n = n/4
			if n == 1{
				return true
			}
		}else{
			return false
		}

	}
	return false
}
func CanConstruct(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine){
		return false
	}
	var cnt [26]int
	for _,v:= range ransomNote{

		cnt[v-'a']++
	}
	for _,v:= range magazine{

		cnt[v-'a']--
		if cnt[v-'a']<0{
			return false
		}
	}
	return true

}
func ReverseKGroup(head *level2.ListNode, k int) *level2.ListNode {

	node := head
	for i:=0;i<k;i++{
		if node == nil{
			return head
		}
		node = node.Next
	}
	//1--2--3--4--5--6--7--8--9
	//3--2--1--4--5--6--7--8--9
	newHead := reverse(head,node)
	//3--2--1--4
	head.Next = ReverseKGroup(node,k)
	return newHead
}
func reverse(first, last *level2.ListNode)*level2.ListNode{

	pre := last
	//2--1--3--4
	for first != last{

		tmp := first.Next
		first.Next = pre
		pre = first
		first = tmp
	}
	return pre

}
/*
Example 1:

Input: nums = [10,2]
Output: "210"
Example 2:

Input: nums = [3,30,9,5,34]
Output: "9534330"
 */
/*

	[9,5,34,3,30]
	[9,5,3,3,3]
 */
func LargestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	//1.先转字符串数组
	strList := toStringArrays(nums)
	//2.字符串数组进行排序
	quickSortString(strList,0,len(strList)-1)
	res := ""
	for _, str := range strList {
		if res == "0" && str == "0" {
			continue
		}
		res = res + str
	}
	return res
	//3.拼接排序后的数组


}
func toStringArrays(nums []int)[]string{

	res := []string{}

	for _,v:= range nums{

		res = append(res,strconv.Itoa(v))
	}
	return res
}
//[9,5,34,3,30]
func partitionString(a[]string,l,h int)int {
	pivot := a[h]

	i :=l-1
	for j:=l;j<h;j++{

		aStr := a[j] + pivot
		bStr := pivot + a[j]
		if aStr > bStr{
			i++
			a[i],a[j] = a[j],a[i]
		}
	}
	a[i+1],a[h] = a[h],a[i+1]
	return i+1
}
func quickSortString(a []string,l,h int){

	if l >h{
		return
	}

	p := partitionString(a,l,h)
	quickSortString(a,l,p-1)
	quickSortString(a,p+1,h)
}
func MinPathSum(grid [][]int) int {



	m :=len(grid)
	if m ==0{
		return 0
	}
	res := 0
	for i:=0;i<m;i++{
		res += grid[i][0]
	}
	for j:=1;j<len(grid[0])-1;j++{
		res += grid[m-1][j]
	}

	dfs(grid,0,0,&res,0)
	return res
}
func dfs(grid [][]int,x,y int ,res *int,sum int){

	if x==len(grid)-1 && y==len(grid[0])-1{
		*res = min(*res,sum+grid[x][y])
	}
	if sum > *res{
		return
	}
	if inBoard(grid ,x,y){
		if inBoard(grid,x,y+1){
			dfs(grid,x,y+1,res,sum+grid[x][y])
		}else if inBoard(grid,x+1,y){
			dfs(grid,x+1,y,res,sum+grid[x][y])
		}

	}
	return

}
func inBoard(grid [][]int ,x,y int)bool{

	return x <len(grid)&& x>=0 && y<len(grid[0])&& y>=0
}
func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}
func PathSum(root *level2.TreeNode, targetSum int) int {


	res :=0
	sum(root,targetSum,&res)
	return res
}
func sum(root *level2.TreeNode, targetSum int ,res *int) {
	if root ==nil{
		return
	}

	if root.Val >targetSum{
		return
	}
	if root.Val == targetSum {
		*res++
	}
	sum(root.Left,targetSum-root.Val,res)
	sum(root.Right,targetSum-root.Val,res)
	return
}
/*
Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
 */
func CanJump(nums []int) bool {
	if len(nums)==1 {
		return true
	}
	//   1 1 1 1 1}
	// 0,1,2,3,4,5
	//[5,3,3,2,1,0,2,3,3,1,0,0]
	//[3,2,1,0,4]


	for i:=0;i<len(nums);i++{
		j:=i
		max := nums[j]
		index := j
		for ;j<len(nums);j++{
			if nums[j] >= max{
				max = nums[j]
				index = j
			}
			if nums[j] == 0 && nums[j+1]!=0{
				break
			}
		}

		if max >= len(nums)-index-1 || j==len(nums){
			return true
		}
		if max <= j-index {
			return false
		}
		i =j
	}

	return true
}
