package level6

import (
	"strconv"
	"strings"
)

func Search81(nums []int, target int) bool {

	l, r := 0, len(nums)-1
	rotate := findRotate(nums)
	if nums[l] <= target {
		return rotateArrBinarySearch(l, rotate, nums, target)
	}
	return rotateArrBinarySearch(rotate+1, r, nums, target)
}
func rotateArrBinarySearch(l, r int, nums []int, target int) bool {

	if l > r {
		return false
	}
	mid := l + ((r - l) >> 1)
	if nums[mid] == target {
		return true
	} else if nums[mid] < target {

		return rotateArrBinarySearch(mid+1, r, nums, target)

	} else if nums[mid] > target {
		return rotateArrBinarySearch(l, mid-1, nums, target)
	}

	return false
}
func findRotate(nums []int) int {
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] > nums[i+1] {
			break
		}
	}
	return i
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrderBottom(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {

		l := len(queue)
		tmp := []int{}
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Left)
			}
			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[l:]
		res = append(res, tmp)
	}

	ans := [][]int{}
	for i := len(res) - 1; i >= 0; i-- {
		ans = append(ans, res[i])
	}
	return ans
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {

	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		l := len(queue)
		tmp := []*Node{}
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i])
		}
		queue = queue[l:]
		for i := 0; i < len(tmp)-1; i++ {
			tmp[i].Next = tmp[i+1]
		}
	}
	return root
}

func TwoSum167(numbers []int, target int) []int {

	ans := []int{}

	i, j := 0, len(numbers)

	for i < j {

		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return ans
}
func CompareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	i, j := 0, 0
	for i < len(v1) && j < len(v2) {
		if v1[i][0] == '0' && len(v1[i]) > 1 {
			v1[i] = removeBeforeZero(v1[i])
		}
		if v2[j][0] == '0' && len(v2[i]) > 1 {
			v2[j] = removeBeforeZero(v2[j])
		}
		vv1, _ := strconv.Atoi(v1[i])
		vv2, _ := strconv.Atoi(v2[i])
		if vv1 < vv2 {
			return -1
		} else if vv1 > vv2 {
			return 1
		}
		i++
		j++
	}

	for i < len(v1) {
		t := removeBeforeZero(v1[i])
		if t[0] != '0' {
			return 1
		}
		i++
	}

	for j < len(v2) {
		t := removeBeforeZero(v2[j])
		if t[0] != '0' {
			return -1
		}
		j++
	}
	return 0
}
func removeBeforeZero(s string) string {

	for s[0] == '0' && len(s) > 1 {
		s = s[1:]
	}
	res := s
	return res
}

func Rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}
	if k > len(nums) {
		k %= len(nums)
	}
	after := make([]int, k)
	before := make([]int, len(nums)-k)
	copy(after, nums[len(nums)-k:])
	copy(before, nums[:len(nums)-k])
	a, b := 0, 0
	i := 0
	for a < len(after) && i < len(nums) {
		nums[i] = after[a]
		a++
		i++
	}
	for b < len(before) && i < len(nums) {
		nums[i] = before[b]
		b++
		i++
	}
	return
}
func Rotate2(nums []int, k int) {

	if k == 0 || k == len(nums) {
		return
	}
	if len(nums) == 1 {
		return
	}
	if k > len(nums) {
		k %= len(nums)
	}

	//1.k 等于len(nums) 一半
	before := len(nums) - k
	after := k
	index := 0
	for before > 0 && after > 0 {
		for before > after {
			c := 0
			for j := index; j < index+after; j++ {
				nums[j], nums[len(nums)-after+c] = nums[len(nums)-after+c], nums[j]
				c++
			}
			before -= after
			index += after
		}

		if after > before {
			for after > before {
				for jj := index; jj < index+before; jj++ {
					nums[jj], nums[jj+before] = nums[jj+before], nums[jj]
				}
				after -= before
				index += before

			}
		}
		if after == before {
			for index+before < len(nums) {
				nums[index], nums[index+before] = nums[index+before], nums[index]
				index++
			}
			after -= before
			before -= after
		}
	}
	//1,2,  3,4,5 k=3

	return
}
