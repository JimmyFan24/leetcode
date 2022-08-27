package level4

import (
	"container/heap"
	"math"
	"sort"
)

/*
Example 1:

Input: licensePlate = "1s3 PSt", words = ["step","steps","stripe","stepple"]
Output: "steps"
Explanation: licensePlate contains letters 's', 'p', 's' (ignoring case), and 't'.
"step" contains 't' and 'p', but only contains 1 's'.
"steps" contains 't', 'p', and both 's' characters.
"stripe" is missing an 's'.
"stepple" is missing an 's'.
Since "steps" is the only word containing all the letters, that is the answer.
*/
func ShortestCompletingWord(licensePlate string, words []string) string {

	formatLetter := make(map[rune]int, 26)
	for _, b := range licensePlate {
		if b >= 'A' && b <= 'Z' {
			b = 'a' + b - 'A'
			formatLetter[b]++
			continue
		}
		if b >= '0' && b <= '9' {
			continue
		}
		if b >= 'a' && b <= 'z' {
			formatLetter[b]++
			continue
		}

	}
	res := ""

	for _, word := range words {
		tmp := make(map[rune]int, 26)
		for k, v := range formatLetter {
			tmp[k] = v
		}
		if isMatchWord(word, tmp) {
			if res == "" {
				res = word
			} else {
				if len(word) < len(res) {
					res = word
				}
			}

		}
	}
	return res
}

func isMatchWord(word string, match map[rune]int) bool {

	for _, v := range word {
		if _, ok := match[v]; ok {
			match[v]--
		}
	}
	for _, v := range match {
		if v > 0 {
			return false
		}
	}
	return true
}

func IsOneBitCharacter(bits []int) bool {

	for i := 0; i < len(bits); {
		if bits[i] == 1 {
			i += 2
		}
		if bits[i] == 0 {
			if i == len(bits)-1 {
				return true
			}

		}
	}
	return false
}

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor703(k int, nums []int) KthLargest {

	k1 := KthLargest{k: k}
	for _, v := range nums {
		k1.Add(v)
	}
	return k1
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}

func (this *KthLargest) Push(v interface{}) {
	this.IntSlice = append(this.IntSlice, v.(int))
}
func (this *KthLargest) Pop() interface{} {
	a := this.IntSlice
	v := a[len(a)-1]
	this.IntSlice = a[:len(a)-1]
	return v
}
func DominantIndex(nums []int) int {

	res := -1

	m := make(map[int]int, len(nums))
	max := math.MinInt32
	for _, v := range nums {
		m[v] = 2 * v
		if v > max {
			max = v
		}
	}
	for i, v := range nums {
		if max < m[v] {
			res = i
		}
	}
	return res
}

func FindShortestSubArray(nums []int) int {
	//1.找出众数
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}

	max := math.MinInt32
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	degree := []int{}
	for k, v := range m {
		if v == max {
			degree = append(degree, k)
		}
	}
	//2.找到了众数的列表，开始比较大小
	res := math.MaxInt32
	for _, v := range degree {

		distance := -1
		for j, vv := range nums {

			if vv == v && distance != -1 && m[v] == 1 {
				distance = j - distance + 1
				break
			}
			if vv == v && m[v] == 1 {
				distance = 1
				break
			}
			if vv == v && distance == -1 {
				distance = j
			}
			if vv == v {
				m[v]--
			}

		}
		if res > distance {
			res = distance
		}
	}
	return res
}

func CountBinarySubstrings(s string) int {

	group := make([]int, len(s))

	k := 0
	group[0] = 1
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			k++
			group[k] = 1
		} else {
			group[k]++
		}
	}
	count := 0
	for i := 1; i <= k; i++ {
		count += getMinCount(group[i], group[i-1])
	}
	return count
}
func getMinCount(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
minimum element in its row and maximum in its column.
*/
func LuckyNumbers(matrix [][]int) []int {
	res := []int{}
	m := len(matrix)
	n := len(matrix[0])

	row := []int{}
	//1.找出行最小的
	for i := 0; i < m; i++ {
		minRow := math.MaxInt32
		for j := 0; j < n; j++ {
			if matrix[i][j] < minRow {
				minRow = matrix[i][j]
			}
		}
		row = append(row, minRow)
	}
	//2.找出列最大的
	rol := []int{}
	for i := 0; i < n; i++ {
		maxRol := math.MinInt32
		for j := 0; j < m; j++ {
			if matrix[j][i] > maxRol {
				maxRol = matrix[j][i]
			}
		}
		rol = append(rol, maxRol)
	}
	for i := 0; i < len(row); i++ {
		for j := 0; j < len(rol); j++ {
			if row[i] == rol[j] {
				res = append(res, row[i])
			}
		}
	}
	return res
}

func FrequencySort(nums []int) []int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	value := []int{}

	for _, v := range m {
		value = append(value, v)
	}
	sort.Ints(value)
	res := []int{}
	for i := 0; i < len(value); i++ {
		if i == len(value)-1 {
			for k, v := range m {
				if value[i] == v {
					for j := 0; j < v; j++ {
						res = append(res, k)
					}
					delete(m, k)
				}
			}
		} else {

			if value[i] != value[i+1] {
				//找到key
				for k, v := range m {
					if value[i] == v {
						for j := 0; j < v; j++ {
							res = append(res, k)
						}
						delete(m, k)
					}
				}
			} else {
				//频次相同，需要降序

				num := []int{}
				for k, v := range m {
					if value[i] == v {
						num = append(num, k)
						delete(m, k)
					}
				}
				sort.Ints(num)
				for q := len(num) - 1; q >= 0; q-- {
					for j := 0; j < value[i]; j++ {
						res = append(res, num[q])
					}
				}
				i++
			}
		}
	}
	return res

}
