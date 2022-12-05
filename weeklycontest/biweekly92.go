package weeklycontest

import "math"

func NumberOfCuts(n int) int {

	if n%2 == 0 {
		return n / 2
	}
	return n
}
func OnesMinusZeros(grid [][]int) [][]int {
	diff := make([][]int, len(grid))
	for i := 0; i < len(diff); i++ {
		diff[i] = make([]int, len(grid[0]))
	}

	//onesRow0 + onesCol0 - zerosRow0 - zerosCol0
	onesRow := []int{}
	onesCol := []int{}
	for i := 0; i < len(grid); i++ {
		count := 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
		onesRow = append(onesRow, count)
	}
	for i := 0; i < len(grid[0]); i++ {
		count := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 1 {
				count++
			}
		}
		onesCol = append(onesCol, count)
	}
	for i := 0; i < len(onesRow); i++ {
		onesRowi := onesRow[i]
		for j := 0; j < len(onesCol); j++ {
			diff[i][j] = 2*onesRowi - len(grid[0]) + 2*onesCol[j] - len(grid)
		}

	}

	return diff

}

func BestClosingTime(customers string) int {

	dp := make([]int, len(customers)+1)
	totalY := 0
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'Y' {
			totalY++
		}
	}
	totolN := len(customers) - totalY
	dp[0] = totalY
	dp[len(customers)] = totolN
	for i := 1; i < len(customers); i++ {
		if customers[i-1] == 'Y' {
			//如果前一个时间点是Y但是当前的时间点也是Y，那么penty+1,但是如果是N
			dp[i] = dp[i-1] - 1

		} else {
			//如果前一个时间点是N,现在也是N
			dp[i] = dp[i-1] + 1

		}

	}
	res := math.MaxInt32
	index := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] < res {
			index = i
			res = dp[i]
		}
	}
	return index
}
func countPenty(customers string, time int) int {

	penty := 0
	for i := 0; i < len(customers); i++ {
		//For every hour when the shop is open and no customers come, the penalty increases by 1.
		if customers[i] == 'N' && i < time {
			penty++
		}
		//For every hour when the shop is closed and customers come, the penalty increases by 1.
		if i >= time && customers[i] == 'Y' {
			penty++
		}

	}
	return penty
}
func CountPalindromes(s string) int {

	res := 0

	dfsPalindromes(s, 0, &res, "")
	return res
}
func dfsPalindromes(s string, start int, res *int, tmp string) {

	if len(tmp) == 5 {
		if isPalindromes(tmp) {
			*res++
		}
		return
	}
	for j := start; j < len(s); j++ {

		tmp += string(s[j])
		dfsPalindromes(s, j+1, res, tmp)
		tmp = tmp[:len(tmp)-1]
	}
	return
}
func isPalindromes(s string) bool {
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
