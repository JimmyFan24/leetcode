package level4

import (
	"math"
	"strconv"
	"strings"
)

/*
Example 1:

Input: s = "leetcode", t = "coats"
Output: 7
Explanation:
- In 2 steps, we can append the letters in "as" onto s = "leetcode", forming s = "leetcodeas".
- In 5 steps, we can append the letters in "leede" onto t = "coats", forming t = "coatsleede".
"leetcodeas" and "coatsleede" are now anagrams of each other.
We used a total of 2 + 5 = 7 steps.
It can be shown that there is no way to make them anagrams of each other with less than 7 steps
*/
func MinSteps(s string, t string) int {

	m := make(map[rune]int, 26)

	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		m[v]++
	}
	count := 0
	sM := make(map[rune]int, 26)
	tM := make(map[rune]int, 26)
	for _, v := range s {
		sM[v]++
	}
	for _, v := range t {
		tM[v]++
	}
	for k, _ := range m {
		if sM[k] == tM[k] && sM[k] != 0 {
			continue
		} else if sM[k] != tM[k] {
			count += int(math.Abs(float64(sM[k]) - float64(tM[k])))
		}
	}
	return count
}

/*
A complex number can be represented as a string on the form "real+imaginaryi" where:

real is the real part and is an integer in the range [-100, 100].
imaginary is the imaginary part and is an integer in the range [-100, 100].
i2 == -1.
Given two complex numbers num1 and num2 as strings, return a string of the complex number that represents their multiplications.
Example 1:

Input: num1 = "1+1i", num2 = "1+1i"
Output: "0+2i"
Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.
*/

/*
"1+-1i"
"1+-1i"
*/
func ComplexNumberMultiply(num1 string, num2 string) string {

	num1Str := strings.Split(num1, "+")
	num2Str := strings.Split(num2, "+")

	num1Real, _ := strconv.Atoi(num1Str[0])
	num2Real, _ := strconv.Atoi(num2Str[0])
	real := num1Real * num2Real

	newNum1 := strings.Replace(num1Str[1], "i", "", -1)
	newNum2 := strings.Replace(num2Str[1], "i", "", -1)
	imageNum1, _ := strconv.Atoi(newNum1)
	imageNum2, _ := strconv.Atoi(newNum2)
	imaginary := num1Real*imageNum2 + imageNum1*num2Real
	real = real - imageNum2*imageNum1
	return strconv.Itoa(real) + "+" + strconv.Itoa(imaginary) + "i"

}

/*
ib.end < ia.start，此时我们应该继续判断B的下一个区间
ib.start > ia.end，此时我们应该继续判断A的下一个区间
ia.start > ib.start and ia.end < ib.end，此时ib包含ia，那么交集就是ia。
ia.start < ib.start and ia.end > ib.end，此时ia包含ib，那么交集就是ib。
剩下就是ia和ib相交的情况，那么我们取max(ia.start, ib.start), min(ia.end, ib.end)即可。
*/
func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {

	i, j := 0, 0
	res := [][]int{}
	for i < len(firstList) && j < len(secondList) {
		if firstList[i][len(firstList[i])-1] < secondList[j][0] {
			i++
			continue
		}
		if firstList[i][0] > secondList[j][len(secondList[j])-1] {
			j++
			continue
		}
		//ia.start > ib.start and ia.end < ib.end
		if firstList[i][0] > secondList[j][0] && firstList[i][len(firstList[i])-1] < secondList[j][len(secondList[j])-1] {

			res = append(res, firstList[i])
			i++
		}
		//ia.start < ib.start and ia.end > ib.end
		if firstList[i][0] < secondList[j][0] && firstList[i][len(firstList[i])-1] > secondList[j][len(secondList[j])-1] {

			res = append(res, secondList[j])
			j++
		}
		//17--20
		//6--8
		//剩下就是ia和ib相交的情况，那么我们取max(ia.start, ib.start), min(ia.end, ib.end)即可。
		res = append(res, []int{max(firstList[i][0], secondList[j][0]), min(firstList[i][len(firstList[i])-1], secondList[j][len(secondList[j])-1])})
		if firstList[i][len(firstList[i])-1] <= secondList[j][len(secondList[j])-1] {
			i++
		} else if firstList[i][len(firstList[i])-1] >= secondList[j][len(secondList[j])-1] {
			j++
		}
	}
	return res
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
