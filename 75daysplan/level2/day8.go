package level2

import "math"

func MyNewPow(x float64, n int) float64 {

	p := x
	if n == 0{
		return float64(1)
	}
	if n >0{
		for n > 1 {
			x = x * p
			n--
		}
	}else{
		n = -n
		for n >1 {
			x = x*p
			n--
		}
		x = 1/x
	}

	return x
}
/*
Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = 3.33333.. which is truncated to 3.

Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = -2.33333.. which is truncated to -2.
 */
func Divide(dividend int, divisor int) int {
	tmpDividend := dividend

	if tmpDividend <0{
		tmpDividend = -tmpDividend
	}

	p := tmpDividend / divisor
	res := 0
	if tmpDividend % divisor == 0{
		res =  p
	}else{
		d := tmpDividend % divisor
		res = (tmpDividend - d) / divisor
	}

	if dividend<0{
		res = -res
	}
	if res > math.MaxInt32{
		res = math.MaxInt32
	}

	return res

}
func Search(nums []int, target int) int {
	//nums = [4,5,6,7,0,1,2], target = 0
	l,r := 0,len(nums)-1

	for l < r{
		mid := l + (r-l)/2
		if nums[mid] == target{
			return mid
		} else if nums[mid] > nums[l]{

			if target >= nums[l] && target < nums[mid]{
				r = mid
			}else{
				l = mid+1
			}
		}else if nums[mid] < nums[r]{

			if target > nums[mid] && target <= nums[r]{
				l = mid +1
			}else{
				r = mid
			}
		} else {
			if nums[l] == nums[mid] {
				l++
			}
			if nums[r] == nums[mid] {
				r--
			}
		}
	}
	if target == nums[l]{
		return l
	}
	return -1
}

func SearchRange(nums []int, target int) []int {

	return []int{findFirst(nums,target),findLast(nums,target)}
}
func findFirst(nums []int,target int) int{
	l,r := 0,len(nums)-1

	if len(nums)==0{
		return  -1
	}
	for l<r{
		mid := l + (r-l)>>1

		//1,2,2,3,3,4,5   target:2
		if target < nums[mid]{
			r =mid
		}
		if target > nums[mid]{
			l = mid+1
		}else{
			if mid == 0 ||nums[mid-1]!= target{
				return mid
			}
			r = mid
		}



	}
	return -1
}
func findLast(nums []int,target int) int{
	l,r := 0,len(nums)-1

	if len(nums)==0{
		return  -1
	}
	for l<r{
		mid := l + (r-l)>>1

		//1,2,2,3,3,4,5   target:2
		if target < nums[mid]{
			r =mid
		}
		if target > nums[mid]{
			l = mid+1
		}else {
			if mid == len(nums)-1 || nums[mid+1]!= target {
				return mid
			}
			l = mid+1
		}

	}
	return -1
}