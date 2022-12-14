package level3

import (
	"math"
	"strconv"
)

func ConvertToBase7(num int) string {

	if num < 7 && num >= 0 {
		return strconv.Itoa(num)
	} else if num < 0 {
		return "-" + ConvertToBase7(-num)
	}
	return ConvertToBase7(num/7) + strconv.Itoa(num%7)
}

/*
Example 1:

Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
Output: [-1,3,-1]
Explanation: The next greater element for each value of nums1 is as follows:
- 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
- 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.
- 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
*/
func NextGreaterElement(nums1 []int, nums2 []int) []int {

	mMap := make(map[int]int, len(nums2))
	for i, v := range nums2 {
		mMap[v] = i
	}
	res := []int{}

	for i := 0; i < len(nums1); i++ {
		p := mMap[nums1[i]]
		if p == len(nums2)-1 {
			res = append(res, -1)
			continue
		}
		j := p + 1
		for ; j < len(nums2); j++ {
			if nums2[j] > nums2[p] {
				res = append(res, nums2[j])
				break
			}
		}
		if j == len(nums2) {
			res = append(res, -1)
		}

	}
	return res
}

/*
Example 1:

Input: a = "aba", b = "cdc"
Output: 3
Explanation: One longest uncommon subsequence is "aba" because "aba" is a subsequence of "aba" but not "cdc".
Note that "cdc" is also a longest uncommon subsequence.
*/
func FindLUSlength(a string, b string) int {
	//abcde  bcde
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)

}
func GetMinimumDifference(root *TreeNode) int {

	res := 0

	if root == nil {
		return 0

	}
	left := GetMinimumDifference(root.Left)
	right := GetMinimumDifference(root.Right)

	res = min(left, right)
	if root.Left == nil && root.Right == nil {
		res = min(root.Val, res)
	} else if root.Left != nil && root.Right == nil {
		res = min(res, int(math.Abs(float64(root.Val-root.Left.Val))))
	} else if root.Right != nil && root.Left == nil {
		res = min(res, int(math.Abs(float64(root.Val-root.Right.Val))))
	}
	res = min(res, min(int(math.Abs(float64(root.Left.Val-root.Right.Val))), min(int(math.Abs(float64(root.Val-root.Right.Val))), int(math.Abs(float64(root.Val-root.Left.Val))))))
	return res
}
