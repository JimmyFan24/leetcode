package level6

/*
An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

Given an integer n, return the nth ugly number.
*/

func NthUglyNumber(n int) int {

	dp, p2, p3, p5 := make([]int, n+1), 1, 1, 1
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CountNumbersWithUniqueDigits(n int) int {

	if n == 0 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 10

	//2.用0

	//n = 3
	//1. 9  8 [0]
	//2. 9 [0] 8   ==> 9*8 *2
	//3. 9  8  7 ==> 9*8....[9-n]
	//n=2 81
	//1. 9 8  ==> 72
	//2. 9 [0] ==> 9
	//n=1 10
	//9*(0)*8*7 *6*5*4*3

	for i := 2; i <= n; i++ {
		//非0
		count := 1
		for j := 9; j > 9-i; j-- {
			count *= j
		}
		//0
		nonzero := 1
		for q := 9; q > 9-i+1; q-- {
			nonzero *= q
		}

		dp[i] += dp[i-1] + nonzero*(i-1) + count
	}

	//9*10*9
	//504+144 = 684 + 10 = 694
	return dp[n]
}
func IntegerBreak(n int) int {

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			// dp[i] = max(dp[i], j * (i - j), j*dp[i-j])
			dp[i] = getMax(getMax(dp[i], (i-j)*j), j*dp[i-j])
		}

	}
	return dp[n]
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfsIntegerBreak(product *int, start int, factorcount int, count int, n, tmpn int, tmp []int) {

	if count == factorcount {
		if tmpn == 0 {
			tmpProduct := 1
			for i := 0; i < len(tmp); i++ {
				tmpProduct *= tmp[i]
			}
			if tmpProduct > *product {
				*product = tmpProduct
			}
			return
		}
		if tmpn < 0 {
			return
		}
	} else if count > factorcount {
		return
	}

	for i := start; i < n; i++ {
		if tmpn-i >= 0 {
			tmp = append(tmp, i)
			dfsIntegerBreak(product, i, factorcount, count+1, n, tmpn-i, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
}
