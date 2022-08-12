package level2

import "sort"

func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int, len(nums1)*len(nums2))

	for _, a := range nums1 {
		for _, b := range nums2 {
			m[a+b]++
		}
	}

	res := 0
	for _, c := range nums3 {
		for _, d := range nums4 {

			res += m[0-c-d]
		}
	}
	return res

}
func Reverse(x int) int {

	tmp := 0
	// tmp = 0 + -3 = -3 x = -12
	// tmp = -3*10 + -12%10 = -30 + -2 = -32 x = -1
	//tmp = -32*10 + -1%10 = -32
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10

	}

	return tmp
}

func NumMatchingSubseq(s string, words []string) int {

	sMap := make(map[rune]int, len(s))
	for _, v := range s {
		sMap[v]++
	}
	res := 0
	for j := 0; j < len(words); j++ {
		tmp := sMap
		for i, v := range words[j] {

			if _, ok := tmp[v]; !ok {
				break
			} else if tmp[v] >= 0 {
				tmp[v]--
				if i == len(words[j])-1 {
					res++
				}

			} else {
				break
			}

		}

	}
	return res
}

func myAtoi(s string) int {

	res := 0
	minus := false
	plus := false
	stack := []rune{}
	for i, w := range s {

		if w == '-' {
			if len(stack) == 0 && !plus {
				if minus {
					break
				} else {
					minus = true
				}

			} else if len(stack) == 0 && plus {
				break
			} else if len(stack) > 0 {
				break
			}
		}
		if w == '+' {
			if len(stack) == 0 && !minus {
				if i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
					plus = true
				} else {
					break
				}
			} else if len(stack) == 0 && minus {
				break
			} else if len(stack) > 0 {
				break
			}
		}

		if w == ' ' {
			if len(stack) != 0 {
				break
			} else {
				continue
			}
		}

		if w >= '0' && w <= '9' || w == '.' {
			stack = append(stack, w)
		}
		if w >= 'a' && w <= 'z' || w >= 'A' && w <= 'Z' {
			break
		}

	}

	for _, v := range stack {

		if v == '0' {
			if res == 0 {
				continue
			}
		}
		if v == '.' {
			if res == 0 {
				return 0
			} else {
				break
			}
		}
		res = res*10 + int(v-'0')

		if res >= (1 << 31) {
			res = (1 << 31)
		} else {
			res = res
		}
	}

	if minus {
		if res >= (1 << 31) {
			res = 0 - res
		} else {
			res = 0 - res
		}

	} else {
		if res >= (1 << 31) {
			res = res - 1
		} else {
			res = res
		}
	}
	return res
}

func ThreeSum(nums []int) [][]int {
	res := [][]int{}

	counter := make(map[int]int, len(nums))
	for _,v := range nums{
		counter[v]++
	}
	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)
	//[-4,-1,0,1,2]

	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i+1; j <len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			//[-4,-1,0,1,2]
			//-1,1,0
			c := 0 - uniqNums[i] - uniqNums[j]
			if c< uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}

		}
	}
	return res

}

func MaxProfit(prices []int) int {

	min := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {

		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}
	return maxProfit
}
