package weeklycontest

import (
	"sort"
)

func ApplyOperations(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}
	//1,4,0,2,0,0
	res := make([]int, len(nums))

	index := 0
	for _, v := range nums {
		if v != 0 {
			res[index] = v
			index++
		}
	}
	return res
}
func MaximumSubarraySum(nums []int, k int) int64 {

	if k == 1 {
		tmp := 0
		for _, v := range nums {
			if tmp < v {
				tmp = v
			}
		}
		return int64(tmp)
	}
	max := 0
	beforeSum := 0
	distinMap := make(map[int]bool, k)
	for i := 0; i <= len(nums)-k; i++ {

		j := i
		if beforeSum == 0 {
			sum := 0
			for j < i+k && j < len(nums) {
				sum += nums[j]
				if distinMap[nums[j]] {
					break
				}
				distinMap[nums[j]] = true
				j++
			}
			if j < i+k {
				q := i
				for q < j {
					delete(distinMap, nums[q])
					q++
				}
				sum = 0
				beforeSum = 0
			} else {
				beforeSum = sum
				if beforeSum > max {
					max = beforeSum
				}
			}
		} else {
			oldNum := nums[i-1]
			newNum := nums[i+k-1]
			delete(distinMap, oldNum)
			if !distinMap[newNum] {
				distinMap[newNum] = true
				beforeSum = beforeSum - oldNum + newNum
				if beforeSum > max {
					max = beforeSum
				}
			} else {
				//
				q := i
				for q < i+k-1 {
					delete(distinMap, nums[q])
					q++
				}
				beforeSum = 0
			}

		}

	}
	return int64(max)
}
func notContainSameNum(arr []int) (bool, int) {
	m := make(map[int]bool, len(arr))
	sum := 0
	for _, v := range arr {
		if m[v] {
			return false, 0
		}
		m[v] = true
		sum += v
	}
	return true, sum
}
func TotalCost(costs []int, k int, candidates int) int64 {

	cost := 0

	m := make(map[int][]int, len(costs))
	for i, v := range costs {
		m[v] = append(m[v], i)
	}
	c := make([]int, len(costs))
	copy(c, costs)
	sort.Ints(c)
	for i := 0; i < k; i++ {
		//index<=candidates-1] || index >=[len(cost)-candidates
		if candidates < len(c) {

			min := c[0]
			arr := make([]int, len(m[min]))
			copy(arr, m[min])
			for i, index := range arr {
				if index <= candidates-1 || index >= len(costs)-candidates {
					cost += min
					m[min] = m[min][i+1:]
					break
				}

			}
			c = c[1:]
		} else {
			min := c[0]
			m[min] = m[min][1:]
			cost += min
			c = c[1:]
		}
	}
	return int64(cost)
}
