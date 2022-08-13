package level3

import (
	"strconv"
	"strings"
)

func CountMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	key := getType(ruleKey)
	for _, item := range items {
		if key != -1 {

			if ruleValue == item[key] {
				res++
			}
		}

	}
	return res
}
func getType(ruleKey string) int {

	switch ruleKey {
	case "type":
		return 0
	case "color":
		return 1
	case "name":
		return 2
	default:
		return -1

	}
}

/*
You are given a 0-indexed, strictly increasing integer array nums and a positive integer diff. A triplet (i, j, k) is an arithmetic triplet if the following conditions are met:

i < j < k,
nums[j] - nums[i] == diff, and
nums[k] - nums[j] == diff
*/
func ArithmeticTriplets(nums []int, diff int) int {

	res := 0

	//j-i = k-j
	//k+i ==2*j
	for j := 1; j < len(nums)-1; j++ {
		// nums = [0,1,4,6,7,10], diff = 3
		i, k := 0, len(nums)-1
		for j > i && k > j {
			if nums[i]+nums[k] == nums[j]*2 {
				if nums[k]-nums[j] == diff {
					res++

				}
				i++
				k--
			} else if nums[k]+nums[i] < nums[j]*2 {
				i++
			} else {
				k--
			}
		}

	}

	return res
}

func DecodeMessage(key string, message string) string {
	kMap := make(map[rune]rune, 26)

	i := 0
	for _, v := range key {
		if v == ' ' {
			continue
		}
		if kMap[v] > 0 {
			continue
		}
		kMap[v] = rune('a' + i)
		i++
	}

	res := ""
	for _, v := range message {

		if v == ' ' {
			res = res + string(' ')
		} else {
			res = res + string(kMap[v])
		}

	}
	return res
}

/*
Given an array of positive integers arr, return the sum of all possible odd-length subarrays of arr.

A subarray is a contiguous subsequence of the array.
*/
func SumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	//1
	i := 1
	res := 0
	for ; i <= n; i = i + 2 {
		for j := 0; j < n-i+1; j++ {
			for p := j; p < j+i; p++ {
				res += arr[p]
			}

		}

	}
	return res
}
func sumDfs(arr []int, start int, index int, count int, res *int) {
	//1,4,2,5,3
	if index == count {
		return
	}

	for j := start; j < len(arr); j++ {
		*res = *res + arr[j]
		sumDfs(arr, start+1, index, count+1, res)
	}
	return
}
func XorOperation(n int, start int) int {

	res := 0
	for i := 0; i < n; i++ {
		res = res ^ (start + 2*i)
	}
	return res

}

/*
Input: s = "K1:L2"
Output: ["K1","K2","L1","L2"]
Explanation:
The above diagram shows the cells which should be present in the list.
The red arrows denote the order in which the cells should be presented.
*/

func CellsInRange(s string) []string {
	sList := strings.Split(s, ":")

	begin, _ := strconv.Atoi(string(sList[0][1]))
	end, _ := strconv.Atoi(string(sList[1][1]))

	res := []string{}
	if begin == end {
		gap := sList[1][0] - sList[0][0]
		for i := 0; i <= int(gap); i++ {
			b := sList[0][0]
			b = b + uint8(i)
			res = append(res, string(b)+string(sList[0][1]))
		}
	} else if begin > end {
		res = append(res, string(sList[0]))
		for i := 1; i <= end; i++ {
			b := string(sList[1][0]) + strconv.Itoa(i)
			res = append(res, b)
		}
	} else {
		for p := sList[0][0]; p <= sList[1][0]; p++ {
			for i := begin; i <= end; i++ {
				b := string(p) + strconv.Itoa(i)
				res = append(res, b)
			}
		}
	}
	return res

}

func FindCenter(edges [][]int) int {

	eMap := make(map[int]int, 2)
	for _, edge := range edges {

		eMap[edge[0]]++
		eMap[edge[1]]++
	}
	res := 0
	for _, egde := range edges {

		if eMap[egde[0]] == len(edges) {
			res = eMap[egde[0]]
		} else if eMap[egde[1]] == len(edges) {
			res = eMap[egde[1]]
		}
	}
	return res
}

func NumberOfMatches(n int) int {
	match := 0

	for n > 1 {
		if n%2 != 0 {
			match += (n - 1) / 2
			n = (n-1)/2 + 1

		} else {
			match += n / 2
			n /= 2
		}

	}
	return match
}
func CountConsistentStrings(allowed string, words []string) int {
	count := 0
	for _, word := range words {

		if isConsistenStrings(allowed, word) {
			count++
		}
	}
	return count
}
func isConsistenStrings(allowed string, word string) bool {

	aMap := make(map[rune]int, len(allowed))
	for _, v := range allowed {
		aMap[v]++
	}
	for _, v := range word {
		if _, ok := aMap[v]; !ok {
			return false
		}
	}
	return true
}
func CountPairs(nums []int, k int) int {
	count := 0
	i := 0
	for ; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				if i*j%2 == 0 {
					count++
				}
			}
		}

	}
	return count
}
func UniqueMorseRepresentations(words []string) int {
	count := 0
	morseMap := make(map[string]int, len(words))
	for _, v := range words {
		morse := getMorseCode(v)
		morseMap[morse]++
	}
	for _, v := range morseMap {
		if v > 0 {
			count++
		}
	}
	return count
}

func getMorseCode(word string) string {
	MorseToLetter := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	res := ""
	for _, w := range word {
		res += MorseToLetter[int(w-'a')]
	}
	return res
}
