package level2

import (
	"strconv"
)

/*
Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 */
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	res := []int{}
	p := 0
	q := 0
	for p<len(nums1) && q<len(nums2){

		if nums1[p] <=nums2[q]{

			res = append(res,nums1[p])
			p++
		}else{
			res = append(res,nums2[q])
			q++
		}
	}

	for p<len(nums1){
		res = append(res,nums1[p])
		p++
	}
	for q<len(nums2){
		res = append(res,nums2[q])
		q++
	}

	mid := len(res)/2
	if len(res)%2==0{
		return (float64(res[mid])+float64(res[mid-1]))/2
	}
	return float64(res[mid])
}

/*
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次
 */


func IsValidSudoku(board [][]byte) bool {



	for i:=0;i<9;i++{
		tmp :=[10]int{}
		for j:=0;j<9;j++{

			cellVal := board[i][j]
			if string(cellVal) != "."{
				index, _ := strconv.Atoi(string(cellVal))
				if index >9 || index<1{
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}

	for i:=0;i<9;i++{
		tmp :=[10]int{}
		for j:=0;j<9;j++{

			cellVal := board[j][i]
			if string(cellVal) != "."{
				index, _ := strconv.Atoi(string(cellVal))
				if index >9 || index<1{
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}

	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			tmp := [10]int{}
			for ii:=i*3;ii<i*3+3;ii++{
				for jj:=j*3;jj<j*3+3;jj++{

					cellVal := board[ii][jj]
					if string(cellVal) != "." {
						index, _ := strconv.Atoi(string(cellVal))
						if tmp[index] == 1 {
							return false
						}
						tmp[index] = 1
					}
				}
			}
		}
	}
	return true
}

 /*
     * clockwise rotate 顺时针旋转
     * first reverse up to down, then swap the symmetry
     * 1 2 3     7 8 9     7 4 1
     * 4 5 6  => 4 5 6  => 8 5 2
     * 7 8 9     1 2 3     9 6 3
*/

func Rotate(matrix [][]int)  {

	length := len(matrix)

	/*
		1 2 3    1 4 7
		4 5 6    2 5 8
		7 8 9    3 6 9
	 */
	for i:=0;i<length;i++{
		for j:=i+1;j<length;j++{
			matrix[i][j],matrix[j][i] = matrix[j][i],matrix[i][j]
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-j-1] = matrix[i][length-j-1], matrix[i][j]
		}
	}
}
/*
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).
 */
func IsMatch(s string, p string) bool {

	sByte := []byte(s)
	pByte := []byte(p)

	/*
	"aab"
	"c*a*b"
	 */




	return sDFS(sByte,pByte,0)
}
func sDFS(sByte []byte,pByte []byte,index int) bool{

	if index == len(pByte){

		return false
	}

	i := 0
	j := 0
	/*
		"pi"
		"*."
	*/
	for ;i<len(pByte)&& j<len(sByte);{
		if sByte[j] == pByte[i]{
			i++
			j++
		}else {

			if string(pByte[i]) == "."{
				i++
				j++
			}else if string(pByte[i]) == "*"{
				i++
				if i==len(pByte){
					return true
				}else{
					for j <len(sByte) && (sByte[j] != pByte[i] || string(pByte[i]) == ".")  {
						j++
					}
					if j+1 == len(sByte){
						return true
					}

				}
			}else{
				return sDFS(sByte,pByte[index+1:],index+1)
			}
		}


	}
	if j == len(sByte){
		return true
	}
	return false
}

/*
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
 */
func MergeKLists(lists []*ListNode) *ListNode {

	l := len(lists)
	if l ==1{
		return lists[0]
	}
	if l <1{
		return nil
	}
	num := l/2
	left := MergeKLists(lists[:num])
	right := MergeKLists(lists[num:])
	return mergeTwoLists(left,right)

}
func mergeTwoLists(list1 *ListNode ,list2 *ListNode)*ListNode{

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val{
		list1.Next = mergeTwoLists(list1.Next,list2)
		return list1
	}
	list2.Next = mergeTwoLists(list2.Next,list2)
	return list2

}
/*
Example 1:

Input: nums = [1,2,0]
Output: 3
Example 2:

Input: nums = [3,4,-1,1]
Output: 2
 */
func FirstMissingPositive(nums []int) int {

	min := nums[0]
	m := 0
	for i,v := range  nums{
		if v < min{
			min = v
			m = i
		}
	}
	return m
}
/*
Example 1:

Input: nums = [1,2,3], target = 4
Output: 7
Explanation:
The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
Note that different sequences are counted as different combinations.
 */
func CombinationSum4(nums []int, target int) int {

	dp := make([]int,target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
	/*//1,2,3
	dp[4] = dp[4-1] + dp[4-2]  = dp[3] + dp[2]
	dp[3] = d[3-1] + dp[3-2] = dp[2] + dp[1]
	dp[2] = dp[2-1] + dp[2-2] = dp[1] + dp[0]
	dp[1] = dp[0]*/
	//定义 dp[i] 为总和为 target = i 的组合总数
	// 总和等于1 的数量,也就是总和 (3-第一个原素i的数量 ) + 总和(3 -第二个元素i) + 总和 (3-最后一个元素)


}
func comDFS(nums []int,res *int,index int,sum int,target int)  {

	if sum > target{
		return
	}

	for i:= index;i<len(nums);i++{


		if sum == target{
			(*res)++
			return
		}else if sum < target {
			(sum) = (sum) + nums[i]
			comDFS(nums,res,index,sum,target)
			sum = sum -nums[i]
		}else  {
			return
		}
	}
	return
}