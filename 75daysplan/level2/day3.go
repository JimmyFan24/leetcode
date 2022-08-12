package level2

import (
	"leetcode/75daysplan/level1"
	"math/rand"
)

func GenerateParenthesis(n int) []string {

	if n==0{
		return []string{}
	}
	res := []string{}
	generate(n,n,"",&res)
	return res

}
func generate(left int,right int,str string,res *[]string){

	if left == 0 && right ==0{
		(*res) = append((*res),str)
		return
	}

	if left>0{
		generate(left-1,right,str+"(",res)
	}
	if right>0 && left<right{
		generate(left,right-1,str+")",res)
	}

}
type ListNode struct {
	    Val int
	     Next *ListNode
}

func RemoveNthFromEnd(head *level1.ListNode, n int) *level1.ListNode {

	queue := []*level1.ListNode{}
	for head !=nil{

		queue = append(queue,head)
		head = head.Next
	}

	queue[len(queue)-n-1].Next = queue[len(queue)-n].Next
	return queue[0]

}
func Subsets(nums []int) [][]int {

	tmp,res := []int{},[][]int{}

	for i:=0;i<len(nums);i++{
		sdfs(&res,nums,0,i,tmp,0)
	}


	return res
}
func sdfs(res *[][]int,nums[]int,index int,end int,tmp []int,start int){

	if end == len(tmp){
		t := make([]int,len(tmp))
		copy(t,tmp)
		(*res) = append((*res),t)
		return
	}
	for i:=start;i<len(nums);i++{

		tmp = append(tmp,nums[i])
		sdfs(res,nums,index+1,end,tmp,i+1)
		tmp = tmp[:len(tmp)-1]
	}
	return



}


type TreeNode struct {
    Val int
    Left *TreeNode
	Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {


	res := []int{}
	preorder(root.Left,&res)
	res = append(res,root.Val)
	preorder(root.Right,&res)

	return res[k]
}
func preorder(root *TreeNode,res *[]int){

	if root == nil{
		return
	}
	if root.Left != nil && root.Right !=nil{
		preorder(root.Left,res)
		(*res) = append((*res),root.Val)
		preorder(root.Right,res)
	}else if root.Left != nil && root.Right ==nil{
		preorder(root.Left,res)
		(*res) = append((*res),root.Val)
	}else if root.Right != nil && root.Left ==nil{
		preorder(root.Right,res)
		(*res) = append((*res),root.Val)
	}
	return
}

/*
	Input: nums = [3,2,1,5,6,4], k = 2
	Output: 5

	1,2,3,4,5,6
*/
func FindKthLargest(nums []int, k int) int {

	m := len(nums) - k+1
	res := selectSmallest(nums,0,len(nums)-1,m)
	return res
}
func selectSmallest(nums[]int,left int,right int,i int)int{

	if left >= right{
		return nums[left]
	}

	q := partition(nums,left,right)
	k := q - left +1
	if k == i{
		return nums[q]
	}
	if i<k{
		return selectSmallest(nums,left,q-1,i)
	}else{
		return selectSmallest(nums,q+1,right,i-k)
	}

}
func partition(nums []int,left,right int)int{

	k := left + rand.Intn(right -left +1)
	nums[k] ,nums[right] = nums[right],nums[k]

	i := left -1
	for j:= left;j<right;j++{
		if nums[j] <=nums[right]{
			i++
			nums[i],nums[j] = nums[j],nums[i]
		}
	}
	nums[i+1] , nums[right] = nums[right],nums[i+1]
	return i+1
}


