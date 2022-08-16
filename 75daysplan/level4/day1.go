package level4

import "strconv"

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
