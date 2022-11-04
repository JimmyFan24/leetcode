package level9

import "math"

func MaximumGap(nums []int) int {

	max := math.MinInt32
	min := math.MaxInt32

	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	distance := max - min
	if distance == 0 {
		return 0
	}

	hasNum := make([]bool, len(nums)+1)
	maxNum := make([]int, len(nums)+1)
	minNum := make([]int, len(nums)+1)
	bid := 0

	for i := 0; i < len(nums); i++ {
		bid = (nums[i] - min) / (max - min) * len(nums)
		if hasNum[bid] {
			maxNum[bid] = getMax(maxNum[bid], nums[i])
			minNum[bid] = getMin(minNum[bid], nums[i])
		} else {
			maxNum[bid] = nums[i]
			minNum[bid] = nums[i]
			hasNum[bid] = true
		}
	}
	res := math.MinInt32
	lastMax := maxNum[0]
	for i := 1; i < len(maxNum); i++ {

		if hasNum[i] {
			res = getMax(res, minNum[i]-lastMax)
			lastMax = maxNum[i]
		}
	}
	return res
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
