package level6

import (
	"sort"
	"strconv"
)

func SearchMatrix(matrix [][]int, target int) bool {

	row := len(matrix)
	rol := len(matrix[0])

	//1.先找再哪一行
	index := 0
	for index < row && matrix[index][rol-1] < target {
		index++
	}
	if index >= row {
		return false
	} else {
		//2.二分查列
		l := 0
		r := rol - 1
		if r == l {
			return matrix[index][l] == target
		}
		for l < r {
			mid := l + ((r - l) >> 1)
			if matrix[index][mid] > target {
				r = mid - 1
			} else if matrix[index][mid] < target {
				l = mid + 1
			} else {
				return true
			}

		}
	}
	return false

}

func SubsetsWithDup(nums []int) [][]int {

	res := [][]int{}
	sort.Ints(nums)
	for i := 1; i <= len(nums); i++ {
		dfsSubsets(i, nums, 0, 0, &res, []int{})
	}
	for i := 1; i < len(nums); i++ {

		for _, v := range res {
			if len(v) == i {

			}
		}
	}
	return res
}
func dfsSubsets(end int, nums []int, start, count int, res *[][]int, tmp []int) {

	if end == count {
		c := make([]int, len(tmp))
		copy(c, tmp)

		*res = append(*res, c)
		return
	}
	for i := start; i < len(nums); i++ {
		// 这里是去重的关键逻辑,本次不取重复数字，下次循环可能会取重复数字
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		tmp = append(tmp, nums[i])
		dfsSubsets(end, nums, i+1, count+1, res, tmp)
		tmp = tmp[:len(tmp)-1]
	}
	return
}

func ZigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}
	//1.从左到右
	//2.下一层从右到左
	res := [][]int{}
	queue := []*TreeNode{root}

	level := 1
	for len(queue) > 0 {

		l := len(queue)
		tmp := make([]int, l)
		j := 0
		if level%2 == 0 {
			j = l - 1
		}
		for i := 0; i < l; i++ {

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

			if level%2 == 0 {
				tmp[j] = queue[i].Val
				j--
			} else {
				tmp[j] = queue[i].Val
				j++
			}
		}
		level++
		c := make([]int, len(tmp))
		copy(c, tmp)
		res = append(res, c)
		queue = queue[l:]
	}
	return res
}

/*
Example 1:

Input: s = "12"
Output: 2
Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: s = "226"
Output: 3
Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
*/
func NumDecodings(s string) int {

	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	} else {
		dp[1] = 0
	}

	//123 4
	//1,2,3|12,3|1,23|
	//1,2,3,4|1,23,4|12,3,4
	//12 34
	//1,2,3 12,3, 1,23
	for i := 2; i <= len(s); i++ {

		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') >= 10 && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
	//226 6
	//22 66
	//dp[i] = if numMapping[i-1:] && if numMapping[i-2:] {
	//dp[i] = dp[i-2] + dp[i-1]
	//}else if numMapping[i-1:]{
	//dp[i] = dp[i-1]
	//}else if numMapping[i-2:]{
	//dp[i] = dp[i-2]
	//}
	//dp[i-1] + if numMapping[i-2:] dp[i-2]

	//dfsNumMapping(&res,0,"",numMapping,s,0)

}
func dfsNumMapping(res *int, count int, tmps string, numMapping map[int]string, s string, start int) {

	if len(s) == count {
		if tmps == s {
			*res++
		}
		return
	}
	for j := 1; j <= 2; j++ {
		//12
		num := 0
		if start+j == len(s) {
			num, _ = strconv.Atoi(s[start:])
		} else if start+j < len(s) {
			num, _ = strconv.Atoi(s[start : start+j])
		} else {
			return
		}

		if _, ok := numMapping[num]; ok {
			tmps += strconv.Itoa(num)
			//如果找到,那么继续decode下一个s的字符
			dfsNumMapping(res, count+j, tmps, numMapping, s, start+j)
			tmps = tmps[:len(tmps)-len(strconv.Itoa(num))]
		}
	}

	return
}

func MaxNumberOfFamilies(n int, reservedSeats [][]int) int {

	res := 0
	isSet := make([][]bool, n)

	for i := 0; i < len(isSet); i++ {
		isSet[i] = make([]bool, 10)
	}
	for i := 0; i < len(reservedSeats); i++ {
		row := reservedSeats[i][0]
		rol := reservedSeats[i][1]
		isSet[row-1][rol-1] = true
	}

	for i := 0; i < len(isSet); i++ {
		rol := 0
		countseat := 0
		for rol < len(isSet[0])-1 {
			for rol < len(isSet[0])-1 && !isSet[i][rol] {
				countseat++
				rol++

			}
			if countseat < 4 {
				for rol < len(isSet[0])-1 && isSet[i][rol] {
					rol++
				}
				countseat = 0
			} else {
				res += countseat / 4
				countseat = 0
			}
		}

	}

	return res
}
