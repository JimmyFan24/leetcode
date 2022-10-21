package level8

import (
	"math"
	"sort"
)

func DestroyTargets(nums []int, space int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	m := make(map[int]int, len(nums))
	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]

	for _, v := range nums {
		m[v]++
	}
	maxFreq := 0
	maxIndex := []int{}
	for _, v := range m {
		if v > maxFreq {
			maxFreq = v
		}
	}
	for k, v := range m {
		if v == maxFreq {
			maxIndex = append(maxIndex, k)
		}
	}
	sort.Ints(maxIndex)
	if space >= max || min+space > max {
		return maxIndex[0]
	}
	res := math.MinInt32
	num := 0
	for i := 0; i < len(nums)-1; i++ {
		//count := 0
		//nums[i] + c * space
		j := i + 1
		for ; j < len(nums)-1; j++ {
			count := 0
			for j < len(nums)-1 && nums[j] == nums[i] {
				j++
			}
			c := nums[j] - nums[i]
			q := j
			for q < len(nums)-1 && nums[q]-nums[q-1] == c {
				q++
			}
			count = q - i + 1

			if count > res {
				num = nums[i]
				res = count
			}

		}

	}

	return num
}
func countDestory(i int, nums []int, space int, c int, m map[int]int) int {
	max := nums[len(nums)-1]
	count := 0

	for j := nums[i]; j <= max; j = nums[i] + c*space {
		if m[j] > 0 {
			count += m[j]
		}
	}

	return count
}
