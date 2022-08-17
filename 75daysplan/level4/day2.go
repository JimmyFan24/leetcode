package level4

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
Example 1:

Input: s = "(1+(2*3)+((8)/4))+1"
Output: 3
Explanation: Digit 8 is inside of 3 nested parentheses in the string.
*/
func MaxDepth(s string) int {

	left := []int{}
	max := math.MinInt32
	for i, v := range s {

		if v == '(' {
			left = append(left, i)
			if len(left) > max {
				max = len(left)
			}
		}
		if v == ')' {
			left = left[:len(left)-1]
		}

	}
	if max == math.MinInt32 {
		return 0
	}
	return max

}

type OrderedStream struct {
	Id    []int
	Value []string
	Ptr   int
}

func CConstructor(n int) OrderedStream {
	return OrderedStream{
		Id:    make([]int, n),
		Ptr:   1,
		Value: make([]string, n),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {

	this.Id[idKey-1] = idKey
	this.Value[idKey-1] = value
	res := []string{}
	index := -1
	for i, id := range this.Id {
		if id == this.Ptr {
			index = i
			break
		}
	}
	if index == -1 {
		return res
	}
	j := index
	for ; j < len(this.Value); j++ {
		if this.Value[j] != "" {
			res = append(res, this.Value[j])
		} else {
			break
		}

	}
	this.Ptr = this.Id[j-1] + 1
	return res
}

/*
Example 1:

Input: start = 10, goal = 7
Output: 3
Explanation: The binary representation of 10 and 7 are 1010 and 0111 respectively. We can convert 10 to 7 in 3 steps:
- Flip the first bit from the right: 1010 -> 1011.
- Flip the third bit from the right: 1011 -> 1111.
- Flip the fourth bit from the right: 1111 -> 0111.
It can be shown we cannot convert 10 to 7 in less than 3 steps. Hence, we return 3.
*/
func MinBitFlips(start int, goal int) int {

	startBit := ""
	goalBit := ""
	for start > 0 {
		startBit = strconv.Itoa(start%2) + startBit
		start = start / 2

	}
	for goal > 0 {
		goalBit = strconv.Itoa(goal%2) + goalBit
		goal = goal / 2
	}
	count := 0
	p := len(goalBit) - 1
	q := len(startBit) - 1
	for p >= 0 && q >= 0 {
		if startBit[q] != goalBit[p] {
			count++

		}
		p--
		q--
	}
	for p >= 0 {
		if goalBit[p] != '0' {
			count++
		}
		p--
	}
	for q >= 0 {
		if startBit[q] != '0' {
			count++
		}
		q--
	}
	return count
}

func MinMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	count := 0
	for i := 0; i < len(students); i++ {
		count += int(math.Abs(float64(students[i]) - float64(seats[i])))
	}
	return count
}

func CountAsterisks(s string) int {

	sList := strings.Split(s, "|")
	//c**o
	count := 0
	for i, v := range sList {

		if i%2 == 0 {

			for _, b := range v {
				if b == '*' {
					count++
				}
			}
		}
	}
	return count
}

/*
Example 1:

Input: arr = [3,0,1,1,9,7], a = 7, b = 2, c = 3
Output: 4
Explanation: There are 4 good triplets: [(3,0,1), (3,0,1), (3,1,1), (0,1,1)].
*/
func countGoodTriplets(arr []int, a int, b int, c int) int {
	if len(arr) < 3 {
		return 0
	}
	i, k := 0, len(arr)-1
	j := i + 1
	count := 0
	for i < len(arr)-2 && j < k {
		tmp := j
		for tmp < k {
			if getMatch(arr, i, tmp, k, a, b, c) {
				count++
			}
			tmp++
		}
		k--

	}

	return countGoodTriplets(arr[1:], a, b, c) + count
}
func getMatch(arr []int, i, j, k int, a, b, c int) bool {

	return int(math.Abs(float64(arr[i])-float64(arr[j]))) <= a && int(math.Abs(float64(arr[j])-float64(arr[k]))) <= b &&
		int(math.Abs(float64(arr[i])-float64(arr[k]))) <= c
}
func BestHand(ranks []int, suits []byte) string {

	suit := suits[0]
	flush := true
	for _, v := range suits {
		if v != suit {
			flush = false
			break
		}
	}
	if flush {
		return "Flush"
	}
	m := make(map[int]int, len(ranks))
	for _, v := range ranks {
		m[v]++
	}
	max := -1
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	if max == 2 {
		return "Pair"
	} else if max >= 3 {
		return "Three of a Kind"
	}
	return "High Card"

}

/*
Example 1:

Input: s = "YazaAay"
Output: "aAa"
Explanation: "aAa" is a nice string because 'A/a' is the only letter of the alphabet in s,
and both 'A' and 'a' appear.
"aAa" is the longest nice substring.
*/
func LongestNiceSubstring(s string) string {

	if len(s) == 1 {
		return ""
	}
	res := ""
	for i := 0; i < len(s)-1; i++ {
		j := i + 1
		for isSameLetter(s[i], s[j]) && j < len(s) {
			j++
		}

		if j > i+1 {

			res = s[i:j]
			break
		}

	}
	return res
}
func isSameLetter(a, b byte) bool {

	if a == b {
		return true
	}
	if a > b {
		return a-'a' == b-'A'
	}
	return a-'A' == b-'a'
}
