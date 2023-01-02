package level9

import (
	"fmt"
	"math"
	"time"
)

/*
You are given a binary array nums (0-indexed).
We define xi as the number whose binary representation is the subarray nums[0..i]
(from most-significant-bit to least-significant-bit).
For example, if nums = [1,0,1], then x0 = 1, x1 = 2, and x2 = 5.
Return an array of booleans answer where answer[i] is true if xi is divisible by 5.
*/
func PrefixesDivBy5(nums []int) []bool {

	ans := []bool{}
	curBinary := 0
	for i := 0; i < len(nums); i++ {
		curBinary = curBinary*2 + nums[i]
		curBinary %= 5
		if int64(curBinary)%5 == 0 {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
	}
	return ans

}
func DaysBetweenDates(date1 string, date2 string) int {

	// date1 = "2019-06-29", date2 = "2019-06-30"
	tm1, _ := time.Parse("2006-01-02", date1)
	tm2, _ := time.Parse("2006-01-02", date2)
	offset := math.Abs(float64(tm2.Unix()) - float64(tm1.Unix()))
	fmt.Println(offset / (24 * 3600))
	return 0

}
