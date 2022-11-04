package level9

import (
	"math"
	"sort"
)

func MakeGood(s string) string {

	if s == "" {
		return s
	}
	arr := []byte{}
	for i := 0; i < len(s); i++ {

		if len(arr) > 0 {
			top := arr[len(arr)-1]
			if isSameLetter(top, s[i]) {
				arr = arr[:len(arr)-1]
			} else {
				arr = append(arr, s[i])
			}
		} else {
			//如果当前栈为空，直接加
			arr = append(arr, s[i])
		}

	}
	return string(arr)
}
func isLowerLetter(b byte) bool {
	return b >= 'a' && b <= 'z'
}
func isSameLetter(a, b byte) bool {

	if a == b {
		return false
	}
	return a-'a' == b-'A' || a-'A' == b-'a'
}
func SearchInsert(nums []int, target int) int {

	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid
		} else {
			r = mid - 1
		}

	}
	return l + 1
}
func MoveZeroes(nums []int) {

	i, j := 0, 0
	for i < len(nums)-1 {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
	return

}
func RemoveDuplicates(nums []int) int {

	cur, finder := 0, 0
	for cur < len(nums)-1 {
		for nums[cur] == nums[finder] {

			finder++
			if finder == len(nums)-1 {
				return cur + 1
			}

		}
		nums[cur+1] = nums[finder]
		cur++
	}
	return cur + 1
}
func MaximumProduct(nums []int) int {

	sort.Ints(nums)

	res := math.MinInt32
	//1.0*1*end

	res = max(res, nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])
	res = max(res, nums[len(nums)-1]*nums[len(nums)-2]*nums[0])
	res = max(res, nums[0]*nums[1]*nums[len(nums)-1])
	return res

}
func NumEquivDominoPairs(dominoes [][]int) int {

	count := 0
	for i := 0; i < len(dominoes)-1; i++ {
		j := i + 1
		for j < len(dominoes) {
			if isDominoes(dominoes[i], dominoes[j]) {
				count++
			}
			j++
		}
	}
	return count
}
func isDominoes(a, b []int) bool {

	return (a[0] == b[0] && a[1] == b[1]) || (a[0] == b[1] && a[1] == b[0])
}

func RemoveDigit(number string, digit byte) string {

	ans := ""
	for i := 0; i < len(number); i++ {

		if number[i] == digit {
			sub := number[:i] + number[i+1:]
			if sub > ans {
				ans = sub
			}
		}
	}
	return ans
}
