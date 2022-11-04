package main

func FirstBadVersion(n int) int {
	if n <= 2 {
		if isBadVersion(n-1) == false && isBadVersion(n) == true {
			return n
		}

	}
	left, right := 0, n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
		if !isBadVersion(mid-1) && isBadVersion(mid) {
			return mid
		}
	}
	return -1
}

var b = []bool{
	true, false, true,
}

func isBadVersion(index int) bool {

	return b[index]
}
