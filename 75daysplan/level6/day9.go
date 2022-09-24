package level6

import (
	"math"
	"sort"
)

/*
Example 2:

Input: n = 3
Output: 27
Explanation: In binary, 1, 2, and 3 corresponds to "1", "10", and "11".
After concatenating them, we have "11011", which corresponds to the decimal value 27.
*/
func ConcatenatedBinary(n int) int {

	upLimit := int(math.Pow(float64(10), float64(9)))
	upLimit += 7
	shift := 0
	res := 0
	for i := 1; i <= n; i++ {

		if (i & (i - 1)) == 0 {
			shift++
		}
		res = ((res << shift) + i) % upLimit
	}
	return res
}

/*
Example 1:

Input: satisfaction = [-1,-8,0,5,-9]
Output: 14
Explanation: After Removing the second and last dish, the maximum total like-time coefficient will be equal to (-1*1 + 0*2 + 5*3 = 14).
Each dish is prepared in one unit of time.
*/
func MaxSatisfaction(satisfaction []int) int {

	//-9,-8,-1,0,5
	//dp[i] : 前i个菜的最满意值,
	//dp[n] : max(dp[n-1],satisfaction[n]*time)
	sort.Ints(satisfaction)
	i := 0
	for i < len(satisfaction) && satisfaction[i] <= 0 {
		i++
	}

	positiveSum, timesSum, time := 0, 0, 1
	for j := i; j < len(satisfaction); j++ {
		positiveSum += satisfaction[j]
		timesSum += satisfaction[j] * time
		time++
	}
	if i == 0 {
		return timesSum
	}
	max := 0
	for q := i - 1; q >= 0; q-- {
		positiveSum += satisfaction[q]
		timesSum = timesSum + positiveSum
		max = bigger(max, timesSum)
	}
	return max

}
func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfsMaxSatisfaction(res *int, satisfaction []int, start int, score int, time int) {

	if start >= len(satisfaction) {

		if score > *res {
			*res = score
		}
		return
	}

	for j := start; j < len(satisfaction); j++ {

		score += satisfaction[j] * time
		dfsMaxSatisfaction(res, satisfaction, j+1, score, time+1)
		score -= satisfaction[j] * time
	}
	return
}
func CanSeePersonsCount(heights []int) []int {

	stack := []int{}

	res := make([]int, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {

		for len(stack) > 0 && heights[i] > stack[len(stack)-1] {
			res[i]++
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i]++
		}
		stack = append(stack, heights[i])

	}
	return res
}
