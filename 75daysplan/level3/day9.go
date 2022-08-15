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
	//1 1*2+5*2
	nba := 123
	return nba

}

/*
Example 1:

Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
Output: 8
Explanation: There are 8 negatives number in the matrix.
*/
func CountNegatives(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] < 0 {
				count++
			}
		}
	}
	return count
}

func JudgeCircle(moves string) bool {

	m := make(map[rune]int, 4)
	for _, v := range moves {
		m[v]++
	}
	return m['L'] == m['R'] && m['U'] == m['D']
}

func FinalPrices(prices []int) []int {

	res := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		j := i + 1
		for j < len(prices) {
			if prices[j] > prices[i] {
				j++
			} else {
				break
			}

		}
		if (j == len(prices)-1 && prices[j] > prices[i]) || j == len(prices) {
			res[i] = prices[i]
		} else if j < len(prices) {
			res[i] = prices[i] - prices[j]
		}

	}
	return res
}

type NewNode struct {
	Val      int
	Children []*NewNode
}

func Postorder(root *NewNode) []int {

	if root == nil {
		return nil
	}
	res := []int{}

	for _, node := range root.Children {

		nList := Postorder(node)
		if len(nList) > 1 {
			for _, v := range nList {
				res = append(res, v)
			}
		} else if len(nList) == 1 {
			res = append(res, node.Val)
		}

	}
	res = append(res, root.Val)
	return res
}

/*
Example 1:

Input: nums = [1,1,1]
Output: 3
Explanation: You can do the following operations:
1) Increment nums[2], so nums becomes [1,1,2].
2) Increment nums[1], so nums becomes [1,2,2].
3) Increment nums[2], so nums becomes [1,2,3].
*/
func MinOperations(nums []int) int {

	count := 0

	if len(nums) <= 1 {
		count = 0
	}

	for i := 0; i < len(nums)-1; i++ {
		j := i + 1
		for nums[j] <= nums[i] {
			nums[j]++
			count++
		}

	}

	return count
}

func ReplaceElements(arr []int) []int {

	res := make([]int, len(arr))
	max := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		//[17,18,5,4,6,1]
		//[0,6,0,0,1,-1]
		if arr[i] > max {
			res[i] = max
			max = arr[i]
			continue
		}
		res[i] = max

	}
	res[len(res)-1] = -1
	return res
}

/*
Example 1:

Input: items1 = [[1,1],[4,5],[3,8]], items2 = [[3,1],[1,5]]
Output: [[1,6],[3,9],[4,5]]
Explanation:
The item with value = 1 occurs in items1 with weight = 1 and in items2 with weight = 5, total weight = 1 + 5 = 6.
The item with value = 3 occurs in items1 with weight = 8 and in items2 with weight = 1, total weight = 8 + 1 = 9.
The item with value = 4 occurs in items1 with weight = 5, total weight = 5.
Therefore, we return [[1,6],[3,9],[4,5]].
*/
func MergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	res := [][]int{}

	/*

		[[1,1],[4,5],[3,8]]
		[[3,1],[1,5]]
	*/
	for len(items1) > 0 && len(items2) > 0 {

		if items1[0][0] < items2[0][0] {
			q := 0
			for ; q < len(items2); q++ {
				if items2[q][0] == items1[0][0] {
					break
				}
			}
			if q < len(items2) {
				res = append(res, []int{items1[0][0], items1[0][1] + items2[q][1]})
				if q == len(items2)-1 {
					items2 = items2[:q]
				} else {
					tmp := items2[q+1:]
					items2 = items2[:q]
					for _, v := range tmp {
						items2 = append(items2, v)
					}
				}
			} else {
				res = append(res, []int{items1[0][0], items1[0][1]})
			}
			items1 = items1[1:]

		} else {
			p := 0
			for ; p < len(items1); p++ {
				if items1[p][0] == items2[0][0] {
					break
				}
			}
			if p < len(items1) {
				res = append(res, []int{items2[0][0], items2[0][1] + items1[p][1]})

				if p == len(items1)-1 {
					items1 = items1[:p]
				} else {
					tmp := items1[p+1:]
					items1 = items1[:p]
					for _, v := range tmp {
						items1 = append(items1, v)
					}
				}
			} else {
				res = append(res, []int{items2[0][0], items2[0][1]})
			}
			items2 = items2[1:]
		}

	}
	for len(items1) > 0 {
		res = append(res, []int{items1[0][0], items1[0][1]})
		items1 = items1[1:]
	}
	for len(items2) > 0 {
		res = append(res, []int{items2[0][0], items2[0][1]})
		items2 = items2[1:]
	}
	return res
}

func IsUnivalTree(root *TreeNode) bool {

	if root == nil {
		return true
	}
	if root.Left != nil && root.Right != nil {
		return IsUnivalTree(root.Left) &&
			IsUnivalTree(root.Right) &&
			root.Val == root.Left.Val &&
			root.Val == root.Right.Val
	} else if root.Left != nil && root.Right == nil {
		return IsUnivalTree(root.Left) &&
			root.Val == root.Left.Val
	} else if root.Left == nil && root.Right != nil {
		return IsUnivalTree(root.Right) &&
			root.Val == root.Right.Val
	}
	return true

}

func CountWords(words1 []string, words2 []string) int {

	m1 := make(map[string]int, len(words1))
	m2 := make(map[string]int, len(words2))

	for _, v := range words1 {
		m1[v]++
	}
	for _, v := range words2 {
		m2[v]++
	}
	count := 0
	for _, v := range words1 {

		if m1[v] == 1 && m2[v] == 1 {
			count++
		}
	}
	return count

}

/*
Input: root = [3,9,20,null,null,15,7]
Output: [3.00000,14.50000,11.00000]
Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
Hence return [3, 14.5, 11].
*/
func AverageOfLevels(root *TreeNode) []float64 {
	queue := []*TreeNode{root}
	res := []float64{}
	for len(queue) > 0 {
		sum := 0
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			sum += queue[i].Val
		}

		avr := float64(sum) / float64(l)
		res = append(res, avr)
		queue = queue[l:]

	}
	return res
}

/*
Example 1:

Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Explanation:
For nums1, nums1[1] = 2 is present at index 0 of nums2,
whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1,
whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].
*/
func FindDifference(nums1 []int, nums2 []int) [][]int {

	m1 := make(map[int]int, len(nums1))
	m2 := make(map[int]int, len(nums2))
	for _, v := range nums1 {
		m1[v]++
	}
	for _, v := range nums2 {
		m2[v]++
	}
	res := [][]int{}
	t1 := []int{}
	for _, v := range nums1 {
		if m1[v] >= 1 && m2[v] >= 1 {
			continue
		} else if m1[v] >= 1 {
			t1 = append(t1, v)
			delete(m1, v)

		}
	}
	res = append(res, t1)
	t2 := []int{}
	for _, v := range nums2 {
		if m1[v] >= 1 && m2[v] >= 1 {
			continue
		} else if m2[v] >= 1 {
			t2 = append(t2, v)
			delete(m2, v)

		}
	}
	res = append(res, t2)
	return res
}

/*
Example 1:

Input: s = "lEeTcOdE"
Output: "E"
Explanation:
The letter 'E' is the only letter to appear in both lower and upper case.
*/
func GreatestLetter(s string) string {

	return ""
}
func toLower(s rune) rune {
	return 'a' + (s - 'A')
}
