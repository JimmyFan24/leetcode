package level10

import (
	"math"
	"strconv"
)

func ReadBinaryWatch(turnedOn int) []string {

	res := []string{}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if oneCount(i)+oneCount(j) == turnedOn {
				t := strconv.Itoa(i) + ":" + strconv.Itoa(j)
				res = append(res, t)
			}
		}
	}
	return res
}
func oneCount(num int) int {
	count := 0
	for num > 0 {
		num &= num - 1
		count++

	}
	return count
}

func ConstructRectangle(area int) []int {

	//The width W should not be larger than the length L, which means L >= W.
	//The difference between length L and width W should be as small as possible.
	res := make([]int, 2)
	diff := math.MaxInt32
	for i := 1; i*i <= area; i++ {

		if area%i == 0 {
			if area/i-i < diff {
				res[0] = area / i
				res[1] = i
				diff = area/i - i
			}
		}
	}
	return res

}

func DetectCapitalUse(word string) bool {

	if len(word) == 1 {
		return true
	}
	if word[1] > 'z' {
		for i := 2; i < len(word); i++ {
			if word[i] <= 'z' && word[i] >= 'a' {
				return false
			}
		}
	} else {
		for i := 2; i < len(word); i++ {
			if word[i] > 'z' {
				return false
			}
		}
	}
	return true
}
