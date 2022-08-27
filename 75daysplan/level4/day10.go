package level4

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
从 s 中选出 最小 的字符，将它 接在 结果字符串的后面。
从 s 剩余字符中选出 最小 的字符，且该字符比上一个添加的字符大，将它 接在 结果字符串后面。
重复步骤 2 ，直到你没法从 s 中选择字符。
从 s 中选出 最大 的字符，将它 接在 结果字符串的后面。
从 s 剩余字符中选出 最大 的字符，且该字符比上一个添加的字符小，将它 接在 结果字符串后面。
重复步骤 5 ，直到你没法从 s 中选择字符。
重复步骤 1 到 6 ，直到 s 中所有字符都已经被选过。
*/
func SortString(s string) string {

	//1.getsmallest
	res := ""

	for len(s) > 0 {

		small, smalls := getSmallest(s)
		res += small
		s = smalls
		big, bigs := getBiggest(s)
		s = bigs
		res += string(big)

	}
	return res

}
func getSmallest(s string) (string, string) {

	small := ""
	index := -1
	sMap := make(map[rune]int, len(s))
	for _, v := range s {
		sMap[v]++
	}
	for len(sMap) > 0 {
		min := rune('z' + 1)
		for i, v := range s {
			if sMap[v] > 0 && v < min {
				min = v
				index = i

			}

		}
		small += string(min)
		s = s[:index] + s[index+1:]
		delete(sMap, min)
	}

	return small, s
}
func getBiggest(s string) (string, string) {
	big := ""
	index := -1
	sMap := make(map[rune]int, len(s))
	for _, v := range s {
		sMap[v]++
	}
	for len(sMap) > 0 {
		max := rune('a' - 1)
		for i, v := range s {
			if sMap[v] > 0 && v > max {
				max = v
				index = i

			}

		}
		big += string(max)
		s = s[:index] + s[index+1:]
		delete(sMap, max)
	}

	return big, s
}

func DigitCount(num string) bool {

	nMap := make(map[string]int, len(num))

	for _, v := range num {
		nMap[string(v)]++
	}
	for i, v := range num {

		its := strconv.Itoa(i)
		vts, _ := strconv.Atoi(string(v))
		if vts != nMap[its] {
			return false
		}
	}
	return true
}

func SumRootToLeaf(root *TreeNode) int {

	sum := 0
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	getLeafNumSum(root.Left, &sum, root.Val)
	getLeafNumSum(root.Right, &sum, root.Val)
	return sum
}
func getLeafNumSum(root *TreeNode, sum *int, index int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		index *= 2
		index += root.Val
		*sum += index
		return
	}
	index *= 2
	index += root.Val

	getLeafNumSum(root.Left, sum, index)
	getLeafNumSum(root.Right, sum, index)

	return
}

func MinCostToMoveChips(position []int) int {
	pMap := make(map[int]int, len(position))
	for _, v := range position {
		pMap[v]++
	}
	min := math.MaxInt32
	for k, _ := range pMap {
		count := countDistance(k, pMap)
		if count < min {
			min = count
		}
	}
	return min

}
func countDistance(k int, pMap map[int]int) int {
	count := 0
	for kk, vv := range pMap {
		if kk != k {
			if int(math.Abs(float64(kk-k)))%2 == 1 {
				count += vv
			}
		}
	}
	return count
}
func FindMiddleIndex(nums []int) int {

	index := -1
	if isSameSum(nums[:len(nums)-1], []int{0}) {
		index = len(nums) - 1
	}
	for i := 0; i < len(nums)-1; i++ {

		if isSameSum(nums[:i], nums[i+1:]) {
			index = i
			break
		}
	}

	return index
}
func isSameSum(num1 []int, num2 []int) bool {

	sum1 := 0
	sum2 := 0
	for _, v := range num1 {
		sum1 += v
	}
	for _, v := range num2 {
		sum2 += v
	}
	return sum1 == sum2
}
func RelativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int, len(arr2))
	for _, v := range arr2 {
		m[v]++
	}
	m2 := make(map[int]int, len(arr1))
	for _, v := range arr1 {
		m2[v]++
	}
	for k, _ := range m2 {
		if m[k] > 0 {
			m2[k]--
		}
	}
	keys := []int{}
	for k, _ := range m2 {

		keys = append(keys, k)
	}
	sort.Ints(keys)
	tmp := []int{}
	for _, k := range keys {
		for i := 0; i < m2[k]; i++ {
			tmp = append(tmp, k)
		}
	}
	res := []int{}
	i, j := 0, 0
	for i < len(tmp) && j < len(arr2) {
		res = append(res, arr2[j])
		for arr2[j] == tmp[i] {
			res = append(res, tmp[i])
			i++
		}
		j++
	}
	for i < len(tmp) {

		res = append(res, tmp[i])
		i++
	}
	for j < len(arr2) {

		res = append(res, arr2[j])
		j++
	}
	return res
}

func NumUniqueEmails(emails []string) int {

	domin := []string{}
	alice := []string{}
	tmp := []string{}
	for _, v := range emails {

		do := strings.Split(v, "@")
		domin = append(domin, do[1])
		alice = append(alice, do[0])
	}
	for _, v := range alice {
		//1.处理逗号
		v = strings.Replace(v, ".", "", -1)
		//2.处理加号
		index := strings.Index(v, "+")
		if index == -1 {
			tmp = append(tmp, v)
			continue
		} else {
			vv := v[:index]
			tmp = append(tmp, vv)
		}
	}

	newEmails := make(map[string]int, len(emails))

	for i := 0; i < len(emails); i++ {

		email := tmp[i] + "@" + domin[i]
		newEmails[email]++
	}

	count := len(newEmails)
	return count
}
func CommonChars(words []string) []string {

	//1.找最短的
	min := math.MaxInt32
	minStr := ""
	for _, v := range words {
		if len(v) < min {
			min = len(v)
			minStr = v
		}
	}
	res := []string{}
	minMap := make(map[rune]int, 26)
	for _, v := range minStr {
		minMap[v]++
	}
	for k, _ := range minMap {

		i := 0

		for ; i < len(words); i++ {
			if words[i] == minStr {
				continue
			}
			wordFreq := hasChar(words[i], k)
			if wordFreq < 1 {
				break
			} else if wordFreq < minMap[k] {
				minMap[k] = wordFreq
			}

		}
		if i == len(words) {
			for minMap[k] > 0 {
				res = append(res, string(k))
				minMap[k]--
			}
		}

	}

	return res
}
func hasChar(word string, b rune) int {

	wMap := make(map[rune]int, 26)
	for _, v := range word {
		wMap[v]++
	}

	return wMap[b]
}
