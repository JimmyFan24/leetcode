package level3

import (
	"sort"
	"strconv"
)

func LargestLocal(grid [][]int) [][]int {

	res := make([][]int, len(grid)-2)

	for i := 0; i < len(grid[0])-2; i++ {
		tmp := []int{}
		for j := 0; j < len(grid)-2; j++ {

			max := getMaxNum(grid, i, i+3, j, j+3)
			tmp = append(tmp, max)
		}
		res = append(res, tmp)
	}
	return res
}
func getMaxNum(grid [][]int, starti, endi, startj, endj int) int {

	m := 0
	for i := starti; i < endi; i++ {
		for j := startj; j < endj; j++ {
			if grid[i][j] > m {
				m = grid[i][j]
			}
		}
	}
	return m
}
func CountPoints(rings string) int {
	rod := make(map[rune][]rune, 10)
	//"B0B6G0R6R0R6G9"
	for i, v := range rings {

		if v >= '0' && v <= '9' {
			rod[v] = append(rod[v], rune(rings[i-1]))
		}
	}
	count := 0
	for _, v := range rod {

		if len(v) < 3 {
			continue
		}
		tMap := make(map[rune]int, 3)
		for _, r := range v {
			tMap[r]++
		}
		if tMap['R'] >= 1 && tMap['G'] >= 1 && tMap['B'] >= 1 {
			count++
		}

	}
	return count
}
func FindNumbers(nums []int) int {
	sum := 0

	for _, v := range nums {
		num := getEvenDigits(v)
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}
func getEvenDigits(num int) int {

	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	return count
}

// -5 5 -4 4 0
func SumZero(n int) []int {
	if n == 1 {
		return []int{0}
	}
	res := []int{}
	res = append(res, n)
	if (n-1)%2 == 0 {
		res = append(res, -n)
		res = append(res, 0)
		for i := 1; i < (n-1)/2; i++ {
			res = append(res, n-i)
			res = append(res, -(n - i))
		}

	} else {
		res = append(res, -n)
		for i := 1; i < n/2; i++ {
			res = append(res, n-i)
			res = append(res, -(n - i))
		}
	}
	return res
}

/*
Example 1:

Input: words = ["pay","attention","practice","attend"], pref = "at"
Output: 2
Explanation: The 2 strings that contain "at" as a prefix are: "attention" and "attend".
*/
func PrefixCount(words []string, pref string) int {

	count := 0
	for _, v := range words {

		if v[:len(pref)] == pref {
			count++
		}
	}
	return count
}
func FindGCD(nums []int) int {

	min := nums[0]
	max := nums[0]
	for _, v := range nums {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	if max%min == 0 {
		return min
	}
	return gcd(max, min)

}
func gcd(a, b int) int {

	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func SelfDividingNumbers(left int, right int) []int {

	res := []int{}
	for i := left; i <= right; i++ {
		if isSelfDividingNumbers(i) {
			res = append(res, i)
		}
	}
	return res
}
func isSelfDividingNumbers(num int) bool {

	for num > 0 {

		tmp := num % 10
		if tmp == 0 {
			return false
		}
		if num%tmp != 0 {
			return false
		}
		num = num / 10
	}
	return true
}
func TargetIndices(nums []int, target int) []int {

	res := []int{}
	sort.Ints(nums)
	for i, v := range nums {
		if v == target {
			res = append(res, i)
		}
	}
	return res
}
func NumberOfPairs(nums []int) []int {

	res := make([]int, 2)
	nMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if _, ok := nMap[v]; !ok {
			nMap[v] = i
		} else {
			res[0]++
			delete(nMap, v)
		}
	}
	res[1] = len(nMap)
	return res
}
func GenerateTheString(n int) string {
	res := ""
	if n%2 == 0 {
		for i := 0; i < n-1; i++ {
			res += "a"
		}
		res += "b"
	} else {
		for i := 0; i < n; i++ {
			res += "a"
		}

	}
	return res
}
func SquareIsWhite(coordinates string) bool {
	oddLetter := []rune{'a', 'c', 'e', 'g'}

	num, _ := strconv.Atoi(string(coordinates[1]))
	if num%2 != 0 {
		for _, v := range oddLetter {
			if rune(coordinates[0]) == v {
				return false
			}
		}
		return true
	} else {
		for _, v := range oddLetter {
			if rune(coordinates[0]) == v {
				return true
			}
		}

	}
	return false
}
func HalvesAreAlike(s string) bool {
	n := len(s)
	a := s[:n/2]
	b := s[n/2:]

	counta := 0
	countb := 0

	for _, v := range a {
		if isVolwels(v) {
			counta++
		}
	}
	for _, v := range b {
		if isVolwels(v) {
			countb++
		}
	}
	return countb == counta
}
func isVolwels(w rune) bool {

	return w == 'a' || w == 'e' || w == 'i' || w == 'o' || w == 'u' || w == 'A' || w == 'E' || w == 'I' || w == 'O' || w == 'U'

}

func ReversePrefix(word string, ch byte) string {

	res := ""
	wByte := []byte(word)
	index := 0
	for i, v := range wByte {
		if v == ch {
			index = i
			break
		}
	}
	for i := index; i >= 0; i-- {
		res += string(wByte[i])
	}
	res += word[index+1:]
	return res
}

type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}
func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:

		if this.Big >= 1 {
			this.Big--
			return true
		} else {
			return false
		}
	case 2:
		if this.Medium >= 1 {
			this.Medium--
			return true
		} else {
			return false
		}
	case 3:
		if this.Small >= 1 {
			this.Small--
			return true
		} else {
			return false
		}
	default:
		return true
	}

}
func SumOfUnique(nums []int) int {

	nmap := make(map[int]int, len(nums))
	sum := 0
	for _, v := range nums {
		nmap[v]++
	}
	for _, v := range nums {
		if nmap[v] == 1 {
			sum += v
		}
	}
	return sum
}
func CountOperations(num1 int, num2 int) int {
	count := 0
	for num1 > 0 && num2 > 0 {

		if num1 > num2 {
			num1 = num1 - num2
			count++
		} else {
			num2 = num2 - num1
			count++
		}
	}
	return count
}
func MergeAlternately(word1 string, word2 string) string {

	res := ""
	for len(word1) > 0 && len(word2) > 0 {
		res += string(word1[0]) + string(word2[0])
		word1 = word1[1:]
		word2 = word2[1:]
	}

	for len(word1) > 0 {
		res += word1
	}
	for len(word2) > 0 {
		res += word2
	}
	return res
}
func ArrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums)-1; {
		sum += getMin(nums[i], nums[i+1])
		i += 2
	}
	return sum
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BusyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0

	for i := 0; i < len(startTime); i++ {
		if endTime[i] >= queryTime && startTime[i] <= queryTime {
			count++
		}
	}
	return count
}
func IsSameAfterReversals(num int) bool {
	if num < 10 {
		return true
	}

	if num%10 == 0 {
		return false
	}
	return true
}

func RemovePalindromeSub(s string) int {

	if len(s) == 0 {
		return 0
	}
	for i, j := 0, len(s)-1; i < j; {

		if s[i] != s[j] {
			return 2
		}
		i++
		j--
	}
	return 1
}

func DivideArray(nums []int) bool {

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; {
		if nums[i] != nums[i+1] {
			return false
		}
		i += 2
	}
	return true
}

func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	if root1 != nil && root2 != nil {
		root.Val = root1.Val + root2.Val
		root.Left = MergeTrees(root1.Left, root2.Left)
		root.Right = MergeTrees(root1.Right, root2.Right)
		return root
	} else if root1 == nil && root2 != nil {
		return root2
	} else if root1 != nil && root2 == nil {
		return root1
	}
	return nil
}
func FirstPalindrome(words []string) string {

	res := ""
	for _, v := range words {
		if isPalindrome(v) {
			res = v
			break
		}
	}
	return res
}
func isPalindrome(s string) bool {

	for i, j := 0, len(s)-1; i < j; {

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
func SubsetXORSum(nums []int) int {

	sum := 0

	for i := 1; i < len(nums); i++ {
		getSubsetSum(nums, 0, 0, i, &sum)
	}
	return sum
}
func getSubsetSum(nums []int, start, count, index int, sum *int) {
	if index == count {

		return
	}

	for j := start; j < len(nums); j++ {
		*sum += *sum ^ nums[j]
		getSubsetSum(nums, j+1, count+1, index, sum)
	}
	return
}
