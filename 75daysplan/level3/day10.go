package level3

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
Example 1:

Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
*/
func DuplicateZeros(arr []int) {

	for i := 0; i < len(arr); {

		if arr[i] == 0 {
			j := len(arr) - 1
			for ; j > i; j-- {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
			arr[j] = 0
			i = j + 2
		} else {
			i++
		}
	}
}

/*
Example 1:

Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
Output: ["girl","student"]
*/

/*
"we will we will rock you"
"we"
"will"
Output
["we"]
Expected
["we","rock"]
*/
func FindOcurrences(text string, first string, second string) []string {
	sList := strings.Split(text, " ")

	res := []string{}
	for i := 0; i < len(sList)-2; i++ {
		if sList[i] == first {
			if sList[i+1] == second {
				res = append(res, sList[i+2])

			}
		}
	}
	return res
}

/*
Input: colors = [1,1,1,6,1,1,1]
Output: 3
Explanation: In the above image, color 1 is blue, and color 6 is red.
The furthest two houses with different colors are house 0 and house 3.
Input: colors = [1,8,3,8,3]
Output: 4
Explanation: In the above image, color 1 is blue, color 8 is yellow, and color 3 is green.
The furthest two houses with different colors are house 0 and house 4.
House 0 has color 1, and house 4 has color 3. The distance between them is abs(0 - 4) = 4.
*/
func MaxDistance(colors []int) int {
	max := math.MinInt32
	for i := 0; i < len(colors); i++ {

		j := i + 1
		for ; j < len(colors); j++ {
			if colors[j] != colors[i] {
				max = getMinHouse(max, j-i)
			}
		}

	}
	return max
}
func getMinHouse(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Given an integer array arr,
return the mean of the remaining integers after removing the smallest 5% and the largest 5% of the elements.
Answers within 10-5 of the actual answer will be considered accepted.
*/
func TrimMean(arr []int) float64 {
	l := len(arr)
	rl := float64(l) * 0.05

	sort.Ints(arr)
	arr = arr[int(rl) : len(arr)-int(rl)]
	sum := 0
	for _, v := range arr {
		sum += v
	}
	avr := float64(sum) / float64(len(arr))
	return avr

}

/*
Example 1:
In one operation
you can increase the time current by 1, 5, 15, or 60 minutes.
You can perform this operation any number of times.
Input: current = "02:30", correct = "04:35"
Output: 3
Explanation:
We can convert current to correct in 3 operations as follows:
- Add 60 minutes to current. current becomes "03:30".
- Add 60 minutes to current. current becomes "04:30".
- Add 5 minutes to current. current becomes "04:35".
It can be proven that it is not possible to convert current to correct in fewer than 3 operations.
*/
func ConvertTime(current string, correct string) int {

	//current = "02:30", correct = "04:35"

	corrextList := strings.Split(correct, ":")
	currentList := strings.Split(current, ":")

	hourleft := getHour(currentList[0], corrextList[0])
	minuteleft := getMinute(currentList[1], corrextList[1])
	return hourleft + minuteleft

}
func getMinute(string1 string, string2 string) int {
	if string1 == string2 {
		return 0
	}
	if string1[0] == '0' {
		string1 = string1[1:]
	}
	if string2[0] == '0' {
		string2 = string2[1:]
	}
	s1, _ := strconv.Atoi(string1)
	s2, _ := strconv.Atoi(string2)
	count := 0

	if s1 > s2 {
		s2 += 60
		count = -1
	}
	for s2 != s1 {

		if s2-s1 == 60 {
			s1 += 60
			count++
		} else if s2-s1 >= 15 {
			s1 += 15
			count++
		} else if s2-s1 >= 5 {
			s1 += 5
			count++
		} else {
			s1++
			count++
		}
	}

	return count
}
func getHour(string1 string, string2 string) int {

	if string1[0] == '0' {
		string1 = string1[1:]
	}
	if string2[0] == '0' {
		string2 = string2[1:]
	}
	s1, _ := strconv.Atoi(string1)
	s2, _ := strconv.Atoi(string2)
	if s1 < s2 {
		return s2 - s1
	}
	return s2 - s1 + 12
}

/*
Example 1:

Input: sentence = "i love eating burger", searchWord = "burg"
Output: 4
Explanation: "burg" is prefix of "burger" which is the 4th word in the sentence.
*/
func IsPrefixOfWord(sentence string, searchWord string) int {

	count := 0
	sList := strings.Split(sentence, " ")
	for i, v := range sList {
		if findPrefix(v, searchWord) {
			count = i
			break
		}
	}
	return count
}
func findPrefix(word string, pre string) bool {

	if len(word) < len(pre) {
		return false
	}
	for i := 0; i < len(pre); i++ {
		if pre[i] != word[i] {
			return false
		}
	}
	return true
}

/*
Input: s = "abcdefghi", k = 3, fill = "x"
Output: ["abc","def","ghi"]
*/
func DivideString(s string, k int, fill byte) []string {

	res := []string{}

	for len(s) > 0 {
		tmp := ""
		if len(s) >= k {
			tmp = s[:k]
			res = append(res, tmp)
			s = s[k:]
		} else {
			tmp = s
			for i := 0; i < k-len(s); i++ {
				tmp += string(fill)
			}
			res = append(res, tmp)
			s = ""
		}

	}
	return res

}
