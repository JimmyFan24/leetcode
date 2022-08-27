package level5

import (
	"sort"
	"strconv"
	"strings"
)

/*
Example 1:

Input: words = ["abc","aabc","bc"]
Output: true
Explanation: Move the first 'a' in words[1] to the front of words[2],
to make words[1] = "abc" and words[2] = "abc".
All the strings are now equal to "abc", so return true.
*/
func MakeEqual(words []string) bool {

	l := 0
	for _, v := range words {
		l += len(v)
	}
	if l%len(words) != 0 {
		return false
	}

	//"abc","aabc","bc"
	m := make(map[rune]int, 26)
	for _, v := range words {

		for _, vv := range v {
			m[vv]++
		}
	}
	for _, v := range m {
		if v%len(words) != 0 {
			return false
		}
	}
	return true

}

type NumArray struct {
	Array []int
}

func Constructor303(nums []int) NumArray {
	return NumArray{Array: nums}
}

func (this *NumArray) SumRange(left int, right int) int {

	tmp := this.Array[left : right+1]
	sum := 0
	for _, v := range tmp {
		sum += v
	}
	return sum
}

/*
n空瓶 ==> 1瓶酒
一瓶酒本身+ 一个瓶 ==> n
纯酒 == n - 1

*/
func NumWaterBottles(numBottles int, numExchange int) int {

	res := numBottles

	for numBottles >= numExchange {
		//9 / 4 == 2
		num := numBottles / numExchange
		//还剩下1个没换的酒瓶: 9%4 ==1
		numBottles %= numExchange
		//加上这次新兑换的酒瓶,重新去兑换
		numBottles += num
		// 9+3 == 12
		res += num
	}
	return res
}

/*
Example 1:

Input: title = "capiTalIze tHe titLe"
Output: "Capitalize The Title"
Explanation:
Since all the words have a length of at least 3,
the first letter of each word is uppercase, and the remaining letters are lowercase.
*/
func CapitalizeTitle(title string) string {

	strList := strings.Split(title, " ")
	res := ""
	for _, v := range strList {
		if len(v) >= 3 {
			res += strings.ToUpper(string(v[0])) + strings.ToLower(v[1:]) + " "
		} else {
			res += strings.ToLower(v) + " "
		}

	}
	return res[:len(res)-1]
}

/*

Example 1:

Input: nums = [11,7,2,15]
Output: 2
Explanation: The element 7 has the element 2 strictly smaller than it and the element 11 strictly greater than it.
Element 11 has element 7 strictly smaller than it and element 15 strictly greater than it.
In total there are 2 elements having both a strictly smaller and a strictly greater element appear in nums.
*/

func CountElements(nums []int) int {
	//-71,-71,93,-71,4
	//-71,-71,-71,4,93
	sort.Ints(nums)
	count := 0

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[0] && nums[i] < nums[len(nums)-1] {
			count++
		}
	}
	return count
}

func SpecialArray(nums []int) int {

	sort.Ints(nums)
	//1,1000
	//0,0,3,4,4

	l := 0
	r := nums[len(nums)-1]
	//mid := l + ((r-l)>>2)
	for l <= r {
		mid := l + ((r - l) >> 1)
		count := binarySearchNum(nums, l, r)
		if mid < count {
			l = mid + 1
		} else if mid == count {
			return mid
		} else {
			r = mid - 1
		}
	}

	return -1
}
func binarySearchNum(nums []int, l, r int) int {
	mid := l + ((r - l) >> 1)

	count := 0
	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] >= mid {
			count++
		}
		if count > mid {
			break
		}
	}
	if count == mid {
		return count
	}
	return count
}

/*
Example 1:

Input: target = [1,2,3,4], arr = [2,4,1,3]
Output: true
Explanation: You can follow the next steps to convert arr to target:
1- Reverse sub-array [2,4,1], arr becomes [1,4,2,3]
2- Reverse sub-array [4,2], arr becomes [1,2,4,3]
3- Reverse sub-array [4,3], arr becomes [1,2,3,4]
There are multiple ways to convert arr to target, this is not the only way to do so.
*/
func CanBeEqual(target []int, arr []int) bool {

	sort.Ints(target)
	sort.Ints(arr)

	for i, v := range target {
		if arr[i] != v {
			return false
		}
	}
	return true
}

/*
Example 1:

Input: nums = [1,5,0,3,5]
Output: 3
Explanation:
In the first operation, choose x = 1. Now, nums = [0,4,0,2,4].
In the second operation, choose x = 2. Now, nums = [0,2,0,0,2].
In the third operation, choose x = 2. Now, nums = [0,0,0,0,0].
*/
func MinimumOperations(nums []int) int {

	sort.Ints(nums)
	count := 0
	for i, v := range nums {
		if v != 0 {
			for j := i; j < len(nums); j++ {
				nums[j] -= v

			}
			count++

		}
	}
	return count
}

func HasAlternatingBits(n int) bool {
	//10111
	//00000
	b := ""
	for n > 0 {
		b = strconv.Itoa(n%2) + b
		n = n / 2
	}
	for i := 0; i < len(b)-1; i++ {
		if b[i] == b[i+1] {
			return false
		}
	}
	return true
}
