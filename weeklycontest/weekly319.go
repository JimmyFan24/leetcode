package weeklycontest

import (
	"sort"
)

func ConvertTemperature(celsius float64) []float64 {

	res := make([]float64, 2)
	res[0] = celsius + 273.15
	res[1] = celsius*1.80 + 32.00
	return res
}
func SubarrayLCM(nums []int, k int) int {

	count := 0
	for i := 0; i < len(nums); i++ {
		if k%nums[i] == 0 {
			flag := false
			if nums[i] == k {
				flag = true
				count++
			}
			j := i + 1
			for j < len(nums) && k%nums[j] == 0 {
				if nums[j] == k {
					flag = true
				}
				if flag {
					count++
				}
				j++
			}
		}
	}
	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MinimumOperations(root *TreeNode) int {

	queue := []*TreeNode{root}
	step := 0
	if root == nil {
		return 0
	}
	for len(queue) > 0 {
		l := len(queue)
		tmp := []int{}
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}
		step += sortStep(tmp)
		queue = queue[l:]
	}
	return step
}
func sortStep(arr []int) int {

	step := 0
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	for i := 0; i < len(arr); i++ {
		if sortedArr[i] != arr[i] {
			for j := i + 1; j < len(arr); j++ {
				if sortedArr[i] == arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
					step++
				}
			}
		}
	}
	return step
}
func MaxPalindromes(s string, k int) int {

	if k == 1 {
		return len(s)
	}
	count := 0
	for i := 0; i < len(s); i++ {

		sub := ""
		for j := i; j < len(s); j++ {
			sub += string(s[i])
			if len(sub) >= k && isPalindrome(sub) {
				count++
				i = j
				break
			}
			if len(sub) > k+1 {
				break
			}
		}
	}
	return count
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func f(s string, k int, i, j int) int {

	if i < 0 {
		return 0
	}
	if j >= len(s) {
		return 0
	}
	if i < j-k+1 {
		return 0
	}
	if isPalindrome(s[i:j]) {
		return 1
	}
	return f(s, k, i, j+1) + f(s, k, i-1, j)
}
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func LongestPalindrome(s string) int {

	count := palindrome(s, 0, len(s)-1)
	return count
}
func palindrome(s string, i, j int) int {

	if i == j {
		return 1
	}
	if s[i] == s[j] {
		if isPalindrome(s[i : j+1]) {
			return palindrome(s, i+1, j-1)
		}

	}
	return palindrome(s, i+1, j)

}
