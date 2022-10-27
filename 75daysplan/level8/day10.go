package level8

import (
	"sort"
	"strconv"
	"strings"
)

var Days = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func CountDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	arriveAliceTime := strings.Split(arriveAlice, "-")
	leaveAliceTime := strings.Split(leaveAlice, "-")
	arriveBobTime := strings.Split(arriveBob, "-")
	leaveBobTime := strings.Split(leaveBob, "-")
	a := getDays(arriveAliceTime)
	b := getDays(leaveAliceTime)
	c := getDays(arriveBobTime)
	d := getDays(leaveBobTime)
	return getmax(0, getmin(b, d)-getmax(a, c)+1)

}

// 需要将日期转换为天数,这样好比较
func getDays(s []string) int {

	res := 0
	m, _ := strconv.Atoi(s[0])
	d, _ := strconv.Atoi(s[1])
	for i := 0; i < m; i++ {
		res += Days[i]
	}
	return res + d
}

/*
You are given an integer array nums and an integer k.
You want to find a subsequence of nums of length k that has the largest sum.

Return any such subsequence as an integer array of length k.

A subsequence is an array that can be derived from another array by deleting some or no elements
without changing the order of the remaining elements.
*/
func MaxSubsequence(nums []int, k int) []int {

	//dp[i] : nums i个元素长度的子序列的最大值
	//dp[i] = dp[i-1]
	arr := [][]int{}
	for i, v := range nums {
		arr = append(arr, []int{i, v})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1] //按照每行的第一个元素排序
	})
	arr2 := make([][]int, k)

	copy(arr2, arr[:k])

	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i][0] < arr2[j][0] //按照每行的第一个元素排序
	})
	res := []int{}
	for _, v := range arr2 {
		res = append(res, v[1])
	}
	return res
}
func DfsSubsequence(nums []int, end int, count int, res *[]int, start int, max *int, sum int, tmp []int) {

	if count >= end {
		if sum > *max {
			c := make([]int, len(tmp))
			copy(c, tmp)
			*res = c
			*max = sum
		}
		return
	}
	for j := start; j < len(nums); j++ {
		sum += nums[j]
		tmp = append(tmp, nums[j])
		DfsSubsequence(nums, end, count+1, res, j+1, max, sum, tmp)
		sum -= nums[j]
		tmp = tmp[:len(tmp)-1]
	}

}
