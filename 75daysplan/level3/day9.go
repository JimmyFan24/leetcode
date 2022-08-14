package level3

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
Example 1:

Input: heights = [1,1,4,2,1,3]
Output: 3
Explanation:
heights:  [1,1,4,2,1,3]
expected: [1,1,1,2,3,4]
Indices 2, 4, and 5 do not match.
*/

func HeightChecker(heights []int) int {

	tmp := make([]int, len(heights))
	copy(tmp, heights)
	sort.Ints(tmp)
	count := 0
	for i, v := range heights {
		if v != tmp[i] {
			count++
		}
	}
	return count
}

/*
Example 1:

Input: s = "abbaca"
Output: "ca"
Explanation:
For example, in "abbaca" we could remove "bb" since
the letters are adjacent and equal, and this is the only possible move.
The result of this move is that the string is "aaca", of which only "aa" is possible,
so the final string is "ca"
*/
func RemoveDuplicates(s string) string {
	//"abbaca"

	for i := 0; i < len(s)-1; {
		j := i + 1
		if s[i] == s[j] {
			if j+1 < len(s) {
				s = s[:i] + s[j+1:]
				i = 0
			} else {
				s = s[:i]
			}
		} else {
			i++
		}

	}
	return s
}

/*
Example 1:

Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
*/
func LastStoneWeight(stones []int) int {

	for len(stones) > 1 {
		sort.Ints(stones)
		if stones[len(stones)-1] != stones[len(stones)-2] {
			first := stones[len(stones)-1]
			second := stones[len(stones)-2]
			stones = stones[:len(stones)-2]
			stones = append(stones, first-second)
		} else {
			stones = stones[:len(stones)-2]
		}
	}

	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

/*
A triplet (arr[i], arr[j], arr[k]) is good if the following conditions are true:

0 <= i < j < k < arr.length
|arr[i] - arr[j]| <= a
|arr[j] - arr[k]| <= b
|arr[i] - arr[k]| <= c
*/
func CountGoodTriplets(arr []int, a int, b int, c int) int {

	i, k := 0, len(arr)-1
	j := i + 1
	count := 0
	for j < k {

		if getAbsInt(arr[i], arr[j]) <= a && getAbsInt(arr[j], arr[k]) <= b && getAbsInt(arr[i], arr[k]) <= c {
			count++
			j++
		} else {
			k--
			j = i + 1
		}
	}
	return count
}
func getAbsInt(a, b int) int {

	return int(math.Abs(float64(a - b)))
}

/*
Example 1:

Input: nums = [5,6,2,7,4]
Output: 34
Explanation: We can choose indices 1 and 3 for the first pair (6, 7)
and indices 2 and 4 for the second pair (2, 4).
The product difference is (6 * 7) - (2 * 4) = 34.
*/
func MaxProductDifference(nums []int) int {
	sort.Ints(nums)
	big := nums[len(nums)-1] * nums[len(nums)-2]
	small := nums[0] * nums[1]
	return big - small
}

/*
Example 1:

Input: s = "abccbaacz"
Output: "c"
Explanation:
The letter 'a' appears on the indexes 0, 5 and 6.
The letter 'b' appears on the indexes 1 and 4.
The letter 'c' appears on the indexes 2, 3 and 7.
The letter 'z' appears on the index 8.
The letter 'c' is the first letter to appear twice,
because out of all the letters the index of its second occurrence is the smallest.
*/
func RepeatedCharacter(s string) byte {
	//abccbaacz
	sMap := make(map[byte]int, 26)
	min := math.MaxInt32
	var res byte
	for i := 0; i < len(s); i++ {

		if sMap[s[i]] > 0 {
			if i < min {
				res = s[i]
				min = i

			}

		}
		sMap[s[i]]++
	}
	return res
}

/*
Example 1:

Input: paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
Output: "Sao Paulo"
Explanation: Starting at "London" city you will reach "Sao Paulo" city which is the destination city.
Your trip consist of: "London" -> "New York" -> "Lima" -> "Sao Paulo".
*/
/*
[["pYyNGfBYbm","wxAscRuzOl"],["kzwEQHfwce","pYyNGfBYbm"]]
*/
func DestCity(paths [][]string) string {
	pMap := make(map[string]int, len(paths)*len(paths[0]))
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[0]); j++ {
			pMap[paths[i][j]]++
		}
	}
	res := ""
	for i := 0; i < len(paths); i++ {
		if pMap[paths[i][len(paths[0])-1]] == 1 {
			res = paths[i][len(paths[0])-1]
			break

		}
	}
	return res
}

/*
You are given an integer n.

Each number from 1 to n is grouped according to the sum of its digits.

Return the number of groups that have the largest size.
*/
/*

Example 1:

Input: n = 13
Output: 4
Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
[1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9].
There are 4 groups with largest size.
*/
func CountLargestGroup(n int) int {

	nMap := make(map[int]int, n)
	for i := 1; i <= n; i++ {

		sum := 0
		tmp := i
		for tmp > 0 {
			sum += tmp % 10
			tmp = tmp / 10
		}

		nMap[sum]++
	}

	max := math.MinInt32
	for _, v := range nMap {
		if v > max {
			max = v
		}
	}
	count := 0
	for _, v := range nMap {
		if v == max {
			count++
		}
	}
	return count
}

/*
Example 1:

Input: n = 34, k = 6
Output: 9
Explanation: 34 (base 10) expressed in base 6 is 54. 5 + 4 = 9.
*/
func SumBase(n int, k int) int {

	tmp := ""
	for n > 0 {

		tmp = tmp + strconv.Itoa(n%k)
		n = n / k
	}
	sum := 0
	for _, v := range tmp {
		i, _ := strconv.Atoi(string(v))
		sum += i
	}
	return sum
}

/*
Example 1:

Input: gain = [-5,1,5,0,-7]
Output: 1
Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.
*/
func LargestAltitude(gain []int) int {
	max := 0
	sum := 0
	for _, v := range gain {
		sum += v
		if sum > max {
			max = sum
		}

	}
	return sum
}
func MaximumUnits(boxTypes [][]int, truckSize int) int {

	max := 0
	value := []int{}
	vMap := make(map[int]int, len(boxTypes))
	for _, v := range boxTypes {
		value = append(value, v[1])
		if vMap[v[1]] > 0 {
			vMap[v[1]] += v[0]
		} else {
			vMap[v[1]] = v[0]
		}

	}
	sort.Ints(value)

	newval := []int{}
	for k, _ := range vMap {
		newval = append(newval, k)
	}
	sort.Ints(newval)
	for i := len(newval) - 1; i >= 0; i-- {

		if truckSize >= vMap[newval[i]] {
			max += newval[i] * vMap[newval[i]]
			truckSize -= vMap[newval[i]]
		} else {
			max += truckSize * newval[i]
			break
		}

	}
	return max

}

/*
A string is good if there are no repeated characters.

Given a string s​​​​​,
return the number of good substrings of length three in s​​​​​​.
Example 1:

Input: s = "xyzzaz"
Output: 1
Explanation: There are 4 substrings of size 3: "xyz", "yzz", "zza", and "zaz".
The only good substring of length 3 is "xyz".
*/
func CountGoodSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s)-2; i++ {

		if NotDupStr(s[i : i+3]) {
			count++
		} else {
			continue
		}

	}
	return count
}
func NotDupStr(s string) bool {

	m := make(map[rune]int, 3)
	for _, v := range s {
		if m[v] > 0 {
			return false
		} else {
			m[v]++
		}
	}
	return true
}

/*
Example 1:

Input: text = "hello world", brokenLetters = "ad"
Output: 1
Explanation: We cannot type "world" because the 'd' key is broken.
*/
func CanBeTypedWords(text string, brokenLetters string) int {

	count := 0
	sList := strings.Split(text, " ")
	for _, text := range sList {
		if !containKey(text, brokenLetters) {
			count++
		}
	}
	return count
}
func containKey(str string, brokenLetters string) bool {
	mStr := make(map[rune]int, 26)
	for _, v := range str {
		mStr[v]++
	}
	for _, v := range brokenLetters {
		if mStr[v] > 0 {
			return true
		}
	}
	return false

}

/*
Example 1:

Input: arr = ["d","b","c","b","c","a"], k = 2
Output: "a"
Explanation:
The only distinct strings in arr are "d" and "a".
"d" appears 1st, so it is the 1st distinct string.
"a" appears 2nd, so it is the 2nd distinct string.
Since k == 2, "a" is returned.
*/
func kthDistinct(arr []string, k int) string {

	res := ""
	m := make(map[string]int, len(arr))
	for _, v := range arr {

		m[v]++
	}
	for _, v := range arr {
		if k == 1 {
			res = v
			break
		}
		if m[v] > 0 {
			k--
		}
	}
	return res
}

func SumRootToLeaf(root *TreeNode) int {

	return 0
}
