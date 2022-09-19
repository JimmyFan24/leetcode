package level6

import (
	"math"
	"sort"
	"strings"
)

/*
Example 1:

Input: paths = ["root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)","root 4.txt(efgh)"]
Output: [["root/a/2.txt","root/c/d/4.txt","root/4.txt"],["root/a/1.txt","root/c/3.txt"]]
*/
func FindDuplicate(paths []string) [][]string {

	pathMap := make(map[string][]string, len(paths))
	for _, path := range paths {
		p := strings.Split(path, " ")
		filerootpath := p[0]
		filepaths := []string{}
		content := []string{}
		count := 1
		for count < len(p) {

			filename := strings.Split(p[count], "(")
			filep := filerootpath + "/" + filename[0]
			filepaths = append(filepaths, filep)
			content = append(content, filename[1][:len(filename[1])-1])
			count++
		}

		for i, c := range content {
			if _, ok := pathMap[c]; !ok {
				pathMap[c] = []string{}
				pathMap[c] = append(pathMap[c], filepaths[i])
			} else {
				pathMap[c] = append(pathMap[c], filepaths[i])
			}
		}

	}
	res := [][]string{}
	for _, v := range pathMap {
		if len(v) > 1 {
			res = append(res, v)
		}

	}
	return res
}
func PermuteUnique(nums []int) [][]int {
	res := [][]int{}
	visited := make([]bool, len(nums))
	sort.Ints(nums)

	dfsPermuteUnique(&visited, &res, nums, 0, len(nums), 0, []int{})

	return res
	//dp[i] 表示从0到i的全排列,那么dp[i+1] = dp[i] 全排列 + nums[i+1]

}
func dfsPermuteUnique(visited *[]bool, res *[][]int, nums []int, start int, end, count int, tmp []int) {

	if end == count {
		c := make([]int, len(tmp))
		copy(c, tmp)
		*res = append(*res, c)
		return
	}

	for j := 0; j < len(nums); j++ {

		if !(*visited)[j] {
			if j > 0 && nums[j] == nums[j-1] && !(*visited)[j-1] { // 这里是去重的关键逻辑
				continue
			}
			(*visited)[j] = true
			tmp = append(tmp, nums[j])
			dfsPermuteUnique(visited, res, nums, j+1, end, count+1, tmp)
			tmp = tmp[:len(tmp)-1]
			(*visited)[j] = false
		}

	}

}

/*
Example 1:

Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/
func CountPrimes(n int) int {
	isNotPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		//判断当前遍历的数时候为质数，如果不是质数就遍历下一个数
		if isNotPrime[i] {
			continue
		}
		///排除当前这个质数的2n、3n倍数..
		for j := i + i; j < n; j = j + i {
			isNotPrime[j] = true
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			count++
		}
	}
	return count

}

func NumSquares(n int) int {

	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}

	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}

	return 3
}

// 判断是否为完全平方数
func isPerfectSquare(n int) bool {
	sq := int(math.Floor(math.Sqrt(float64(n))))
	return sq*sq == n
}
func checkAnswer4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}

/*
Example 1:

Input: prices = [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]
*/
func MaxProfit(prices []int) int {

	if len(prices) == 1 {
		return 0
	}
	max := math.MinInt32
	min := math.MaxInt32
	i := 0
	j := len(prices) - 1
	for i < j {
		if max < prices[i] {
			max = prices[i]
		}
		if min > prices[j] {
			min = prices[j]
		}
		i++
		j--
	}
	profit := 0

	for i := 0; i < len(prices)-1; i++ {

		j := i
		//第一个最高点
		for j+1 < len(prices) && prices[j+1] >= prices[j] {
			j++
		}
		q := j
		//第一个最低点
		for q+1 < len(prices) && prices[q+1] < prices[q] {
			q++
		}
		sellday := 0
		if q-j == 1 && q < len(prices)-1 {
			sellday = j - 1
		} else {
			sellday = j
		}

		for i < sellday {
			profit += prices[i+1] - prices[i]
			i++
		}
		i = q - 1

	}
	if profit < max-min {
		profit = max - min
	}
	return profit
}
