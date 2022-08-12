package level3

import "strconv"

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

func findLUSlength(a string, b string) int {

	tmp := 9
	return tmp
}
