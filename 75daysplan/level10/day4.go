package level10

import "strings"

func MostFrequentEven(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if v%2 == 0 {
			m[v]++
		}
	}
	res := 100001
	mostFreq := 0
	if len(m) == 0 {
		return -1
	}
	for num, freq := range m {
		if freq > mostFreq {
			res = num
		} else if freq == mostFreq {
			if res > num {
				res = num

			}
		}
	}
	return res
}
func LargeGroupPositions(s string) [][]int {

	res := [][]int{}
	for i := 0; i < len(s); i++ {
		j := i + 1
		for j < len(s) && s[j] == s[i] {
			j++
		}
		if j-i >= 3 {
			res = append(res, []int{i, j - 1})
		}
	}
	return res
}

func FindSubarrays(nums []int) bool {

	// two subarrays of length 2 with equal sum
	for i := 0; i < len(nums)-1; i++ {
		j := i + 1
		sum := nums[i] + nums[j]
		for q := j; q < len(nums)-1; q++ {
			k := q + 1
			if nums[k]+nums[q] == sum {
				return true
			}

		}
	}
	return false
}
func StringMatching(words []string) []string {

	res := []string{}
	for i := 0; i < len(words); i++ {
		j := i + 1
		for ; j < len(words); j++ {
			if strings.Contains(words[i], words[j]) {

				res = append(res, words[j])
			}
			if strings.Contains(words[j], words[i]) {

				res = append(res, words[i])
			}
		}
	}
	return res
}
