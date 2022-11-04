package level9

import (
	"math"
	"sort"
	"strings"
)

func CanThreePartsEqualSum(arr []int) bool {

	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	iSum := 0
	i := 0
	for ; i < len(arr); i++ {
		iSum += arr[i]
		if iSum == sum/3 {
			break
		}
	}
	j := i + 1
	for ; j < len(arr); j++ {
		iSum += arr[j]
		if iSum == sum/3*2 {
			break
		}
	}
	if j < len(arr)-1 {
		return true
	}
	return false
}
func twoPartionEuqalArr(arr []int) bool {

	index := 0
	sum := 0
	for _, v := range arr {
		sum += v
	}
	beforeSum := arr[index]
	afterSum := sum - beforeSum
	for index < len(arr) {
		if beforeSum == afterSum {
			return true
		}
		index++
		beforeSum += arr[index]
		afterSum -= arr[index]

	}
	return false
}
func ContainsPattern(arr []int, m int, k int) bool {

	//1,2,3,4
	for i := 0; i <= len(arr); i++ {
		j := i + m
		count := k - 1
		for j < len(arr) && count > 0 {
			if j+m <= len(arr) && isEqualArr(arr[j:j+m], arr[i:i+m]) {
				count--
				j += m
			} else {
				break
			}
		}
		if count <= 0 {
			return true
		}
	}

	return false
}
func isEqualArr(arr1 []int, arr2 []int) bool {

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
func RepeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return true
	}
	for i := 1; i <= len(s)/2; i++ {

		tmp := s[:i]
		q := i
		for q < len(s) && s[q:q+i] == tmp {
			q += i
		}
		if q >= len(s) {
			return true
		}
	}
	return false
}
func ReorderSpaces(text string) string {

	spaceCount := 0
	textStr := strings.Split(text, " ")
	tStr := []string{}
	for _, v := range textStr {
		if v != "" {
			tStr = append(tStr, v)
		}
	}
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			spaceCount++
		}
	}
	//直接返回
	if len(tStr) < 1 {
		return text
	}
	//说明只有一个字符串，直接把空格放到最后
	if len(tStr) == 1 {

		allSpace := ""
		for i := 0; i < len(text)-len(tStr[0]); i++ {
			allSpace += " "
		}
		return tStr[0] + allSpace
	}
	newSpace := spaceCount / (len(tStr) - 1)
	space := ""
	for i := 0; i < newSpace; i++ {
		space += " "
	}
	res := ""
	mod := spaceCount % (len(tStr) - 1)
	modspace := ""
	for i := 0; i < mod; i++ {
		modspace += " "
	}
	for i := 0; i < len(tStr); i++ {
		if i == len(tStr)-1 {
			res = res + tStr[i] + modspace
		} else {
			res = res + tStr[i] + space
		}
	}
	return res
}

func FindMaxAverage(nums []int, k int) float64 {

	oriSum := 0
	for i := 0; i < k; i++ {
		oriSum += nums[i]
	}
	res := oriSum
	for i := 0; i < len(nums)-k; i++ {

		oriSum = oriSum - nums[i] + nums[i+k]
		res = max(oriSum, res)
	}
	return float64(res) / float64(k)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindSecondMinimumValue(root *TreeNode) int {
	//root.val = min(root.left.val, root.right.val)
	if root == nil {
		return -1
	}
	minNum := root.Val
	secMinNum := math.MaxInt64
	findSecMinSum(root, &minNum, &secMinNum)
	if secMinNum != math.MaxInt64 {
		return -1
	}
	return secMinNum
}
func findSecMinSum(root *TreeNode, minNum, secMinNum *int) {

	if root == nil {
		return
	}
	if root.Val > *minNum && root.Val < *secMinNum {
		*secMinNum = root.Val
	}
	findSecMinSum(root.Left, minNum, secMinNum)
	findSecMinSum(root.Right, minNum, secMinNum)
}
func MostCommonWord(paragraph string, banned []string) string {
	paragraphArr := getStr(paragraph)
	m := make(map[string]int)
	for _, v := range paragraphArr {
		m[v]++
	}

	bannedmap := make(map[string]bool)
	for _, v := range banned {
		bannedmap[v] = true
	}
	maxFreq := 0
	res := ""
	for _, v := range paragraphArr {
		if m[v] >= maxFreq && !bannedmap[v] {
			maxFreq = m[v]
			res = v
		}
	}
	return res

}
func getStr(paragraph string) []string {
	res := []string{}
	for i := 0; i < len(paragraph); i++ {

		j := i
		tmp := []byte{}
		for j < len(paragraph) && isLetter(paragraph[j]) {
			tmp = append(tmp, paragraph[j])
			j++
		}
		if len(tmp) > 0 {
			res = append(res, strings.ToLower(string(tmp)))
			i = j - 1
		}
	}
	return res
}
func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'z')
}
func AddToArrayForm(num []int, k int) []int {

	kArr := []int{}
	for k > 0 {
		kArr = append(kArr, k%10)
		k /= 10
	}
	for i := 0; i < len(kArr)/2; i++ {
		kArr[i], kArr[len(kArr)-i-1] = kArr[len(kArr)-i-1], kArr[i]
	}
	kArrIndex := len(kArr) - 1
	numIndex := len(num) - 1
	c := 0
	ans := make([]int, max(len(num), len(kArr))+1)
	ansIndex := len(ans) - 1
	for numIndex >= 0 && kArrIndex >= 0 {

		ans[ansIndex] = (num[numIndex] + kArr[kArrIndex] + c) % 10
		c = (num[numIndex] + kArr[kArrIndex] + c) / 10
		ansIndex--
		numIndex--
		kArrIndex--
	}

	for numIndex >= 0 {
		ans[ansIndex] = (num[numIndex] + c) % 10
		c = (num[numIndex] + c) / 10
		numIndex--
		ansIndex--
	}
	for kArrIndex >= 0 {
		ans[ansIndex] = (kArr[kArrIndex] + c) % 10
		c = (kArr[kArrIndex] + c) / 10
		kArrIndex--
		ansIndex--
	}
	if c != 0 {
		ans[0] = c
		return ans
	}
	return ans[1:]
}
func AreAlmostEqual(s1 string, s2 string) bool {

	byteS1 := []byte(s1)
	byteS2 := []byte(s2)
	//先找出第一个不一样的
	i := 0
	for i < len(byteS1) && byteS1[i] == byteS2[i] {
		i++
	}
	if i >= len(byteS1) {
		return true
	}
	j := i + 1
	for j < len(byteS1) && byteS1[j] == byteS2[j] {
		j++
	}
	if j >= len(byteS1) {
		return false
	}
	byteS1[i], byteS1[j] = byteS1[j], byteS1[i]
	return string(byteS1) == string(byteS2)

}
func FindClosestNumber(nums []int) int {

	mostClose := math.MaxInt64
	res := 0
	sort.Ints(nums)
	for _, v := range nums {
		if int(math.Abs(float64(v)-float64(0))) <= mostClose {
			res = v
			mostClose = int(math.Abs(float64(v) - float64(0)))
		}
	}
	return res
}
