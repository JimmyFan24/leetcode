package weeklycontest

import "sort"

func UnequalTriplets(nums []int) int {

	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for k := len(nums) - 1; k > i+1; k-- {
			j := i + 1
			for j < k {
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					count++
				}
				j++
			}

		}
	}

	return count
}
func ClosestNodes(queries []int) [][]int {
	ans := [][]int{}
	//sort.Ints(ans)
	arr := []int{4, 9}
	sort.Ints(arr)
	for _, v := range queries {
		ans = append(ans, findValue(v, arr))
	}

	return ans
}
func preOrder(root *TreeNode) []int {
	//res := []int{}
	if root == nil {
		return nil
	}
	left := preOrder(root.Left)

	left = append(left, root.Val)
	right := preOrder(root.Right)
	for _, v := range right {
		right = append(right, v)
	}

	return left

}
func findValue(value int, arr []int) []int {
	res := make([]int, 2)
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == value {
			res[0] = value
			res[1] = value
			return res
		} else if arr[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if arr[left] == value {
		res[0] = value
		res[1] = value
		return res
	} else if arr[left] < value {
		res[0] = arr[left]
		if left+1 < len(arr) {
			res[1] = arr[left+1]
		} else {
			res[1] = -1
		}
	} else {
		res[1] = arr[left]
		if left-1 >= 0 {
			res[0] = arr[left-1]
		} else {
			res[0] = -1
		}
	}
	return res

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
