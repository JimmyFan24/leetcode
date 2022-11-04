package weeklycontest

import "sort"

func DistinctAverages(nums []int) int {

	if len(nums) == 2 {
		return 1
	}
	sort.Ints(nums)
	m := make(map[int]int)
	for len(nums) > 1 {
		minOne := nums[0]
		maxOne := nums[len(nums)-1]
		m[minOne+maxOne]++
		nums = nums[1 : len(nums)-1]
	}
	return len(m)
}
func CountGoodStrings(low int, high int, zero int, one int) int {

	n := high / zero

	m := high / one

	// "00", "11", "000","110".,"011",
	count := 0

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			tmp := zero*i + one*j
			if tmp >= low && tmp <= high {
				count += Count(i, j, zero, one)
			}
		}
	}

	return count % (100000007)
}
func Count(i, j int, zero, one int) int {

	if i == 0 || j == 0 {
		return 1
	}

	sum := i + j
	oneCount := j
	zeroCOunt := i
	tmp := 1
	for q := 0; q < sum; q++ {
		tmp *= sum
		sum--
	}
	t := 1
	for p := 0; p < i; p++ {
		t *= zeroCOunt
		zeroCOunt--
	}
	tt := 1
	for k := 0; k < j; k++ {
		tt *= oneCount
		oneCount--
	}

	return tmp / (t * tt)
}
