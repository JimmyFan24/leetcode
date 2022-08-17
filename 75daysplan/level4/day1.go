package level4

import (
	"math"
	"sort"
	"strconv"
)

func NumOfStrings(patterns []string, word string) int {

	count := 0
	for _, v := range patterns {
		if isSubset(v, word) {
			count++
		}
	}
	return count
}
func isSubset(str string, word string) bool {

	if len(str) > len(word) {
		return false
	}
	if str == word {
		return true
	}

	for j := 0; j < len(word)-len(str)+1; j++ {
		if str[0] == word[j] && str == word[j:j+len(str)] {
			return true
		}
	}

	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func EvaluateTree(root *TreeNode) bool {

	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return false
		}
		return true
	}
	if root.Left != nil && root.Right != nil {
		if root.Val == 2 {
			return EvaluateTree(root.Left) || EvaluateTree(root.Right)
		} else if root.Val == 3 {
			return EvaluateTree(root.Left) && EvaluateTree(root.Right)
		}
	}
	return false
}

func DiagonalSum(mat [][]int) int {

	sum := 0
	m := len(mat)
	n := len(mat[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				sum += mat[i][j]
			}
			if i != j && i+j == m-1 {
				sum += mat[i][j]
			}
		}
	}
	return sum
}

/*
Example 1:

Input: num = 9669
Output: 9969
Explanation:
Changing the first digit results in 6669.
Changing the second digit results in 9969.
Changing the third digit results in 9699.
Changing the fourth digit results in 9666.
The maximum number is 9969.
*/
func Maximum69Number(num int) int {

	str := strconv.Itoa(num)

	res := ""
	for i, v := range str {
		if v == '6' {
			res = res + string('9')
			res = res + str[i+1:]
			break
		} else {
			res = res + string(v)
		}
	}
	n, _ := strconv.Atoi(res)
	return n
}

/*
Example 1:

Input: s = "foobar", letter = "o"
Output: 33
Explanation:
The percentage of characters in s that equal the letter 'o' is 2 / 6 * 100% = 33% when rounded down, so we return 33.
*/
func PercentageLetter(s string, letter byte) int {

	m := make(map[rune]int, 26)
	for _, v := range s {
		m[v]++
	}
	tmp := float64(m[rune(letter)]) / float64(len(s)) * 100
	return int(tmp)
}

/*
Example 1:

Input: firstWord = "acb", secondWord = "cba", targetWord = "cdb"
Output: true
Explanation:
The numerical value of firstWord is "acb" -> "021" -> 21.
The numerical value of secondWord is "cba" -> "210" -> 210.
The numerical value of targetWord is "cdb" -> "231" -> 231.
We return true because 21 + 210 == 231.
*/
func IsSumEqual(firstWord string, secondWord string, targetWord string) bool {
	n1 := ""
	n2 := ""
	t := ""
	for _, v := range firstWord {
		if v == 'a' {
			n1 = n1 + "0"
		} else {
			n1 = n1 + strconv.Itoa(int(v-'a'))
		}
	}
	for _, v := range secondWord {
		if v == 'a' {
			n2 = n2 + "0"
		} else {
			n2 = n2 + strconv.Itoa(int(v-'a'))
		}
	}
	for _, v := range targetWord {
		if v == 'a' {
			t = t + "0"
		} else {
			t = t + strconv.Itoa(int(v-'a'))
		}
	}
	for n1[0] == '0' && len(n1) > 1 {
		n1 = n1[1:]
	}
	for n2[0] == '0' && len(n2) > 1 {
		n2 = n2[1:]
	}
	for t[0] == '0' && len(t) > 1 {
		t = t[1:]
	}

	i1, _ := strconv.Atoi(n1)
	i2, _ := strconv.Atoi(n2)
	t1, _ := strconv.Atoi(t)
	return i1+i2 == t1
}

/*
Example 1:

Input: ops = ["5","2","C","D","+"]
Output: 30
Explanation:
"5" - Add 5 to the record, record is now [5].
"2" - Add 2 to the record, record is now [5, 2].
"C" - Invalidate and remove the previous score, record is now [5].
"D" - Add 2 * 5 = 10 to the record, record is now [5, 10].
"+" - Add 5 + 10 = 15 to the record, record is now [5, 10, 15].
The total sum is 5 + 10 + 15 = 30.
*/
func CalPoints(ops []string) int {

	sum := 0
	arr := []int{}
	for _, v := range ops {
		switch v {
		case "C":
			arr = arr[:len(arr)-1]
		case "D":
			cur := arr[len(arr)-1]
			arr = append(arr, cur*2)
		case "+":
			pre := arr[len(arr)-1]
			ppre := arr[len(arr)-2]
			arr = append(arr, pre+ppre)
		default:
			pre, _ := strconv.Atoi(v)
			arr = append(arr, pre)
		}
	}

	for _, v := range arr {
		sum += v
	}
	return sum
}

/*
Example 1:

Input: nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
Output: [3,2]
Explanation: The values that are present in at least two arrays are:
- 3, in all three arrays.
- 2, in nums1 and nums2.
*/
func TwoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m1 := make(map[int]int, len(nums1))
	m2 := make(map[int]int, len(nums2))
	m3 := make(map[int]int, len(nums3))
	for _, v := range nums1 {
		m1[v]++
	}
	for _, v := range nums2 {
		m2[v]++
	}
	for _, v := range nums3 {
		m3[v]++
	}
	res := []int{}

	for _, v := range nums1 {
		if m2[v] > 0 || m3[v] > 0 {
			res = append(res, v)
		}
	}
	for _, v := range nums2 {
		if m1[v] > 0 || m3[v] > 0 {
			res = append(res, v)
		}
	}
	for _, v := range nums3 {
		if m1[v] > 0 || m2[v] > 0 {
			res = append(res, v)
		}
	}

	m4 := make(map[int]int, len(res))
	for _, v := range res {
		m4[v]++
	}
	keys := []int{}
	for k, _ := range m4 {
		keys = append(keys, k)
	}
	return keys
}

/*
Example 1:

Input: nums = [5,3,6,1,12], original = 3
Output: 24
Explanation:
- 3 is found in nums. 3 is multiplied by 2 to obtain 6.
- 6 is found in nums. 6 is multiplied by 2 to obtain 12.
- 12 is found in nums. 12 is multiplied by 2 to obtain 24.
- 24 is not found in nums. Thus, 24 is returned.
*/
func FindFinalValue(nums []int, original int) int {

	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	for m[original] > 0 {

		original = original * 2
	}
	return original
}

/*
Example 1:

Input: words = ["a","b","c","ab","bc","abc"], s = "abc"
Output: 3
Explanation:
The strings in words which are a prefix of s = "abc" are:
"a", "ab", and "abc".
Thus the number of strings in words which are a prefix of s is 3.
*/
func CountPrefixes(words []string, s string) int {

	count := 0
	for _, v := range words {

		if len(v) > len(s) {
			continue
		}
		if v == s {
			count++
			continue
		}

		if v == s[:len(v)] {
			count++
		}
	}
	return count
}

/*
Example 1:

Input: s = "abbb"
Output: true
Explanation:
The 'a's are at indices 0, 1, and 2, while the 'b's are at indices 3, 4, and 5.
Hence, every 'a' appears before every 'b' and we return true.
*/
func CheckString(s string) bool {

	a := -1
	b := -1
	for i, v := range s {
		if v == 'a' {
			a = i
		}
		if v == 'b' && b == -1 {
			b = i
		}
	}

	if b == -1 || a == -1 {
		return true
	}
	return a < b

}

/*
Example 1:

Input: nums = [0,1,2]
Output: 0
Explanation:
i=0: 0 mod 10 = 0 == nums[0].
i=1: 1 mod 10 = 1 == nums[1].
i=2: 2 mod 10 = 2 == nums[2].
All indices have i mod 10 == nums[i], so we return the smallest index 0.
*/
func SmallestEqual(nums []int) int {
	index := -1

	for i, v := range nums {
		if i%10 == v {
			index = i
			break
		}
	}
	return index
}

type RecentCounter struct {
	RequestQueue []int
}

func Constructor() RecentCounter {

	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {

	this.RequestQueue = append(this.RequestQueue, t)

	index := -1
	for i := 0; i < len(this.RequestQueue); i++ {
		//[1],[100],[3001],[3002]
		if this.RequestQueue[i] < t-3000 {
			index = i
		}
	}

	if index == -1 {
		return len(this.RequestQueue)
	} else {
		return len(this.RequestQueue[index+1:])
	}

}

/*
Example 1:

Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1.
No two values have the same number of occurrences.
*/
func UniqueOccurrences(arr []int) bool {
	m := make(map[int]int, len(arr))
	for _, v := range arr {
		m[v]++
	}
	key := make(map[int]int, len(arr))
	for k, v := range m {
		if _, ok := key[v]; !ok {
			key[v] = k
		} else {
			return false
		}
	}
	return true
}

/*
Example 1:

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].
*/
func SortedSquares(nums []int) []int {
	res := []int{}
	for _, v := range nums {

		res = append(res, v*v)
	}
	sort.Ints(res)
	return res
}

func KWeakestRows(mat [][]int, k int) []int {
	res := []int{}
	soldiers := make(map[int]int, k)
	for i, v := range mat {
		sum := 0
		for _, vv := range v {
			sum += vv
		}
		soldiers[i] = sum
	}
	return res

}

func CountBalls(lowLimit int, highLimit int) int {

	n := highLimit - lowLimit + 1
	m := make(map[int]int, n)
	max := math.MinInt32
	for i := lowLimit; i <= highLimit; i++ {

		sum := 0
		tmp := i
		for tmp > 0 {
			sum += tmp % 10
			tmp /= 10
		}
		m[sum]++
		if m[sum] > max {
			max = m[sum]
		}
	}
	return max

}
func MinSubsequence(nums []int) []int {

	sort.Ints(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	tmp := []int{}
	bigger := 0
	for i := len(nums) - 1; i >= 0; i-- {
		others := sum - bigger
		if others >= bigger {

			bigger += nums[i]
			tmp = append(tmp, nums[i])
		} else {
			break
		}
	}
	return tmp
}
