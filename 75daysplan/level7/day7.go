package level7

/*
You are given two integers m and n, which represent the dimensions of a matrix.
You are also given the head of a linked list of integers.

Generate an m x n matrix that contains the integers in the linked list presented in spiral order (clockwise), starting from the top-left of the matrix. If there are remaining empty spaces, fill them with -1.

Return the generated matrix.
*/
func SpiralMatrix(m int, n int, head *ListNode) [][]int {
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for _, v := range res {
		for j := 0; j < len(v); j++ {
			v[j] = -1
		}
	}

	tmp := head
	top, buttom, left, right := 0, m-1, 0, n-1
	for tmp != nil {
		//第一行
		for j := left; j <= right; j++ {
			if tmp == nil {
				return res
			}
			res[top][j] = tmp.Val
			tmp = tmp.Next
		}
		top++
		//最后一列
		for j := top; j <= buttom; j++ {
			if tmp == nil {
				return res
			}
			res[j][right] = tmp.Val
			tmp = tmp.Next
		}
		right--
		//最后一行
		for j := right; j >= left; j-- {
			if tmp == nil {
				return res
			}
			res[buttom][j] = tmp.Val
			tmp = tmp.Next

		}
		buttom--
		//第一列
		for j := buttom; j >= top; j-- {
			if tmp == nil {
				return res
			}
			res[j][left] = tmp.Val
			tmp = tmp.Next

		}
		left++

	}

	return res

}
func CountBattleships(board [][]byte) int {

	//1.横向,一行都是X 并且
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				res++
				//1.正上方不能有,左边也不能有
				if i-1 >= 0 && board[i-1][j] == 'X' || j-1 >= 0 && board[i][j-1] == 'X' {
					res--
				}
			}
		}
	}
	//2.纵向

	return res

}

/*
Given a string s, partition the string into one or more substrings
such that the characters in each substring are unique.
That is, no letter appears in a single substring more than once.
Return the minimum number of substrings in such a partition.
Note that each character should belong to exactly one substring in a partition
*/

func PartitionString(s string) int {

	//"abacaba"
	// "hdklqkcssgxlvehva"
	//s:2 h:2 v:2 d:1 k:2 l:2 q:1 c :1 g:1 x:1 e:1 a:1
	//a:4 b:2 c :1
	tmpMap := make(map[byte]int, 26)
	i := 0
	res := 0
	for i < len(s) {

		if tmpMap[s[i]] < 1 {
			tmpMap[s[i]]++

		} else {
			tmpMap = make(map[byte]int, 26)
			tmpMap[s[i]]++
			res++
		}
		i++
	}

	return res + 1
}

/*
You are given a 0-indexed string
pattern of length n consisting of the characters 'I' meaning increasing and 'D' meaning decreasing.
A 0-indexed string num of length n + 1 is created using the following conditions:

num consists of the digits '1' to '9', where each digit is used at most once.
If pattern[i] == 'I', then num[i] < num[i + 1].
If pattern[i] == 'D', then num[i] > num[i + 1].
Return the lexicographically smallest possible string num that meets the conditions.
*/
func SmallestNumber(pattern string) string {

	//"IIIDIDDD"

	//设置初始值为0
	num := []int{0}
	subArr := [][]int{}
	used := make(map[int]bool, len(pattern))
	//"III DIDDD"
	//01232321
	res := ""
	for i := 0; i < len(pattern); i++ {
		if dfsSmallestNumber(&subArr, num, pattern, &used, i) {
			break
		}
	}

	return res
}
func dfsSmallestNumber(subArr *[][]int, num []int, pattern string, used *map[int]bool, begin int) bool {

	if len(num) >= len(pattern) {
		c := make([]int, len(num))
		copy(c, num)
		(*subArr) = append(*subArr, c)
		return true
	}
	if len(num) == 0 {
		num = append(num, begin)
	}
	index := len(num) - 1
	if pattern[index] == 'I' {
		top := num[len(num)-1] + 1
		for (*used)[top] {
			top++
		}
		for top < len(pattern) {
			num = append(num, top)
			if dfsSmallestNumber(subArr, num, pattern, used, top) {
				return true
			}
			num = num[:len(num)-1]
			top++
		}

	} else {
		top := num[len(num)-1] - 1
		for (*used)[top] {
			top--
		}
		for top > -len(pattern) {
			num = append(num, top)
			if dfsSmallestNumber(subArr, num, pattern, used, top) {
				return true
			}
			num = num[:len(num)-1]
			top++
		}
	}

	return false

}
