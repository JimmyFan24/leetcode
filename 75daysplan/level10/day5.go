package level10

import (
	"sort"
	"strconv"
)

func DivisorGame(n int) bool {

	//Choosing any x with 0 < x < n and n % x == 0.
	//Replacing the number n on the chalkboard with n - x.
	//Alice starting first.
	return n%2 == 0
}
func DigitSum(s string, k int) string {

	for len(s) > k {
		s = divideStr(s, k)
	}
	return s
}
func divideStr(s string, k int) string {

	res := ""
	for i := 0; i < len(s); i++ {
		index := i
		num := 0
		for index < i+k && index < len(s) {
			num += int(s[index] - '0')
			index++
		}
		res += strconv.Itoa(num)
		i = index - 1
	}
	return res
}
func SortEvenOdd(nums []int) []int {
	oddNums := []int{}
	evenNums := []int{}
	for i, v := range nums {
		if i%2 == 0 {
			evenNums = append(evenNums, v)
		} else {
			oddNums = append(oddNums, v)
		}
	}
	sort.Ints(evenNums)
	sort.Ints(oddNums)
	oddIndex := len(oddNums) - 1
	evenIndex := 0
	res := make([]int, len(nums))
	for i, _ := range res {
		if i%2 == 0 {
			res[i] = evenNums[evenIndex]
			evenIndex--
		} else {
			res[i] = oddNums[oddIndex]
			oddIndex--
		}

	}
	return res
}
