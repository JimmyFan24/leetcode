package level9

import (
	"fmt"
	"strconv"
	"strings"
)

func HaveConflict(event1 []string, event2 []string) bool {

	// ["01:15","02:00"], event2 = ["02:00","03:00"]
	//max(parseTime(event1[0]), parseTime(event2[0]) 左边界，取最大
	//min(parseTime(event1[1]), parseTime(event2[1])) 有边界，取最小
	if min(parseTime(event1[1]), parseTime(event2[1]))-max(parseTime(event1[0]), parseTime(event2[0])) < 0 {

		return false
	}
	return true
}
func parseTime(string2 string) int {
	time := strings.Split(string2, ":")
	hour := time[0]
	minute := time[1]

	res := 0
	i := 0
	for i < len(hour) && hour[i] == '0' {
		i++
	}
	h, _ := strconv.Atoi(hour[i:])
	res += h * 60
	m, _ := strconv.Atoi(minute)
	res += m
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Check(nums []int) bool {

	originNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {

		originNums[(i+1)%len(nums)] = nums[i]
	}
	fmt.Println(originNums)
	fmt.Println(nums)
	return true

}
func FindJudge(n int, trust [][]int) int {

	in := make(map[int]int)
	out := make(map[int]int)
	for _, v := range trust {
		out[v[0]]++
		in[v[1]]++
	}
	for i := 1; i <= n; i++ {
		if in[i] == n-1 && out[i] == 0 {
			return i
		}
	}
	return -1
}
