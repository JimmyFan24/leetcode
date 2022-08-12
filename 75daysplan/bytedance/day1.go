package main

func MySqrt(x int) int {

	//1,2,3,4,5,6,7,8
	l, r := 1, x

	//
	for l < r {
		mid := l + (r-l)/2
		if mid*mid == x {
			return mid
		}
		if mid*mid > x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l - 1
}
func NextGreatestLetter(letters []byte, target byte) byte {

	l, r := 0, len(letters)

	for l < r {

		mid := l + (r-l)/2

		if letters[mid] > target {
			r = mid
		} else if letters[mid] <= target {
			l = mid + 1
		}
	}

	find := letters[l]
	if find <= target {
		return letters[0]
	}
	return find
}

func FindDuplicate(nums []int) int {

	//3 3 4 1 2
	//0 1 2 3 4
	slow :=  nums[0]
	fast := nums[nums[0]]

	for slow != fast{

		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	meet := 0
	for meet != slow{

		meet = nums[meet]
		slow = nums[slow]
	}
	return meet
}